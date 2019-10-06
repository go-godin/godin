package grpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Config struct {
	Port                string        `mapstructure:"port"`
	PortMetrics         int           `mapstructure:"port-metrics"`
	Hostname            string        `mapstructure:"hostname"`
	GrpcShutdownTimeout time.Duration `mapstructure:"grpc-server-shutdown-timeout"`
	HttpShutdownTimeout time.Duration `mapstructure:"http-metrics-shutdown-timeout"`
}

type Server struct {
	logger     *zap.Logger
	config     *Config
	GoogleGrpc *grpc.Server
	listener   net.Listener
}

var (
	healthy int32
	ready   int32
)

func NewServer(config *Config, logger *zap.Logger) *Server {
	var err error

	srv := &Server{
		logger: logger,
		config: config,
	}

	// setup plain gRPC server
	srv.GoogleGrpc = grpc.NewServer()
	srv.listener, err = net.Listen("tcp", fmt.Sprintf(":%v", srv.config.Port))
	if err != nil {
		srv.logger.Fatal("failed to listen on port", zap.Error(err))
	}

	return srv
}

func (s *Server) setupMetricsServer() *http.Server {
		mux := http.DefaultServeMux
		mux.Handle("/metrics", promhttp.Handler())
		mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("OK"))
		})

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%v", s.config.PortMetrics),
			Handler: mux,
		}

		return srv
}

func (s *Server) ListenAndServe(stopCh <-chan struct{}) {
	var metricsServer *http.Server
	if s.config.PortMetrics > 0 {
		metricsServer = s.setupMetricsServer()
		go func() {
			s.logger.Info("starting HTTP metrics server", zap.Int("port", s.config.PortMetrics))
			if err := metricsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				s.logger.Fatal("metrics server failed", zap.Error(err))
			}
		}()
	}

	// run server as goroutine
	go func() {
		if err := s.GoogleGrpc.Serve(s.listener); err != nil {
			s.logger.Fatal("gRPC server crashed", zap.Error(err))
		}
	}()

	// signal kubernetes that the server is ready to receive traffic
	atomic.StoreInt32(&healthy, 1)
	atomic.StoreInt32(&ready, 1)

	// wait for signal handler to fire
	<-stopCh
	s.logger.Info("shutting down gRPC server", zap.Duration("timeout", s.config.GrpcShutdownTimeout))
	if metricsServer != nil {
		s.logger.Info("shutting down HTTP metrics server", zap.Duration("timeout", s.config.HttpShutdownTimeout))
	}

	// health checks fail from now on
	atomic.StoreInt32(&healthy, 0)
	atomic.StoreInt32(&ready, 0)

	// HTTP metrics graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), s.config.HttpShutdownTimeout)
	defer cancel()

	if err := metricsServer.Shutdown(ctx); err != nil {
		s.logger.Warn("HTTP metrics graceful shutdown failed", zap.Error(err))
	} else {
		s.logger.Info("HTTP metrics server stopped")
	}

	// gRPC server graceful shutdown
	stopped := make(chan struct{})
	go func() {
		s.GoogleGrpc.GracefulStop()
		close(stopped)
	}()
	t := time.NewTicker(s.config.GrpcShutdownTimeout * time.Second)
	select {
	case <-t.C:
		s.logger.Warn("gRPC server graceful shutdown timed-out")
	case <-stopped:
		s.logger.Info("gRPC server stopped")
		t.Stop()
	}
}
