package main

import (
	"log"
	"net"
	"os"

	"github.com/xxii22w/grpc-example/modlue5/internal/auth"
	"github.com/xxii22w/grpc-example/modlue5/internal/interceptor"
	"github.com/xxii22w/grpc-example/modlue5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		log.Fatal("JWT_SECRET is required")
	}

	authService, err := auth.NewService(jwtSecret)
	if err != nil {
		log.Fatal(err)
	}

	middleware, err := interceptor.NewMiddleware(authService)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryAuthMiddleware),
	)

	interceptorService := interceptor.Service{}

	proto.RegisterInterceptorServiceServer(grpcServer, interceptorService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting grpc server on address %s", ":50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
