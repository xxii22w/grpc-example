package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/xxii22w/grpc-example/module2/internal/hello"
	"github.com/xxii22w/grpc-example/module2/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
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
	grpcServer := grpc.NewServer()
	helloService := hello.Service{}
	// 注册grpc服务
	proto.RegisterHelloServiceServer(grpcServer, helloService)

	const addr = ":50051"

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		// 监听端口
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("failed to listen on address %q: %w", addr, err)
		}

		slog.Info("starting grpc server on address", slog.String("address", addr))

		// grpcServer.Serve 会阻塞
		if err := grpcServer.Serve(listen); err != nil {
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
