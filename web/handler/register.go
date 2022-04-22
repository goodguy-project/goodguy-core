package handler

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/util"
	"github.com/goodguy-project/goodguy-core/web/token"
)

func Register(ctx context.Context, req *idl.RegisterRequest) (resp *idl.RegisterResponse, err error) {
	defer func() {
		if resp == nil {
			resp = new(idl.RegisterResponse)
		}
	}()
	if !viper.GetBool(util.OpenRegisterConfigName) {
		return nil, status.Error(codes.Unavailable, "register is unavailable")
	}
	if req.Member == nil {
		return nil, status.Error(codes.DataLoss, "member is empty")
	}
	if req.GetMember().GetSid().GetValue() == "" {
		return nil, status.Error(codes.DataLoss, "sid is empty")
	}
	if len(req.GetMember().GetSid().GetValue()) > 50 {
		return nil, status.Error(codes.Unavailable, fmt.Sprintf("sid %s is too long", req.GetMember().Sid))
	}
	if req.Pwd == "" {
		return nil, status.Error(codes.DataLoss, "pwd is empty")
	}
	member := req.GetMember()
	err = model.GetDB().Create(&model.Member{
		Sid:          member.GetSid().GetValue(),
		Name:         member.GetName().GetValue(),
		School:       member.GetSchool().GetValue(),
		Grade:        member.GetGrade().GetValue(),
		Clazz:        member.GetClazz().GetValue(),
		IsOfficial:   false,
		CodeforcesId: member.GetCodeforcesId().GetValue(),
		AtcoderId:    member.GetAtcoderId().GetValue(),
		CodechefId:   member.GetCodechefId().GetValue(),
		NowcoderId:   member.GetNowcoderId().GetValue(),
		VjudgeId:     member.GetVjudgeId().GetValue(),
		LeetcodeId:   member.GetLeetcodeId().GetValue(),
		LuoguId:      member.GetLuoguId().GetValue(),
		Email:        member.GetEmail().GetValue(),
		IsSubscribe:  false,
		IsAdmin:      false,
		Pwd:          util.Hashing(req.Pwd),
	}).Error
	if err != nil {
		return nil, status.Error(codes.Unavailable, fmt.Sprintf("sid %s exists", req.GetMember().Sid))
	}
	return &idl.RegisterResponse{Token: token.SetMemberSid(member.GetSid().GetValue())}, nil
}
