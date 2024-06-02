package service

import (
	"context"
	"io"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var (
	name   = "user-service"
	tracer = otel.GetTracerProvider().Tracer(name)
)

func GetUser(ctx context.Context) (string, error) {
	ctx, span := tracer.Start(ctx, "getUser")
	defer span.End()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://go_user:8000/user", nil)
	if err != nil {
		// if error should set status to span
		span.RecordError(err)
		span.SetStatus(codes.Error, codes.Error.String())

		return "", err
	}
	// use client to call other service
	client := &http.Client{
		Transport: otelhttp.NewTransport(nil),
	}
	resp, err := client.Do(request)
	if err != nil {
		// if error should set status to span
		span.RecordError(err)
		span.SetStatus(codes.Error, codes.Error.String())

		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
