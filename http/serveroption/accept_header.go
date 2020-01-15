package serveroption

import (
	"context"
	"net/http"

	mhttp "joi-energy-golang/http"
	"joi-energy-golang/http/contextkeys"
	"joi-energy-golang/http/header"
)

// ExtractAcceptHeaderIntoContext extracts content type from an http request and injects it into the provided context.
func ExtractAcceptHeaderIntoContext(ctx context.Context, r *http.Request) context.Context {
	if acceptHeaderValue := r.Header.Get(header.Accept); acceptHeaderValue != "" {
		return context.WithValue(ctx, contextkeys.AcceptHeader, acceptHeaderValue)
	}
	if acceptHeaderValue := r.Header.Get(header.ContentType); acceptHeaderValue != "" {
		return context.WithValue(ctx, contextkeys.AcceptHeader, acceptHeaderValue)
	}
	acceptHeaderValue := mhttp.MimeJSON
	return context.WithValue(ctx, contextkeys.AcceptHeader, acceptHeaderValue)
}
