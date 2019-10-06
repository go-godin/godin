package main

import (
	"fmt"
	"github.com/go-godin/godin/examples/ticket/internal/endpoint"
	"github.com/go-godin/godin/examples/ticket/internal/ticket"
	grpc2 "github.com/go-godin/godin/examples/ticket/internal/transport/grpc"
	ticket_v1 "github.com/go-godin/ticket-service/api"
	"os"
	"strings"
	"time"

	"github.com/go-godin/godin/examples/ticket/pkg/grpc"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/go-godin/godin/examples/ticket/pkg/logging"
	"github.com/go-godin/godin/examples/ticket/pkg/signals"
)

const EnvPrefix = "TICKET" // Prefix to use for all flags if configured via environment variables

func main() {
	fs := pflag.NewFlagSet("server", pflag.ContinueOnError)
	fs.Int("port", 50051, "gRPC port")
	fs.Int("port-metrics", 3000, "HTTP metrics server port")
	fs.Duration("grpc-server-shutdown-timeout", 5*time.Second, "gRPC server graceful shutdown duration")
	fs.Duration("http-metrics-shutdown-timeout", 3*time.Second, "HTTP metrics erver graceful shutdown duration")
	fs.String("level", "info", "log level debug, info, warn, error, flat or panic")

	// parse and bind flags to env variables
	parseFlags(fs)
	bindFlags(fs)

	// configure logging
	logger, _ := logging.SetupZapLogger(viper.GetString("level"))
	defer logger.Sync()
	stdLog := zap.RedirectStdLog(logger)
	defer stdLog()

	// load gRPC config
	var grpcConfig grpc.Config
	if err := viper.Unmarshal(&grpcConfig); err != nil {
		logger.Panic("config unmarshal failed", zap.Error(err))
	}

	// setup service layer
	repo := ticket.NewInMemory()
	service := ticket.NewService(repo, logger)
	endpoints := endpoint.NewEndpointSet(service)
	ticket := grpc2.NewTicketServiceServer(endpoints)

	// start gRPC server
	logger.Info("starting ticket gRPC server",
		zap.String("version", "1.0.0"),
		zap.String("port", viper.GetString("port")),
	)
	stopCh := signals.SetupSignalHandler()
	srv := grpc.NewServer(&grpcConfig, logger)
	ticket_v1.RegisterTicketServiceServer(srv.GoogleGrpc, ticket)
	srv.ListenAndServe(stopCh)
}


func parseFlags(fs *pflag.FlagSet) {
	err := fs.Parse(os.Args[1:])
	switch {
	case err == pflag.ErrHelp:
		os.Exit(0)
	case err != nil:
		_, _ = fmt.Fprint(os.Stderr, "Error: %s\n\n", err.Error())
		fs.PrintDefaults()
		os.Exit(2)
	}
}

func bindFlags(fs *pflag.FlagSet) {
	_ = viper.BindPFlags(fs)
	hostname, _ := os.Hostname()
	viper.Set("hostname", hostname)
	viper.SetEnvPrefix(EnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}
