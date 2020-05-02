package movietransport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	proto "github.com/reynaldipane/microservice-test/movieservice/pb"
	"github.com/reynaldipane/microservice-test/movieservice/pkg/models"
	"github.com/reynaldipane/microservice-test/movieservice/pkg/movieendpoint"
	"github.com/reynaldipane/microservice-test/movieservice/pkg/movieservice"
	"google.golang.org/grpc"
)

type grpcServer struct {
	findMoviesServer grpctransport.Handler
}

func (s *grpcServer) FindMovies(ctx context.Context, r *proto.MovieRequest) (*proto.MovieResponse, error) {
	_, resp, err := s.findMoviesServer.ServeGRPC(ctx, r)

	if err != nil {
		return nil, err
	}

	return resp.(*proto.MovieResponse), nil
}

func NewGRPCServer(ctx context.Context, endpoint movieendpoint.Set) proto.MovieServiceServer {
	return &grpcServer{
		findMoviesServer: grpctransport.NewServer(
			endpoint.FindMoviesEndpoint,
			models.DecodeGRPCFindMoviesRequest,
			models.EncodeGRPCFindMoviesResponse,
		),
	}
}

func NewGRPCClient(conn *grpc.ClientConn) movieservice.Service {
	var findMoviesEndpoint = grpctransport.NewClient(
		conn,
		"proto.MovieService",
		"FindMovies",
		models.EncodeGRPCFindMoviesRequest,
		models.DecodeGRPCFindMoviesResponse,
		proto.MovieResponse{},
	).Endpoint()

	return movieendpoint.Set{
		FindMoviesEndpoint: findMoviesEndpoint,
	}
}
