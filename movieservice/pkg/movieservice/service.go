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

type OmdbAPI struct {
	Client  *http.Client
	BaseURL string
	APIKey  string
}

type MovieService struct {
	OmdbAPI OmdbAPI
}

func NewMovieService(client *http.Client, baseURL, apiKey string) MovieService {
	return MovieService{
		OmdbAPI: OmdbAPI{
			Client:  client,
			BaseURL: baseURL,
			APIKey:  apiKey,
		},
	}
}

func (m MovieService) FindMovies(_ context.Context, movieRequest models.MovieRequest) (models.MovieResponse, error) {
	var movieResponse models.MovieResponse

	url := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%s", m.OmdbAPI.BaseURL, m.OmdbAPI.APIKey, movieRequest.Title, movieRequest.Page)

	response, err := m.OmdbAPI.Client.Get(url)
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
