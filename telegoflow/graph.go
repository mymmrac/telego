package telegoflow

import (
	"slices"
	"strings"
)

const (
	graphBranch = "├─> "
	graphLast   = "└─> "
	graphPipe   = "│   "
	graphSpace  = "    "
)

// Graph returns a deterministic text representation of the flow transitions.
//
// The graph starts from the configured start step and follows transitions
// declared with [Step.CanGo]. Steps with no outgoing transitions are marked as
// terminal. Cycles are marked explicitly and not expanded recursively.
func (f *Flow[T]) Graph() string {
	if f == nil {
		return ""
	}

	var builder strings.Builder
	writeGraphString(&builder, "flow ")
	writeGraphString(&builder, f.id)
	writeGraphString(&builder, " (start: ")
	writeGraphString(&builder, f.startStep)
	writeGraphString(&builder, ")\n")

	visited := make(map[string]bool, len(f.steps))
	stack := make(map[string]bool, len(f.steps))

	if f.startStep != "" {
		f.writeGraphStep(&builder, f.startStep, "", "", visited, stack)
	}

	unreachable := make([]string, 0)
	for stepID := range f.steps {
		if !visited[stepID] {
			unreachable = append(unreachable, stepID)
		}
	}
	slices.Sort(unreachable)

	if len(unreachable) > 0 {
		writeGraphString(&builder, "\nunreachable\n")
		for i, stepID := range unreachable {
			f.writeGraphStep(&builder, stepID, "", graphConnector(i == len(unreachable)-1), visited, stack)
		}
	}

	return strings.TrimRight(builder.String(), "\n")
}

func (f *Flow[T]) writeGraphStep(
	builder *strings.Builder,
	stepID string,
	prefix string,
	connector string,
	visited map[string]bool,
	stack map[string]bool,
) {
	writeGraphString(builder, prefix)
	writeGraphString(builder, connector)
	writeGraphString(builder, stepID)

	step, ok := f.steps[stepID]
	if !ok {
		writeGraphString(builder, " (missing)\n")
		return
	}

	if stack[stepID] {
		writeGraphString(builder, " (cycle)\n")
		return
	}
	if visited[stepID] {
		writeGraphString(builder, " (visited)\n")
		return
	}

	visited[stepID] = true
	stack[stepID] = true
	defer delete(stack, stepID)

	transitions := step.transitions()
	if len(transitions) == 0 {
		writeGraphString(builder, " (terminal)\n")
		return
	}

	writeGraphString(builder, "\n")
	nextPrefix := prefix + graphPrefix(connector)
	for i, transition := range transitions {
		f.writeGraphStep(builder, transition, nextPrefix, graphConnector(i == len(transitions)-1), visited, stack)
	}
}

func (s *Step[T]) transitions() []string {
	transitions := make([]string, 0, len(s.canGo))
	for stepID := range s.canGo {
		transitions = append(transitions, stepID)
	}
	slices.Sort(transitions)
	return transitions
}

func graphConnector(last bool) string {
	if last {
		return graphLast
	}
	return graphBranch
}

func graphPrefix(connector string) string {
	switch connector {
	case graphLast:
		return graphSpace
	case graphBranch:
		return graphPipe
	default:
		return ""
	}
}

func writeGraphString(builder *strings.Builder, value string) {
	_, _ = builder.WriteString(value)
}
