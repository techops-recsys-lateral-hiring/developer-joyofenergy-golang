package http

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/http/contextkeys"
)

func TestDecodeJSONRequest(t *testing.T) {
	bodyValue := "Test"
	body := fmt.Sprintf("\"%s\"", bodyValue)

	ctx := context.WithValue(context.Background(), contextkeys.ContentType, MimeJSON)
	httpReq := httptest.NewRequest("POST", "https://some-host/some-path", strings.NewReader(body))

	var decodedRequest string
	err := DecodeRequest(ctx, httpReq, &decodedRequest)

	assert.NoError(t, err)
	assert.Equal(t, bodyValue, decodedRequest)
}

func TestEncodeJSONResponse(t *testing.T) {
	bodyValue := "Test"
	expected := fmt.Sprintf("\"%s\"", bodyValue)

	ctx := context.WithValue(context.Background(), contextkeys.AcceptHeader, MimeJSON)
	r := httptest.NewRecorder()

	err := EncodeResponse(ctx, r, &bodyValue)
	assert.NoError(t, err)

	actual, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)

	assert.JSONEq(t, expected, string(actual))
}
