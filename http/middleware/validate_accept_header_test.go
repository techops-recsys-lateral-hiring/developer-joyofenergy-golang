package middleware

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/http/contextkeys"
)

func TestMakeAcceptHeaderValidationMiddlewareSucceedsWithJSON(t *testing.T) {
	m := MakeAcceptHeaderValidationMiddleware()
	ep := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return "", nil
	}

	ctx := context.WithValue(context.Background(), contextkeys.AcceptHeader, "application/json")
	_, err := m(ep)(ctx, "")
	assert.NoError(t, err)
}

func TestMakeAcceptHeaderValidationMiddlewareFailsForInvalidHeaders(t *testing.T) {
	negativeTestCases := []struct {
		name         string
		acceptHeader string
		expected     string
	}{
		{
			name:         "empty header",
			acceptHeader: "",
			expected:     "invalid accept header: '' not allowed",
		},
		{
			name:         "text",
			acceptHeader: "text/plain",
			expected:     "invalid accept header: 'text/plain' not allowed",
		},
		{
			name:         "wildcard",
			acceptHeader: "*/*",
			expected:     "invalid accept header: '*/*' not allowed",
		},
		{
			name:         "category wildcard",
			acceptHeader: "application/*",
			expected:     "invalid accept header: 'application/*' not allowed",
		},
	}

	for _, tt := range negativeTestCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := MakeAcceptHeaderValidationMiddleware()
			ep := func(ctx context.Context, request interface{}) (response interface{}, err error) {
				return "", nil
			}

			ctx := context.WithValue(context.Background(), contextkeys.AcceptHeader, tt.acceptHeader)
			_, err := m(ep)(ctx, "")
			assert.EqualError(t, err, tt.expected)
		})
	}
}
