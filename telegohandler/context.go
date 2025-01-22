package telegohandler

import (
	"context"
	"time"

	"github.com/mymmrac/telego"
)

// Context is a wrapper around [context.Context] with bot handler specific methods
type Context struct {
	bot      *telego.Bot
	ctx      context.Context
	updateID int

	group           *HandlerGroup
	middlewareIndex int
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

// WithContext sets new underling [context.Context] returning same [Context]
//
// Warning: Panics if nil context passed
func (c *Context) WithContext(ctx context.Context) *Context {
	if ctx == nil {
		panic("Telego: nil context not allowed")
	}

	c.ctx = ctx
	return c
}

// WithValue sets new value in underling [context.Context] returning same [Context]
func (c *Context) WithValue(key any, value any) *Context {
	c.ctx = context.WithValue(c.ctx, key, value)
	return c
}

// WithTimeout sets new timeout in underling [context.Context] returning same [Context]
func (c *Context) WithTimeout(timeout time.Duration) (*Context, context.CancelFunc) {
	var cancel context.CancelFunc
	c.ctx, cancel = context.WithTimeout(c.ctx, timeout)
	return c, cancel
}

// Bot returns [telego.Bot]
func (c *Context) Bot() *telego.Bot {
	return c.bot
}

// UpdateID returns update ID
func (c *Context) UpdateID() int {
	return c.updateID
}

// Next executes the next handler in the stack that matches current update
func (c *Context) Next(update telego.Update) error {
	c.middlewareIndex++
	if len(c.group.middlewares) > c.middlewareIndex {
		return c.group.middlewares[c.middlewareIndex](c, update)
	}

	for _, r := range c.group.routes {
		if r.match(c.ctx, update) {
			if r.handler != nil {
				return r.handler(c, update)
			}
			c.group = r.group
			c.middlewareIndex = -1
			return c.Next(update)
		}
	}

	return nil
}
