package handlers

import (
	"context"
	"io"
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

var (
	name   = "hello-handlers"
	tracer = otel.GetTracerProvider().Tracer(name)
	logger = otelslog.NewLogger(name)
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// get tracer provider from otel package and inject ctx to tracer
	ctx, span := tracer.Start(r.Context(), "hello")
	defer span.End()

	logger.Info("call say hello")
	if _, err := io.WriteString(w, sayHello(ctx)); err != nil {
		log.Printf("Write failed: %v\n", err)
	}

}

func sayHello(ctx context.Context) string {
	_, span := tracer.Start(ctx, "say hello")
	defer span.End()
	return "Hello"
}
