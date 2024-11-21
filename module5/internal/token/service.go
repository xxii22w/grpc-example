package token

import (
	"context"
	"errors"

	"github.com/xxii22w/grpc-example/modlue5/proto"
	"google.golang.org/grpc/credentials"
)

type TokenIssure interface {
	IssueToken(ctx context.Context, userID string) (string, error)
}

type jwtCredentials struct {
	tokenIssure TokenIssure
}

func NewJWTCredentials(tokenIssure TokenIssure) (*jwtCredentials, error) {
	if tokenIssure == nil {
		return nil, errors.New("token issure must not be nil")
	}
	return &jwtCredentials{tokenIssure: tokenIssure}, nil
}

func (j jwtCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	info, ok := credentials.RequestInfoFromContext(ctx)
	if !ok || info.Method != proto.InterceptorService_Protected_FullMethodName {
		return nil, nil
	}

	token, err := j.tokenIssure.IssueToken(ctx, "user-id-12345")
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"authorization": token,
	}, nil
}

func (j jwtCredentials) RequireTransportSecurity() bool {
	return false
}
