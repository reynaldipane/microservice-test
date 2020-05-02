package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	conf "github.com/reynaldipane/microservice-test/gateway/pkg/configs"
	"github.com/reynaldipane/microservice-test/gateway/pkg/endpoints"
	"github.com/reynaldipane/microservice-test/gateway/pkg/services"
	"github.com/reynaldipane/microservice-test/gateway/pkg/transports"
	"github.com/reynaldipane/microservice-test/movieservice/pkg/movieendpoint"
	"github.com/reynaldipane/microservice-test/movieservice/pkg/movietransport"
	"google.golang.org/grpc"
)

func main() {
	conf := conf.Config{}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	grpcConn := initGRPCClient(conf.Get("MOVIE_SERVICE_GRPC_ADDRESS"))
	services := services.New(movietransport.NewGRPCClient(grpcConn))
	endpoints := endpoints.New(movieendpoint.New(services.MovieService))

	router := transports.MakeHttpHandler(endpoints, logger)

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Get("PORT")), router)
		if err != nil {
			panic(err)
		}
		errc <- err
	}()

	fmt.Println(<-errc)
}

func initGRPCClient(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return conn
}
