package http

import (
	"context"

	"github.com/currency/pkg/service"
	"github.com/currency/pkg/types"
	"github.com/go-kit/kit/endpoint"
)

type getCurrencyRequest struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type createCurrencyRequest struct {
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	BaseCurrency bool   `json:"base_currency"`
}

type getCurrencysResponse struct {
	Currencys []*types.Currency `json:"currencys"`
}

type updateCurrencyRequest struct {
	ID       int            `json:"id"`
	Currency types.Currency `json:"currency"`
}

type deleteCurrencyRequest struct {
	ID int `json:"id"`
}

func makeCreateCurrencyEndpoint(svc service.CurrencyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createCurrencyRequest)
		currency := types.Currency{
			Name:          req.Name,
			Symbol:        req.Symbol,
			Base_currency: req.BaseCurrency,
		}
		err := svc.CreateCurrency(&currency)
		if err != nil {
			return nil, err
		}
		return currency, nil
	}
}
func makeGetCurrencysEndpoint(svc service.CurrencyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCurrencyRequest)
		currency, err := svc.GetCurrencys(req.ID, req.Name)
		if err != nil {
			return nil, err
		}
		return currency, nil
	}
}

func makeUpdateCurrencyEndpoint(svc service.CurrencyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateCurrencyRequest)
		err := svc.UpdateCurrency(&req.Currency)
		if err != nil {
			return nil, err
		}
		return req.Currency, nil
	}
}

func makeDeleteCurrencyEndpoint(svc service.CurrencyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteCurrencyRequest)
		err := svc.DeleteCurrency(req.ID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}
