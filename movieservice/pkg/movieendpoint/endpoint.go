package movieendpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/reynaldipane/microservice-test/movieservice/pkg/models"
	services "github.com/reynaldipane/microservice-test/movieservice/pkg/movieservice"
)

type Set struct {
	FindMoviesEndpoint endpoint.Endpoint
}

func New(svc services.Service) Set {
	findMoviesEndpoint := MakeFindMovieEndpoint(svc)
	return Set{
		FindMoviesEndpoint: findMoviesEndpoint,
	}
}

func MakeFindMovieEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.MovieRequest)
		return svc.FindMovies(ctx, req)
	}
}

func (s Set) FindMovies(ctx context.Context, movieRequest models.MovieRequest) (models.MovieResponse, error) {
	resp, err := s.FindMoviesEndpoint(ctx, movieRequest)
	if err != nil {
		return models.MovieResponse{}, err
	}

	return resp.(models.MovieResponse), nil
}
