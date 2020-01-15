package serveroption

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/http/contextkeys"
)

func TestExtractAcceptHeaderIntoContext_ExtractsAcceptHeaderValueIntoContext(t *testing.T) {
	req := httptest.NewRequest("", "http://test", nil)
	req.Header.Set("Accept", "application/x-download")

	ctx := context.Background()
	textCtx := ExtractAcceptHeaderIntoContext(ctx, req)

	actual := textCtx.Value(contextkeys.AcceptHeader).(string)
	expected := "application/x-download"

	assert.Equal(t, expected, actual)
}

func TestExtractAcceptHeaderIntoContext_ExtractsContentTypeHeaderValueIntoContextWhenAcceptHeaderEmpty(t *testing.T) {
	req := httptest.NewRequest("", "http://test", nil)
	req.Header.Set("Content-Type", "application/important")

	ctx := context.Background()
	newContext := ExtractAcceptHeaderIntoContext(ctx, req)

	actual := newContext.Value(contextkeys.AcceptHeader).(string)
	expected := "application/important"

	assert.Equal(t, expected, actual)
}

func TestExtractAcceptHeaderIntoContext_DefaultsToJSONWhenContentTypeAndAcceptHeaderEmpty(t *testing.T) {
	req := httptest.NewRequest("", "http://test", nil)

	ctx := context.Background()
	newContext := ExtractAcceptHeaderIntoContext(ctx, req)

	actual := newContext.Value(contextkeys.AcceptHeader).(string)
	expected := "application/json"

	assert.Equal(t, expected, actual)
}
