package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/jirawan-chuapradit/tracing-go/otel"
)

func main() {

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	otel.InitOpenTelemetry(ctx)
	log.Println("hello")

}
