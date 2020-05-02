package movietransport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/reynaldipane/microservice-test/movieservice/pkg/models"
	"github.com/reynaldipane/microservice-test/movieservice/pkg/movieendpoint"
)

var (
	ErrBadRequest = errors.New("request body is wrong")
)

type errorer interface {
	error() error
}

func decodeHTTPFindMoviesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	searchWord, ok := vars["searchword"]
	if !ok {
		return nil, ErrBadRequest
	}
	pagination, ok := vars["pagination"]
	if !ok {
		return nil, ErrBadRequest
	}

	return models.MovieRequest{
		Title: searchWord,
		Page:  pagination,
	}, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error})
}

func NewHTTPHandler(endpoint movieendpoint.Set, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	r := mux.NewRouter()

	r.Methods("GET").Path("/search/{searchword}/{pagination}").Handler(httptransport.NewServer(
		endpoint.FindMoviesEndpoint,
		decodeHTTPFindMoviesRequest,
		encodeHTTPGenericResponse,
		options...,
	))

	return r
}
