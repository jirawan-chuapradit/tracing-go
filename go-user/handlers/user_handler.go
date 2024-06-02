package handlers

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

var (
	name   = "user-handlers"
	tracer = otel.GetTracerProvider().Tracer(name)
	logger = otelslog.NewLogger(name)
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// get tracer provider from otel package and inject ctx to tracer
	ctx, span := tracer.Start(r.Context(), "user")
	defer span.End()

	logger.Info("get user")
	if _, err := io.WriteString(w, getUser(ctx)); err != nil {
		log.Printf("Write failed: %v\n", err)
	}

}

func getUser(ctx context.Context) string {
	_, span := tracer.Start(ctx, "get user")
	defer span.End()
	time.Sleep(3 * time.Second)
	return "Jirawan"
}
