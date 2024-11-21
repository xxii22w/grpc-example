package token

import (
	"context"

	"github.com/xxii22w/grpc-example/modlue5-exercise/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	proto.UnimplementedTokenServiceServer
}

func (s Service) Validate(ctx context.Context, _ *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	claims, ok := ctx.Value(claimKey).(map[string]string)
	if !ok {
		return nil, status.Error(codes.FailedPrecondition, "claims missing from from context")
	}

	return &proto.ValidateResponse{Claims: claims}, nil
}
