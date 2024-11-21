package main

import (
	"context"
	"io"
	"log"

	"github.com/xxii22w/grpc-example/module3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	// first initialize grpc connection
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create the client
	client := proto.NewStreamingServiceClient(conn)

	// init the stream
	stream, err := client.StreamServerTime(ctx, &proto.StreamServerTimeRequest{
		IntervalSeconds: 2,
	})
	if err != nil {
		log.Fatal(err)
	}

	// loop through all the response we get back from the server
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		log.Printf("received time from server: %s", res.CurrentTime.AsTime())
	}
	log.Println("server stream closed")
}
