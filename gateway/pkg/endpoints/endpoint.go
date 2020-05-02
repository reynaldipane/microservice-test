package endpoints

import "github.com/reynaldipane/microservice-test/movieservice/pkg/movieendpoint"

type Endpoints struct {
	MovieEndpointSet movieendpoint.Set
}

func New(movieEndpointSet movieendpoint.Set) Endpoints {
	return Endpoints{
		MovieEndpointSet: movieEndpointSet,
	}
}
