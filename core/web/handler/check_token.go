package handler

import (
	"context"
	"github.com/goodguy-project/goodguy-core/core/web/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/goodguy-project/goodguy-core/idl"
)

func CheckToken(ctx context.Context, req *idl.CheckTokenRequest) (*idl.CheckTokenResponse, error) {
	sid, ok := token.GetMemberSid(req.Token)
	if !ok {
		return new(idl.CheckTokenResponse), status.Error(codes.Unauthenticated, "token is expired")
	}
	return &idl.CheckTokenResponse{
		Ok:  ok,
		Sid: sid,
	}, nil
}
