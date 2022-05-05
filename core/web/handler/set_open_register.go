package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/goodguy-project/goodguy-core/core/constant"
	"github.com/goodguy-project/goodguy-core/core/web/token"
	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/util/conf"
)

func OpenRegister(ctx context.Context, req *idl.OpenRegisterRequest) (*idl.OpenRegisterResponse, error) {
	sid, ok := token.Auth(ctx)
	if ok && sid == "admin" && req.GetOpenRegister() != nil {
		conf.Viper().Set(constant.OpenRegisterConfigName, req.GetOpenRegister().GetValue())
		return &idl.OpenRegisterResponse{OpenRegister: wrapperspb.Bool(req.GetOpenRegister().GetValue())}, nil
	}
	return &idl.OpenRegisterResponse{
		OpenRegister: wrapperspb.Bool(conf.Viper().GetBool(constant.OpenRegisterConfigName)),
	}, nil
}
