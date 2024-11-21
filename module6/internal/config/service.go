package config

import (
	"context"
	"errors"
	"log"
	"math/rand/v2"
	"time"

	"github.com/xxii22w/grpc-example/module6/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	proto.UnimplementedConfigServiceServer
	name string
}

func NewService(name string) (*service, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	return &service{
		name: name,
	}, nil
}

func (s service) LongRunning(ctx context.Context, request *proto.LongRunningRequest) (*proto.LongRunningResponse, error) {
	select {
	case <-time.Tick(time.Second * 5):
		log.Println("finish request")
	case <-ctx.Done():
		log.Println("context done")
	}

	return &proto.LongRunningResponse{}, nil
}

func (s service) Flaky(ctx context.Context, request *proto.FlakyRequest) (*proto.FlakyResponse, error) {
	// Generate a random number between 0 and 2
	if rand.IntN(3) != 0 {
		log.Println("error response returned")
		return nil, status.Error(codes.Internal, "flaky error occurred") // Return an error 2 in 3 times
	}
	log.Println("successful response returned")

	return &proto.FlakyResponse{}, nil
}

func (s service) GetServerAddress(ctx context.Context, request *proto.GetServerAddressRequest) (*proto.GetServerAddressResponse, error) {
	log.Printf("request received on server: %s", s.name)

	return &proto.GetServerAddressResponse{
		Address: s.name,
	}, nil
}
