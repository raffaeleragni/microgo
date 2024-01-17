package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()

	engine := gin.Default()
	setupSentry(engine)

	engine.GET("/status", func(c *gin.Context) {
		c.Status(200)
	})

	NewApp().ApiSampleSetup(engine)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}
	engine.Run(":" + port)
}

func setupSentry(r *gin.Engine) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_URL"),
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	r.Use(sentrygin.New(sentrygin.Options{}))
}
