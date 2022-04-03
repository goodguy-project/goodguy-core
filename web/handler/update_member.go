package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/web/token"
)

func UpdateMember(ctx context.Context, req *idl.UpdateMemberRequest) (*idl.UpdateMemberResponse, error) {
	if req.Member == nil {
		return new(idl.UpdateMemberResponse), status.Error(codes.Unauthenticated, "auth failed")
	}
	sid, ok := token.Auth(ctx)
	if !ok {
		return new(idl.UpdateMemberResponse), status.Error(codes.Unauthenticated, "auth failed")
	}
	db := model.GetDB()
	member := &model.Member{}
	err := db.Model(&model.Member{}).Where("sid = ?", sid).First(member).Error
	if err != nil {
		return new(idl.UpdateMemberResponse), status.Error(codes.Internal, "database error")
	}
	doUpdateMember(req.Member, member)
	db = model.GetDB()
	db.Save(member)
	return new(idl.UpdateMemberResponse), nil
}

func doUpdateMember(reqMember *idl.Member, member *model.Member) {
	if reqMember.Name != "" {
		member.Name = reqMember.Name
	}
	if reqMember.School != "" {
		member.School = reqMember.School
	}
	if reqMember.Grade != 0 {
		member.Grade = reqMember.Grade
	}
	if reqMember.Clazz != "" {
		member.Clazz = reqMember.Clazz
	}
	if reqMember.CodeforcesId != "" {
		member.CodeforcesId = reqMember.CodeforcesId
	}
	if reqMember.AtcoderId != "" {
		member.AtcoderId = reqMember.AtcoderId
	}
	if reqMember.CodechefId != "" {
		member.CodechefId = reqMember.CodechefId
	}
	if reqMember.NowcoderId != "" {
		member.NowcoderId = reqMember.NowcoderId
	}
	if reqMember.VjudgeId != "" {
		member.VjudgeId = reqMember.VjudgeId
	}
	if reqMember.LeetcodeId != "" {
		member.LeetcodeId = reqMember.LeetcodeId
	}
	if reqMember.LuoguId != "" {
		member.LuoguId = reqMember.LuoguId
	}
	if reqMember.Email != "" {
		member.Email = reqMember.Email
	}
	if reqMember.IsSubscribe != idl.Bool_Bool_Undefined {
		member.IsSubscribe = reqMember.IsSubscribe == idl.Bool_Bool_True
	}
}
