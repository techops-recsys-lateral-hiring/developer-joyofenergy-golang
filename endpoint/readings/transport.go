package readings

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	kitlogrus "github.com/go-kit/kit/log/logrus"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sirupsen/logrus"

	"joi-energy-golang/domain"
	mhttp "joi-energy-golang/http"
	"joi-energy-golang/http/middleware"
	"joi-energy-golang/http/serveroption"
)

// MakeStoreReadingsHandler returns a handler for the Readings service.
func MakeStoreReadingsHandler(
	s Service,
	logger *logrus.Entry,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitlogrus.NewLogrusLogger(logger))),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc(logger)),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
		makeValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeStoreReadingsEndpoint(s)),
		decodeRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

// MakeGetReadingsHandler returns a handler for the Readings service.
func MakeGetReadingsHandler(
	s Service,
	logger *logrus.Entry,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitlogrus.NewLogrusLogger(logger))),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc(logger)),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetReadingsEndpoint(s)),
		decodeSmartMeterIdFromRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.StoreReadings
	err := mhttp.DecodeRequest(ctx, r, &request)
	return request, err
}

func decodeSmartMeterIdFromRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return strings.Split(r.URL.Path, "/")[3], nil
}
