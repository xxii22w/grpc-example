package main

import (
	"context"
	"log"

	"github.com/xxii22w/grpc-example/module2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // 意味着客户端和服务端之间的连接是纯文本
		grpc.WithBlock(), // 意味着拨号功能在此阻塞，知道客户端和服务端连接成功
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: ""})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Fatalf("status code: %s, error: %s", status.Code().String(), status.Message())
		}
		log.Fatal(err)
	}
	log.Printf("response received: %s", res.Message)
}
