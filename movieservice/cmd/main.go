package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	proto "github.com/reynaldipane/microservice-test/movieservice/pb"
	endpoints "github.com/reynaldipane/microservice-test/movieservice/pkg/movieendpoint"
	services "github.com/reynaldipane/microservice-test/movieservice/pkg/movieservice"
	transports "github.com/reynaldipane/microservice-test/movieservice/pkg/movietransport"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	var svc services.Service

	// svc = services.MovieService{}
	svc = services.NewMovieService(http.DefaultClient, "http://www.omdbapi.com", "faf7e5bb")
	errChan := make(chan error)

	endpoints := endpoints.Set{
		FindMoviesEndpoint: endpoints.New(svc).FindMoviesEndpoint,
	}
	go func() {
		listener, err := net.Listen("tcp", ":4042")
		if err != nil {
			errChan <- err
			return
		} else {
			fmt.Println("Server started on :4042")
		}
		handler := transports.NewGRPCServer(ctx, endpoints)
		grpcServer := grpc.NewServer()
		proto.RegisterMovieServiceServer(grpcServer, handler)
		errChan <- grpcServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
