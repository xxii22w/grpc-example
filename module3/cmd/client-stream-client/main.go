package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/xxii22w/grpc-example/module3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// initialize a grpc connection
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create out client
	client := proto.NewStreamingServiceClient(conn)

	// initialize the client stream
	stream, err := client.LogStream(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// send some log message
	for i := range 5 {
		req := &proto.LogStreamRequest{
			Timestamp: timestamppb.New(time.Now()),
			Level:     proto.LogLevel_LOG_LEVEL_INFO,
			Message:   fmt.Sprintf("Hello log: %s", i),
		}
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}

	// close the stream
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	// log the response from server
	log.Printf("number of sent: %d", res.GetEntriesLogged())
}
