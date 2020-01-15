package serveroption

import (
	"context"
	"net/http"
	"net/url"

	"joi-energy-golang/http/contextkeys"
)

// ExtractQueryIntoContext extracts query values from an http request and injects them into the provided context.
func ExtractQueryIntoContext(ctx context.Context, r *http.Request) context.Context {
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return ctx
	}
	return context.WithValue(ctx, contextkeys.QueryValues, values)
}
