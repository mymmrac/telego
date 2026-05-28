package telegoflow

import (
	"sync"
	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

// Manager coordinates flows, sessions, and telegohandler integration.
type Manager struct {
	storage Storage
	keyFunc KeyFunc
	nowFunc func() time.Time

	mu    sync.Mutex
	flows map[string]FlowRunner
	locks map[SessionKey]*sessionLock
}

type sessionLock struct {
	ch chan struct{}
}

// ManagerOption configures a manager.
type ManagerOption func(manager *Manager)

// WithKeyFunc configures custom session key extraction.
func WithKeyFunc(keyFunc KeyFunc) ManagerOption {
	return func(manager *Manager) {
		if keyFunc != nil {
			manager.keyFunc = keyFunc
		}
	}
}

// NewManager creates a flow manager.
func NewManager(storage Storage, options ...ManagerOption) *Manager {
	if storage == nil {
		storage = NewMemoryStorage()
	}

	manager := &Manager{
		storage: storage,
		keyFunc: DefaultKeyFunc,
		nowFunc: time.Now,
		flows:   make(map[string]FlowRunner),
		locks:   make(map[SessionKey]*sessionLock),
	}

	for _, option := range options {
		option(manager)
	}

	return manager
}

// Register registers flows in the manager.
func (m *Manager) Register(flows ...FlowRunner) error {
	for _, flow := range flows {
		if flow == nil {
			continue
		}
		if _, ok := m.flows[flow.ID()]; ok {
			return DuplicateFlowError{FlowID: flow.ID()}
		}
		if err := flow.register(m); err != nil {
			return err
		}
		m.flows[flow.ID()] = flow
	}
	return nil
}

// Middleware returns a telegohandler middleware that routes active sessions to flows.
//
// If the update doesn't belong to an active session, the middleware calls ctx.Next(update).
func (m *Manager) Middleware() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		key, ok := m.keyFunc(update)
		if !ok {
			return ctx.Next(update)
		}

		lock := m.lockFor(key)
		lock.lock()
		state, exists, err := m.storage.LoadSession(ctx.Context(), key)
		if err != nil {
			lock.unlock()
			return err
		}
		if !exists {
			lock.unlock()
			return ctx.Next(update)
		}
		defer lock.unlock()

		flow, ok := m.flows[state.FlowID]
		if !ok {
			return FlowNotFoundError{FlowID: state.FlowID}
		}
		if state.Expired(m.now()) {
			return flow.timeoutSession(ctx, update, state)
		}

		return flow.handleUpdate(ctx, update, state)
	}
}

// Cancel cancels the active session for this update, if any.
func (m *Manager) Cancel(ctx *th.Context, update telego.Update) error {
	key, ok := m.keyFunc(update)
	if !ok {
		return NoSessionKeyError{}
	}

	lock := m.lockFor(key)
	lock.lock()
	defer lock.unlock()

	state, exists, err := m.storage.LoadSession(ctx.Context(), key)
	if err != nil || !exists {
		return err
	}

	flow, ok := m.flows[state.FlowID]
	if !ok {
		return FlowNotFoundError{FlowID: state.FlowID}
	}

	return flow.cancelSession(ctx, update, state)
}

// ActiveSession returns information about the active session for this update, if any.
func (m *Manager) ActiveSession(ctx *th.Context, update telego.Update) (*SessionInfo, bool, error) {
	key, ok := m.keyFunc(update)
	if !ok {
		return nil, false, nil
	}

	state, exists, err := m.storage.LoadSession(ctx.Context(), key)
	if err != nil || !exists {
		return nil, false, err
	}
	if state.Expired(m.now()) {
		return nil, false, nil
	}

	return newSessionInfo(state), true, nil
}

func (m *Manager) startFlow(ctx *th.Context, update telego.Update, flow FlowRunner, data any) error {
	key, ok := m.keyFunc(update)
	if !ok {
		return NoSessionKeyError{}
	}

	lock := m.lockFor(key)
	lock.lock()
	defer lock.unlock()

	now := m.now()
	state := &SessionState{
		Key:         key,
		FlowID:      flow.ID(),
		CurrentStep: flow.startStepID(),
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if timeout := flow.timeoutDuration(); timeout > 0 {
		expiresAt := now.Add(timeout)
		state.ExpiresAt = &expiresAt
	}

	if err := flow.encodeInitialData(data, state); err != nil {
		return err
	}
	if err := m.storage.SaveSession(ctx.Context(), state); err != nil {
		return err
	}

	return flow.enterSession(ctx, update, state)
}

func (m *Manager) lockFor(key SessionKey) *sessionLock {
	m.mu.Lock()
	defer m.mu.Unlock()

	lock, ok := m.locks[key]
	if !ok {
		lock = &sessionLock{ch: make(chan struct{}, 1)}
		lock.ch <- struct{}{}
		m.locks[key] = lock
	}
	return lock
}

func (m *Manager) now() time.Time {
	return m.nowFunc()
}

func (l *sessionLock) lock() {
	<-l.ch
}

func (l *sessionLock) unlock() {
	l.ch <- struct{}{}
}
