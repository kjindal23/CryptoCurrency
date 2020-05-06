package cryptoerver

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	GetCurrencyEndpoint endpoint.Endpoint
	StatusEndpoint      endpoint.Endpoint
}

// MakeGetCurrencyEndpoint returns the response from our service "getcurrency"
func MakeGetCurrencyEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		symbol := req.symbol
		d, err := srv.GetCurrency(ctx, symbol)
		if err != nil {
			return getResponse{d, err.Error()}, nil
		}
		return getResponse{d, ""}, nil
	}
}

// MakeStatusEndpoint returns the response from our service "status"
func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(statusRequest) // we really just need the request, we don't use any value from it
		s, err := srv.Status(ctx)
		if err != nil {
			return statusResponse{s}, err
		}
		return statusResponse{s}, nil
	}
}

// Get endpoint mapping
func (e Endpoints) getcurrancy(ctx context.Context) (string, error) {
	req := getRequest{}
	resp, err := e.GetCurrencyEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(getResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Currency, nil
}

// Status endpoint mapping
func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := statusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	statusResp := resp.(statusResponse)
	return statusResp.Status, nil
}
