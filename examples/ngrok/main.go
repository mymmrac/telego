package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"

	"github.com/mymmrac/telego"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Initialize signal handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Initialize done chan
	done := make(chan struct{}, 1)

	// Create a new Ngrok tunnel to connect local network with the Internet & have HTTPS domain for bot
	tun, err := ngrok.Listen(context.Background(),
		// Forward connections to localhost:8080
		config.HTTPEndpoint(config.WithForwardsTo(":8080")),
		// Authenticate into Ngrok using NGROK_AUTHTOKEN env (optional)
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Prepare fast HTTP server
	srv := &fasthttp.Server{}

	// Get an update channel from webhook using Ngrok
	updates, _ := bot.UpdatesViaWebhook("/bot"+bot.Token(),
		// Set func server with fast http server inside that will be used to handle webhooks
		telego.WithWebhookServer(telego.FuncWebhookServer{
			Server: telego.FastHTTPWebhookServer{
				Logger: bot.Logger(),
				Server: srv,
				Router: router.New(),
			},
			// Override default start func to use Ngrok tunnel
			// Note: When server is stopped, the Ngrok tunnel always returns an error, so it should be handled by user
			StartFunc: func(_ string) error {
				return srv.Serve(tun)
			},
		}),

		// Calls SetWebhook before starting webhook and provide dynamic Ngrok tunnel URL
		telego.WithWebhookSet(&telego.SetWebhookParams{
			URL: tun.URL() + "/bot" + bot.Token(),
		}),
	)

	// Handle stop signal (Ctrl+C)
	go func() {
		// Wait for stop signal
		<-sigs

		fmt.Println("Stopping...")

		// Stop reviving updates from update channel and shutdown webhook server
		_ = bot.StopWebhook()
		fmt.Println("Webhook done")

		// Notify that stop is done
		done <- struct{}{}
	}()

	// Start server for receiving requests from the Telegram
	go func() {
		_ = bot.StartWebhook("")
	}()

	// Loop through all updates when they came
	go func() {
		for update := range updates {
			fmt.Printf("Update: %+v\n", update)
		}
	}()

	// Wait for the stop process to be completed
	<-done
	fmt.Println("Done")
}
