package internal

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	faucetpb "github.com/Liangyu-Zhou/registry-demo/proto/faucet"
	gateway "github.com/rauljordan/minimal-grpc-gateway"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	log = logrus.WithField("prefix", "server")
)

// Config for the faucet server.
type Config struct {
	GrpcPort       int      `mapstructure:"grpc-port"`
	GrpcHost       string   `mapstructure:"grpc-host"`
	HttpPort       int      `mapstructure:"http-port"`
	HttpHost       string   `mapstructure:"http-host"`
	AllowedOrigins []string `mapstructure:"allowed-origins"`
}

// Server capable of funding requests for faucet ETH via gRPC and REST HTTP.
type Server struct {
	faucetpb.UnimplementedFaucetServer
	cfg *Config
}

// NewServer initializes the server from configuration values.
func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		cfg: cfg,
	}, nil
}

// Start a faucet server by serving a gRPC connection, an http JSON server, and a rate limiter.
func (s *Server) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Initialize and register gRPC handlers.
	grpcServer := s.initializeGRPCServer()

	grpcAddress := fmt.Sprintf("%s:%d", s.cfg.GrpcHost, s.cfg.GrpcPort)
	// Start a gRPC server.
	go func() {
		log.Infof("Starting gRPC server %s", grpcAddress)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GrpcPort))
		if err != nil {
			log.WithError(err).Fatalf("Could not listen on port %d", s.cfg.GrpcPort)
		}
		if err := grpcServer.Serve(lis); err != nil {
			log.WithError(err).Fatal("Stopped server")
		}
	}()

	// Start a gRPC Gateway to serve http JSON requests.
	gatewayAddress := fmt.Sprintf("%s:%d", s.cfg.HttpHost, s.cfg.HttpPort)
	gatewaySrv := gateway.New(ctx, &gateway.Config{
		GatewayAddress:      gatewayAddress,
		RemoteAddress:       grpcAddress,
		AllowedOrigins:      s.cfg.AllowedOrigins,
		EndpointsToRegister: []gateway.RegistrationFunc{faucetpb.RegisterFaucetHandlerFromEndpoint},
	})
	log.Infof("Starting JSON http server %s", gatewayAddress)
	gatewaySrv.Start()

	// Listen for any process interrupts.
	stop := make(chan struct{})
	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		<-sigc
		logrus.Info("Got interrupt, shutting down...")
		grpcServer.GracefulStop()
		stop <- struct{}{}
	}()

	// Wait for stop channel to be closed.
	<-stop
}

// Initialize a gRPC server and register handlers.
func (s *Server) initializeGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	faucetpb.RegisterFaucetServer(grpcServer, s)
	reflection.Register(grpcServer)
	return grpcServer
}
