package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	ticket_v1 "github.com/go-godin/ticket-service/api"

	"github.com/go-godin/godin/examples/ticket/internal/endpoint"
	grpcTransport "github.com/go-godin/godin/examples/ticket/internal/transport/grpc"
	"google.golang.org/grpc"

	"github.com/go-godin/godin/examples/ticket/internal/ticket"
	"github.com/go-godin/log"
	cfg "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	"github.com/oklog/run"
)

type Configuration struct {
	GRPC struct {
		Host string
		Port int
	}
}

var (
	config Configuration
	group  run.Group
)

func main() {
	logger := log.NewLoggerFromEnv()

	// load configuration
	err := cfg.Load(
		env.NewSource(),
	)
	err = cfg.Scan(&config)
	if err != nil {
		panic(err)
	}

	// signal handling
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	defer signal.Stop(interrupt)

	// service layer
	repo := ticket.NewInMemory()
	service := ticket.NewService(repo)
	endpoints := endpoint.NewEndpointSet(service)

	// GRPC server
	{
		// set-up our grpc transport
		var addr = fmt.Sprintf("%s:%d", config.GRPC.Host, config.GRPC.Port)

		ticketGrpcServer := grpcTransport.NewTicketServiceServer(endpoints)
		googleGrpcServer := grpc.NewServer()
		grpcListener, err := net.Listen("tcp", addr)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to listen on %s", addr), "err", err)
			os.Exit(1)
		}

		ticket_v1.RegisterTicketServiceServer(googleGrpcServer, ticketGrpcServer)

		group.Add(func() error {
			logger.Info("gRPC server started", "addr", addr)
			return googleGrpcServer.Serve(grpcListener)
		}, func(e error) {
			logger.Error("GRPC server closed", "err", err)
			grpcListener.Close()
		})
	}

	// signal handler
	{
		var (
			cancelInterrupt = make(chan struct{})
			c               = make(chan os.Signal, 2)
		)
		defer close(c)

		group.Add(func() error {
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	logger.Error("exit", group.Run())
}
