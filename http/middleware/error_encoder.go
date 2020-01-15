package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sirupsen/logrus"

	"joi-energy-golang/domain"
	mhttp "joi-energy-golang/http"
	"joi-energy-golang/http/contextkeys"
)

// encode errors from business-logic
func MakeEncodeErrorFunc(logger *logrus.Entry) kithttp.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		contentType := ctx.Value(contextkeys.AcceptHeader).(string)
		if strings.Contains(contentType, "text/html") {
			err = sendBrowserDoc(w)
			if err == nil {
				return
			}
		}

		w.Header().Set("Content-Type", contentType)

		message := &domain.Error{}
		message.ErrorMessage = toErrorMessage(err)
		w.WriteHeader(toHTTPStatusCode(err))

		if contentType != mhttp.MimeJSON {
			logger.WithContext(ctx).WithField("contentType", contentType).Error("contentType unknown. Please use JSON.")
			return
		}
		data, err := json.MarshalIndent(message, "", "  ")
		if err != nil {
			logger.WithContext(ctx).WithError(err).Error("failed marshalling error message to JSON.")
		}
		_, err = w.Write(data)
		if err != nil {
			logger.WithContext(ctx).WithError(err).Error("failed writing error message.")
		}
	}
}

func toErrorMessage(err error) string {
	return err.Error()
}

func toHTTPStatusCode(err error) int {
	switch {
	case errors.Is(err, mhttp.ErrInvalidContentType):
		return http.StatusUnsupportedMediaType
	case errors.Is(err, mhttp.ErrInvalidAcceptHeader):
		return http.StatusNotAcceptable
	case errors.Is(err, domain.ErrMissingArgument):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrInvalidMessageType):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func sendBrowserDoc(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusUnsupportedMediaType)
	b, err := ioutil.ReadFile("browser.htm")
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
