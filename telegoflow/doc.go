/*
Package telegoflow provides typed conversation flows for Telego.

The package is designed as a stateful layer on top of telegohandler. A flow is a
controlled graph of steps. Each session stores the current step and typed data,
while handlers can move between steps using Context.Go, Context.Stay,
Context.Complete, and Context.Cancel.

Typical setup:

	manager := telegoflow.NewManager(telegoflow.NewMemoryStorage())
	bh.Use(manager.Middleware())

	flow, err := telegoflow.New[Registration]("registration").
		Steps(
			telegoflow.NewStep[Registration]("name").
				Enter(func(ctx *telegoflow.Context[Registration]) error {
					_, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "What's your name?"))
					return err
				}).
				Handle(func(ctx *telegoflow.Context[Registration]) error {
					ctx.Data().Name = ctx.Text()
					return ctx.Go("done")
				}, th.AnyMessageWithText()).
				CanGo("done"),
			telegoflow.NewStep[Registration]("done").
				Enter(func(ctx *telegoflow.Context[Registration]) error {
					_, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "Thanks, "+ctx.Data().Name+"!"))
					if err != nil {
						return err
					}
					return ctx.Complete()
				}),
		).
		Build()
	if err != nil {
		panic(err)
	}

	_ = manager.Register(flow)
	bh.Handle(flow.Start, th.CommandEqual("register"))
*/
package telegoflow
