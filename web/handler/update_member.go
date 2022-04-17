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
	member.Name = reqMember.Name
	member.School = reqMember.School
	member.Grade = reqMember.Grade
	member.Clazz = reqMember.Clazz
	member.CodeforcesId = reqMember.CodeforcesId
	member.AtcoderId = reqMember.AtcoderId
	member.CodechefId = reqMember.CodechefId
	member.NowcoderId = reqMember.NowcoderId
	member.VjudgeId = reqMember.VjudgeId
	member.LeetcodeId = reqMember.LeetcodeId
	member.LuoguId = reqMember.LuoguId
	member.Email = reqMember.Email
	member.IsSubscribe = reqMember.IsSubscribe == idl.Bool_Bool_True
}
