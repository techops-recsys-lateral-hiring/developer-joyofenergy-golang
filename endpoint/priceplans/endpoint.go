package priceplans

import (
	"context"
	"net/url"
	"strconv"

	"github.com/go-kit/kit/endpoint"

	"joi-energy-golang/http/contextkeys"
)

func makeCompareAllPricePlansEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		err := validateSmartMeterId(req)
		if err != nil {
			return nil, err
		}
		r, err := s.CompareAllPricePlans(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

func makeRecommendPricePlansEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		limitString := ctx.Value(contextkeys.QueryValues).(url.Values).Get("limit")
		limit, err := strconv.ParseUint(limitString, 10, 64)
		if limitString != "" && err != nil {
			return nil, err
		}
		req := request.(string)
		err = validateSmartMeterId(req)
		if err != nil {
			return nil, err
		}
		r, err := s.RecommendPricePlans(req, limit)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}
