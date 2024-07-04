package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/currency/pkg/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(svc service.CurrencyService) http.Handler {
	router := mux.NewRouter()

	router.Methods("GET").Path("/rms/master_currencys").Handler(httptransport.NewServer(
		makeGetCurrencysEndpoint(svc),
		decodeGetCurrencyRequest,
		encodeResponse,
	))
	
	router.Methods("POST").Path("/rms/master_currencys").Handler(httptransport.NewServer(
		makeCreateCurrencyEndpoint(svc),
		decodeCreateCurrencyRequest,
		encodeResponse,
	))

	router.Methods("PUT").Path("/rms/master_currencys/{id}").Handler(httptransport.NewServer(
		makeUpdateCurrencyEndpoint(svc),
		decodeUpdateCurrencyRequest,
		encodeResponse,
	))

	router.Methods("DELETE").Path("/rms/master_currencys/{id}").Handler(httptransport.NewServer(
		makeDeleteCurrencyEndpoint(svc),
		decodeDeleteCurrencyRequest,
		encodeResponse,
	))

	return router
}

func decodeGetCurrencyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	query := r.URL.Query()
	idStr := query.Get("id")
	name := query.Get("name")

	var req getCurrencyRequest

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		req.ID = &id
	}

	if name != "" {
		req.Name = &name
	}

	return req, nil
}

func decodeUpdateCurrencyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	var req updateCurrencyRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Currency); err != nil {
		return nil, err
	}
	req.ID = id
	return req, nil
}

func decodeCreateCurrencyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createCurrencyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteCurrencyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return nil, err
	}
	return deleteCurrencyRequest{ID: id}, nil
}

func decodeEmptyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
