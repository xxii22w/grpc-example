package main

import (
	"context"
	"log"
	"os"

	"github.com/xxii22w/grpc-example/modlue5/internal/auth"
	"github.com/xxii22w/grpc-example/modlue5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx := context.Background()

	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		log.Fatal("JWT_SECRET is required")
	}

	authService, err := auth.NewService(jwtSecret)
	if err != nil {
		log.Fatalf("failed to initialise auth service: %v", err)
	}

	token, err := authService.IssueToken(ctx, "user-id-1234")
	if err != nil {
		log.Fatalf("failed to issue token: %v", err)
	}

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewInterceptorServiceClient(conn)

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	res, err := client.Protected(ctx, &proto.ProtectedRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("response with userID: %s",res.GetUserId())
}