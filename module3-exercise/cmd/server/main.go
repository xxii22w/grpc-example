package main

import (
	"log"
	"net"

	"github.com/xxii22w/grpc-example/module3-exercise/internal/stream"
	"github.com/xxii22w/grpc-example/module3-exercise/proto"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	streamingService := &stream.Service{}

	proto.RegisterFileUploadServiceServer(grpcServer,streamingService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting grpc server on address: %s", ":50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
