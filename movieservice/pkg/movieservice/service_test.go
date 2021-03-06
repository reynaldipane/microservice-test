package movieservice

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	"github.com/reynaldipane/microservice-test/movieservice/pkg/models"
)

func TestFindMovies(t *testing.T) {
	ctx := context.Background()
	movieService := NewMovieService(http.DefaultClient, "http://www.omdbapi.com", "faf7e5bb")

	movieRequest := models.MovieRequest{
		Title: "batman",
		Page:  "1",
	}
	movieResponse, err := movieService.FindMovies(ctx, movieRequest)

	if len(movieResponse.Search) == 0 {
		t.Error("Expected to at least had 1 result from search keyword, got 0")
	}

	if movieResponse.Response != "True" {
		t.Error("Got false response from search result")
	}

	totalResult, err := strconv.Atoi(movieResponse.TotalResults)
	if err != nil {
		t.Errorf("Fail to convert totalResult from response : %v", err)
	}

	if totalResult == 0 {
		t.Error("Expected to at least had 1 total result from search keyword, got 0")
	}

	if err != nil {
		t.Errorf("Error is not nil : %v", err)
	}

}
