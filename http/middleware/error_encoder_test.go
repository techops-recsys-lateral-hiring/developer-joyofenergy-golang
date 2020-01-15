package middleware

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
	mhttp "joi-energy-golang/http"
	"joi-energy-golang/http/contextkeys"
)

func TestEncodeError(t *testing.T) {
	testCases := []struct {
		name               string
		err                error
		expectedOutput     string
		expectedStatusCode int
	}{
		{
			name:               "invalid content type",
			err:                mhttp.ErrInvalidContentType,
			expectedOutput:     `{
  "errorMessage": "invalid content-type"
}`,
			expectedStatusCode: http.StatusUnsupportedMediaType,
		},
		{
			name:               "missing argument",
			err:                domain.ErrMissingArgument,
			expectedOutput:     `{
  "errorMessage": "missing argument"
}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "invalid message type",
			err:                domain.ErrInvalidMessageType,
			expectedOutput:     `{
  "errorMessage": "invalid message-type"
}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "invalid accept header",
			err:                mhttp.ErrInvalidAcceptHeader,
			expectedOutput:     `{
  "errorMessage": "invalid accept header"
}`,
			expectedStatusCode: http.StatusNotAcceptable,
		},
		{
			name:               "not found",
			err:                domain.ErrNotFound,
			expectedOutput:     `{
  "errorMessage": "not found"
}`,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "unknown error",
			err:                errors.New("👻 I see dead code"),
			expectedOutput:     `{
  "errorMessage": "👻 I see dead code"
}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRecorder()
			ctx := context.WithValue(context.Background(), contextkeys.AcceptHeader, "application/json")
			MakeEncodeErrorFunc(logrus.NewEntry(logrus.StandardLogger()))(ctx, tt.err, r)

			resp := r.Result()
			defer resp.Body.Close()

			actual, _ := ioutil.ReadAll(resp.Body)
			assert.Equal(t, tt.expectedOutput, string(actual))
			assert.Equal(t, tt.expectedStatusCode, r.Code)
		})
	}
}

func TestEncodeErrorWithInvalidContentType(t *testing.T) {
	l := logrus.New()
	h := test.NewLocal(l)

	r := httptest.NewRecorder()
	ctx := context.WithValue(context.Background(), contextkeys.AcceptHeader, "application/x-download")
	MakeEncodeErrorFunc(logrus.NewEntry(l))(ctx, domain.ErrMissingArgument, r)

	resp := r.Result()
	defer resp.Body.Close()

	actual, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "", string(actual))
	assert.Equal(t, 400, r.Code)

	assert.Equal(t, "application/x-download", h.LastEntry().Data["contentType"])
}
