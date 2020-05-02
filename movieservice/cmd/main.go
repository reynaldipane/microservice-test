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
	conf "github.com/reynaldipane/microservice-test/movieservice/pkg/configs"
	endpoints "github.com/reynaldipane/microservice-test/movieservice/pkg/movieendpoint"
	services "github.com/reynaldipane/microservice-test/movieservice/pkg/movieservice"
	transports "github.com/reynaldipane/microservice-test/movieservice/pkg/movietransport"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conf := conf.Config{}

	var svc services.Service

	svc = services.NewMovieService(http.DefaultClient, conf.Get("OMDB_API_BASE_URL"), conf.Get("OMDB_API_KEY"))
	errChan := make(chan error)

	endpoints := endpoints.Set{
		FindMoviesEndpoint: endpoints.New(svc).FindMoviesEndpoint,
	}
	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Get("PORT")))
		if err != nil {
			errChan <- err
			return
		} else {
			fmt.Printf("Server started on :%s", conf.Get("PORT"))
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
