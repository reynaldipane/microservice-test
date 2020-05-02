package services

import (
	"github.com/reynaldipane/microservice-test/movieservice/pkg/movieservice"
)

type Service struct {
	MovieService movieservice.Service
}

func New(movieService movieservice.Service) Service {
	return Service{
		MovieService: movieService,
	}
}
