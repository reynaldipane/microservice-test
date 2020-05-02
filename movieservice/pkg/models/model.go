package models

import (
	"context"

	proto "github.com/reynaldipane/microservice-test/movieservice/pb"
)

type Search struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieResponse struct {
	Search       []Search `json:"Search"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
}

type MovieRequest struct {
	Title string
	Page  string
}

func DecodeGRPCFindMoviesRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.MovieRequest)
	return MovieRequest{
		Title: req.Title,
		Page:  req.Page,
	}, nil
}

func EncodeGRPCFindMoviesResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(MovieResponse)

	search := []*proto.Search{}
	for _, v := range resp.Search {
		s := &proto.Search{
			Title:  v.Title,
			Year:   v.Year,
			ImdbID: v.ImdbID,
			Type:   v.Type,
			Poster: v.Poster,
		}
		search = append(search, s)
	}

	return &proto.MovieResponse{
		Search:       search,
		TotalResults: resp.TotalResults,
		Response:     resp.Response,
	}, nil
}

func EncodeGRPCFindMoviesRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(MovieRequest)
	return &proto.MovieRequest{
		Title: req.Title,
		Page:  req.Page,
	}, nil
}

func DecodeGRPCFindMoviesResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*proto.MovieResponse)

	search := []Search{}
	for _, v := range resp.Search {
		s := Search{
			Title:  v.Title,
			Year:   v.Year,
			ImdbID: v.ImdbID,
			Type:   v.Type,
			Poster: v.Poster,
		}
		search = append(search, s)
	}

	return MovieResponse{
		Search:       search,
		TotalResults: resp.TotalResults,
		Response:     resp.Response,
	}, nil
}
