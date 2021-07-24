package application

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/khihadysucahyo/go-echo-boilerplate/config"
	"github.com/khihadysucahyo/go-echo-boilerplate/server"
	"github.com/khihadysucahyo/go-echo-boilerplate/server/routes"
)

// Start function
func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.Sentry.Dsn,
		TracesSampleRate: cfg.Sentry.TracesSampleRate,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	defer sentry.Flush(5 * time.Second)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
