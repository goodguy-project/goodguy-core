package handler

import (
	"context"
	"fmt"
	"github.com/goodguy-project/goodguy-core/core/web/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/util"
)

func Login(ctx context.Context, req *idl.LoginRequest) (*idl.LoginResponse, error) {
	member := &model.Member{}
	err := model.GetDB().Where("sid = ?", req.Sid).Where("pwd = ?", util.Hashing(req.Pwd)).First(member).Error
	if err != nil {
		return new(idl.LoginResponse), status.Error(codes.Unavailable, "login failed")
	}
	fmt.Println(member)
	return &idl.LoginResponse{Token: token.SetMemberSid(member.Sid)}, nil
}
