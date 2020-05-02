package movieservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/reynaldipane/microservice-test/movieservice/pkg/models"
)

type Service interface {
	FindMovies(ctx context.Context, movieRequest models.MovieRequest) (models.MovieResponse, error)
}

type MovieService struct{}

func (m MovieService) FindMovies(_ context.Context, movieRequest models.MovieRequest) (models.MovieResponse, error) {
	var movieResponse models.MovieResponse
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=faf7e5bb&s=%s&page=%s", movieRequest.Title, movieRequest.Page)

	response, err := http.Get(url)
	if err != nil {
		return movieResponse, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(responseData, &movieResponse)
	if err != nil {
		return movieResponse, err
	}

	return movieResponse, nil
}
