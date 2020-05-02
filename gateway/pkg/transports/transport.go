package transports

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/reynaldipane/microservice-test/gateway/pkg/endpoints"
	"github.com/reynaldipane/microservice-test/movieservice/pkg/movietransport"
)

func MakeHttpHandler(endpoints endpoints.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	r.PathPrefix("/").Handler(movietransport.NewHTTPHandler(endpoints.MovieEndpointSet, logger))

	return r
}