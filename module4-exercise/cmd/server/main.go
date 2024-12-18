package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/xxii22w/grpc-example/module4-exercise/internal/stream"
	"github.com/xxii22w/grpc-example/module4-exercise/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("error running application", slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Info("closing server gracefully")
}

func run(ctx context.Context) error {
	// load the server's certificate and private key
	serverCert,err := tls.LoadX509KeyPair("certs/server.crt","certs/server.key")
	if err != nil {
		return fmt.Errorf("failed to load tls credentials: %w",err)
	}

	// load the CA's certificate to verify clients
	caCert,err := os.ReadFile("certs/ca.crt")
	if err != nil {
		return fmt.Errorf("failed to load CA certificate: %w",err)
	}

	// append the CA's certificate to the cert pool
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert){
		return errors.New("failed to append CA certificate to pool")
	}

	// create the TLS config to require client authentication
	tlsCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs: caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	})

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))

	helloService := stream.Service{}

	proto.RegisterFileUploadServiceServer(grpcServer, &helloService)

	const addr = ":50051"

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("failed to listen on address %q: %w", addr, err)
		}

		slog.Info("starting grpc server on address", slog.String("address", addr))

		if err := grpcServer.Serve(lis); err != nil {
			return fmt.Errorf("failed to serve grpc service: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		grpcServer.GracefulStop()

		return nil
	})

	return g.Wait()
}