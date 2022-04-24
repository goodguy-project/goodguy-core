package handler

import (
	"context"
	"github.com/goodguy-project/goodguy-core/core/web/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
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
	if reqMember.Name != nil {
		member.Name = reqMember.Name.GetValue()
	}
	if reqMember.School != nil {
		member.School = reqMember.School.GetValue()
	}
	if reqMember.Grade != nil {
		member.Grade = reqMember.Grade.GetValue()
	}
	if reqMember.Clazz != nil {
		member.Clazz = reqMember.Clazz.GetValue()
	}
	if reqMember.CodeforcesId != nil {
		member.CodeforcesId = reqMember.CodeforcesId.GetValue()
	}
	if reqMember.AtcoderId != nil {
		member.AtcoderId = reqMember.AtcoderId.GetValue()
	}
	if reqMember.CodechefId != nil {
		member.CodechefId = reqMember.CodechefId.GetValue()
	}
	if reqMember.NowcoderId != nil {
		member.NowcoderId = reqMember.NowcoderId.GetValue()
	}
	if reqMember.VjudgeId != nil {
		member.VjudgeId = reqMember.VjudgeId.GetValue()
	}
	if reqMember.LeetcodeId != nil {
		member.LeetcodeId = reqMember.LeetcodeId.GetValue()
	}
	if reqMember.LuoguId != nil {
		member.LuoguId = reqMember.LuoguId.GetValue()
	}
	if reqMember.Email != nil {
		member.Email = reqMember.Email.GetValue()
	}
	if reqMember.SubscribeStatus.GetIsSubscribe() != nil {
		member.IsSubscribe = reqMember.SubscribeStatus.IsSubscribe.GetValue()
	}
	if reqMember.SubscribeStatus.GetEmailBit() != nil {
		member.EmailBit = reqMember.SubscribeStatus.GetEmailBit().GetValue()
	}
}
