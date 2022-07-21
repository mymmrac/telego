package telegohandler

import "github.com/mymmrac/telego"

type conditionalHandler struct {
	Handler    Handler
	Predicates []Predicate
}

// match Matches the current update and handler
func (ch *conditionalHandler) match(update telego.Update) bool {
	ok := true
	for _, p := range ch.Predicates {
		if !p(update.Clone()) {
			ok = false
			break
		}
	}
	return ok
}
