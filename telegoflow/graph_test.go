package telegoflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlow_Graph(t *testing.T) {
	t.Run("linear", func(t *testing.T) {
		flow, err := New[testData]("registration").
			Steps(
				NewStep[testData]("name").CanGo("age"),
				NewStep[testData]("age").CanGo("email"),
				NewStep[testData]("email").CanGo("confirm"),
				NewStep[testData]("confirm"),
			).
			Build()
		require.NoError(t, err)

		assert.Equal(t, `flow registration (start: name)
name
└─> age
    └─> email
        └─> confirm (terminal)`, flow.Graph())
	})

	t.Run("branch_cycle_and_terminal", func(t *testing.T) {
		flow, err := New[testData]("registration").
			Steps(
				NewStep[testData]("name").CanGo("age"),
				NewStep[testData]("age").CanGo("email", "retry"),
				NewStep[testData]("email").CanGo("confirm"),
				NewStep[testData]("retry").CanGo("age"),
				NewStep[testData]("confirm"),
			).
			Build()
		require.NoError(t, err)

		assert.Equal(t, `flow registration (start: name)
name
└─> age
    ├─> email
    │   └─> confirm (terminal)
    └─> retry
        └─> age (cycle)`, flow.Graph())
	})

	t.Run("unreachable", func(t *testing.T) {
		flow, err := New[testData]("flow").
			Steps(
				NewStep[testData]("start"),
				NewStep[testData]("orphan").CanGo("orphan_done"),
				NewStep[testData]("orphan_done"),
			).
			Build()
		require.NoError(t, err)

		assert.Equal(t, `flow flow (start: start)
start (terminal)

unreachable
├─> orphan
│   └─> orphan_done (terminal)
└─> orphan_done (visited)`, flow.Graph())
	})

	t.Run("nil_flow", func(t *testing.T) {
		assert.Empty(t, (*Flow[testData])(nil).Graph())
	})
}
