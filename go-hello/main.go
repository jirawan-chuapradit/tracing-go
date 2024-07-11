package main

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/jirawan-chuapradit/tracing-go/handlers"
	"github.com/jirawan-chuapradit/tracing-go/otel"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// load ssl/tls certificates
	log.Printf("load ssl/tls certificates")
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Printf("cert %v", cert)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	otel.InitOpenTelemetry(ctx)
	port := "443"
	// Start HTTP server.
	srv := &http.Server{
		Addr:         ":" + port,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
		TLSConfig:    tlsConfig,
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServeTLS("", "")
	}()

	log.Debug().Msg("starting server at http://localhost:" + port)

	// Wait for interruption.
	select {
	case err := <-srvErr:
		// Error when starting HTTP server.
		log.Fatal().Err(err)
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal().Err(err)
	}
}

func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handlers.HelloHandler)
	handler := otelhttp.NewHandler(mux, "go-hello")
	return handler
}
