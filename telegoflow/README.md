# Telegoflow • Typed Conversations for Telego

Telegoflow is a stateful conversation builder for [`telego`](https://github.com/mymmrac/telego).
It helps build multi-step dialogs as controlled flows with typed session data, explicit transitions, storage-backed
state, and native integration with [`telegohandler`](../telegohandler).

> Note: Telegoflow is an optional high-level package. It does not replace `telegohandler`; it works on top of it as
> middleware and uses the same predicates.

### :clipboard: Table Of Content

<details>
<summary>Click to show • hide</summary>

- [:zap: Getting Started](#zap-getting-started)
    - [:jigsaw: Basic setup](#jigsaw-basic-setup)
    - [:left_right_arrow: Controlled transitions](#left_right_arrow-controlled-transitions)
    - [:chart_with_upwards_trend: Text graph](#chart_with_upwards_trend-text-graph)
    - [:floppy_disk: Session storage](#floppy_disk-session-storage)
    - [:hourglass_flowing_sand: Timeouts](#hourglass_flowing_sand-timeouts)
    - [:no_entry_sign: Canceling flows](#no_entry_sign-canceling-flows)
    - [:gear: Custom session keys](#gear-custom-session-keys)
- [:bricks: Core Concepts](#bricks-core-concepts)
- [:test_tube: Testing](#test_tube-testing)

</details>

## :zap: Getting Started

[▲ Go Up ▲](#telegoflow--typed-conversations-for-telego)

### :jigsaw: Basic setup

Telegoflow plugs into `telegohandler` as middleware. Register global commands that should interrupt dialogs first,
then add flow middleware, then register flow start handlers.

```go
package main

import (
    "context"
    "fmt"
    "net/mail"
    "os"
    "strconv"

    "github.com/mymmrac/telego"
    tf "github.com/mymmrac/telego/telegoflow"
    th "github.com/mymmrac/telego/telegohandler"
    tu "github.com/mymmrac/telego/telegoutil"
)

type Registration struct {
    Name  string
    Age   uint
    Email string
}

func main() {
    ctx := context.Background()
    botToken := os.Getenv("TOKEN")

    bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    updates, _ := bot.UpdatesViaLongPolling(ctx, nil)
    bh, _ := th.NewBotHandler(bot, updates)

    flows := tf.NewManager(tf.NewMemoryStorage())

    // Commands like /cancel should be registered before flow middleware.
    bh.Handle(func(ctx *th.Context, update telego.Update) error {
        return flows.Cancel(ctx, update)
    }, th.CommandEqual("cancel"))

    // Active sessions are routed by this middleware.
    bh.Use(flows.Middleware())

    registration, err := tf.New[Registration]("registration").
        Steps(
            tf.NewStep[Registration]("name").
                Enter(func(ctx *tf.Context[Registration]) error {
                    _, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "What's your name?"))
                    return err
                }).
                Handle(func(ctx *tf.Context[Registration]) error {
                    ctx.Data().Name = ctx.Text()
                    return ctx.Go("age")
                }, th.AnyMessageWithText()).
                CanGo("age"),

            tf.NewStep[Registration]("age").
                Enter(func(ctx *tf.Context[Registration]) error {
                    _, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "How old are you?"))
                    return err
                }).
                Handle(func(ctx *tf.Context[Registration]) error {
                    age, err := strconv.ParseUint(ctx.Text(), 10, 32)
                    if err != nil || age == 0 {
                        _, err = ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "Invalid age, please try again"))
                        return err
                    }

                    ctx.Data().Age = uint(age)
                    return ctx.Go("email")
                }, th.AnyMessageWithText()).
                CanGo("email"),

            tf.NewStep[Registration]("email").
                Enter(func(ctx *tf.Context[Registration]) error {
                    _, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "What's your email?"))
                    return err
                }).
                Handle(func(ctx *tf.Context[Registration]) error {
                    addr, err := mail.ParseAddress(ctx.Text())
                    if err != nil {
                        _, err = ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "Invalid email, please try again"))
                        return err
                    }

                    ctx.Data().Email = addr.Address
                    return ctx.Go("confirm")
                }, th.AnyMessageWithText()).
                CanGo("confirm"),

            tf.NewStep[Registration]("confirm").
                Enter(func(ctx *tf.Context[Registration]) error {
                    data := ctx.Data()
                    _, err := ctx.Bot().SendMessage(ctx, tu.Messagef(
                        ctx.ChatID(),
                        "Your name is %s, your age is %d, and your email is %s. Confirm? (Y/N)",
                        data.Name, data.Age, data.Email,
                    ))
                    return err
                }).
                Handle(func(ctx *tf.Context[Registration]) error {
                    switch ctx.Text() {
                    case "Y", "y":
                        return ctx.Complete()
                    case "N", "n":
                        return ctx.Go("name")
                    default:
                        _, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "Please answer Y or N"))
                        return err
                    }
                }, th.AnyMessageWithText()).
                CanGo("name"),
        ).
        StartWith("name").
        OnComplete(func(ctx *tf.Context[Registration]) error {
            _, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "Thanks for your data!"))
            return err
        }).
        Build()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    _ = flows.Register(registration)

    // Start command is checked only when there is no active session.
    bh.Handle(registration.Start, th.CommandEqual("register"))

    defer func() { _ = bh.Stop() }()
    _ = bh.Start()
}
```

### :left_right_arrow: Controlled transitions

[▲ Go Up ▲](#telegoflow--typed-conversations-for-telego)

Each step declares where it can go:

```go
tf.NewStep[Data]("name").
    Handle(func(ctx *tf.Context[Data]) error {
        return ctx.Go("age")
    }).
    CanGo("age")
```

If a handler tries to jump to a step that wasn't declared with `CanGo`, Telegoflow returns
`TransitionNotAllowedError`. During `Build`, Telegoflow also validates that all declared transition targets exist.

This keeps dialog graphs explicit and easier to audit.

### :chart_with_upwards_trend: Text graph

[▲ Go Up ▲](#telegoflow--typed-conversations-for-telego)

A built flow can print its transition graph for quick visual analysis:

```go
fmt.Println(registration.Graph())
```

Example output:

```text
flow registration (start: name)
name
└─> age
    ├─> email
    │   └─> confirm (terminal)
    └─> retry
        └─> age (cycle)
```

The graph is deterministic: transitions are sorted, terminal steps are marked with `(terminal)`, and loops are marked
with `(cycle)`. Steps that are not reachable from the start step are shown in an `unreachable` section.

### :floppy_disk: Session storage

[▲ Go Up ▲](#telegoflow--typed-conversations-for-telego)

Flow sessions are saved through the `Storage` interface:

```go
type Storage interface {
    LoadSession(ctx context.Context, key telegoflow.SessionKey) (*telegoflow.SessionState, bool, error)
    SaveSession(ctx context.Context, session *telegoflow.SessionState) error
    DeleteSession(ctx context.Context, key telegoflow.SessionKey) error
}
```

Telegoflow includes `NewMemoryStorage` for development, tests, and simple bots:

```go
flows := telegoflow.NewManager(telegoflow.NewMemoryStorage())
```

For production, implement `Storage` with Redis, PostgreSQL, or another persistent backend. Session data is stored as
JSON in `SessionState.Data`, so a new bot process can restore and continue the dialog as long as it registers a flow
with the same ID and step IDs.

### :hourglass_flowing_sand: Timeouts

[▲ Go Up ▲](#telegoflow--typed-conversations-for-telego)

Use `WithTimeout` to expire inactive sessions:

```go
flow, err := telegoflow.New[Data]("settings").
    Steps(...).
    WithTimeout(30 * time.Minute).
    OnTimeout(func(ctx *telegoflow.Context[Data]) error {
        _, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "Session expired"))
        return err
    }).
    Build()
```

When the next update for an expired session arrives, Telegoflow runs `OnTimeout` and deletes the session.

### :no_entry_sign: Canceling flows

[▲ Go Up ▲](#telegoflow--typed-conversations-for-telego)

Register cancel commands before `flows.Middleware()` so they can interrupt an active session:

```go
bh.Handle(func(ctx *th.Context, update telego.Update) error {
    return flows.Cancel(ctx, update)
}, th.CommandEqual("cancel"))

bh.Use(flows.Middleware())
```

A flow can also cancel itself from any handler:

```go
return ctx.Cancel()
```

### :gear: Custom session keys

[▲ Go Up ▲](#telegoflow--typed-conversations-for-telego)

By default, sessions are scoped by `chat_id:user_id`. This means the same user can have independent sessions in private
chats and groups.

You can override this behavior:

```go
flows := telegoflow.NewManager(storage, telegoflow.WithKeyFunc(func(update telego.Update) (telegoflow.SessionKey, bool) {
    if update.Message == nil || update.Message.From == nil {
        return telegoflow.SessionKey{}, false
    }
    return telegoflow.SessionKey{ChatID: update.Message.Chat.ID, UserID: update.Message.From.ID}, true
}))
```

## :bricks: Core Concepts

[▲ Go Up ▲](#telegoflow--typed-conversations-for-telego)

- **Manager** stores registered flows and routes active sessions through `telegohandler` middleware.
- **Flow** is a typed dialog graph with lifecycle hooks: `OnComplete`, `OnCancel`, `OnTimeout`, and `OnError`.
- **Step** is a state in the graph. It can have `Enter`, input `Handle` routes, `Fallback`, middleware, and `CanGo` transitions.
- **Context** embeds `*telegohandler.Context` and adds flow helpers: `Data`, `Go`, `Stay`, `Complete`, `Cancel`, `Message`, `CallbackQuery`, `ChatID`, and `UserID`.
- **Storage** persists the current step and typed data encoded as JSON.
