package telegohandler

import (
	"context"
	"time"

	"github.com/mymmrac/telego"
)

// Context is a wrapper around [context.Context] with bot handler specific methods
type Context struct {
	*ctxBase

	ctx context.Context
}

// ctxBase is a base struct for [Context] that is used to copy context without a need to copy all fields
type ctxBase struct {
	bot      *telego.Bot
	updateID int

	group *HandlerGroup
	stack []int
}

// Deadline implements [context.Context.Deadline]
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

// Done implements [context.Context.Done]
func (c *Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

// Err implements [context.Context.Err]
func (c *Context) Err() error {
	return c.ctx.Err()
}

// Value implements [context.Context.Value]
func (c *Context) Value(key any) any {
	return c.ctx.Value(key)
}

// Context returns underling [context.Context]
func (c *Context) Context() context.Context {
	return c.ctx
}

// WithContext creates new [Context] with its underling [context.Context] changed to the one provided by user
//
// Warning: Panics if nil context passed
func (c *Context) WithContext(ctx context.Context) *Context {
	if ctx == nil {
		panic("Telego: nil context not allowed")
	}

	return &Context{
		ctx:     ctx,
		ctxBase: c.ctxBase,
	}
}

// WithValue creates new [Context] with set value
func (c *Context) WithValue(key any, value any) *Context {
	return &Context{
		ctx:     context.WithValue(c.ctx, key, value),
		ctxBase: c.ctxBase,
	}
}

// WithTimeout creates new [Context] with timeout
func (c *Context) WithTimeout(timeout time.Duration) (*Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.ctx, timeout)
	return &Context{
		ctx:     ctx,
		ctxBase: c.ctxBase,
	}, cancel
}

// WithCancel creates new [Context] with cancel
func (c *Context) WithCancel() (*Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(c.ctx)
	return &Context{
		ctx:     ctx,
		ctxBase: c.ctxBase,
	}, cancel
}

// WithoutCancel creates new [Context] without cancel
func (c *Context) WithoutCancel() *Context {
	return &Context{
		ctx:     context.WithoutCancel(c.ctx),
		ctxBase: c.ctxBase,
	}
}

// Bot returns [telego.Bot]
func (c *Context) Bot() *telego.Bot {
	return c.bot
}

// UpdateID returns update ID
func (c *Context) UpdateID() int {
	return c.updateID
}

// Next executes the next handler in the stack that matches the current update
func (c *Context) Next(update telego.Update) error {
	// Go though all middlewares, subgroups and handlers
	for i := c.stack[len(c.stack)-1] + 1; i < len(c.group.routes); i++ {
		r := c.group.routes[i]
		if r.match(c.ctx, update) {
			// Update last checked route
			c.stack[len(c.stack)-1] = i

			// Go into handler or middleware
			if r.handler != nil {
				return r.handler(c, update)
			}

			// Go into subgroup
			c.group = r.group
			c.stack = append(c.stack, -1)
			return c.Next(update)
		}
	}

	// Go back to parent if nothing matches in the current group
	if c.group.parent != nil {
		c.group = c.group.parent
		c.stack = c.stack[:len(c.stack)-1]
		return c.Next(update)
	}

	// Nothing matches in any group
	return nil
}
