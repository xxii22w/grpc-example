package main

import (
	"log"
	"net"

	"github.com/xxii22w/grpc-example/module2-todo/internal/todo"
	"github.com/xxii22w/grpc-example/module2-todo/proto"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	todoServer := todo.NewService()

	proto.RegisterTodoServiceServer(grpcServer,todoServer)

	listen,err := net.Listen("tcp",":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server running at address: %s",":50051")

	if err := grpcServer.Serve(listen);err != nil {
		log.Fatal(err)
	}
}