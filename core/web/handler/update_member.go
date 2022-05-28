package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/goodguy-project/goodguy-core/core/web/token"
	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/util"
)

func UpdateMember(ctx context.Context, req *idl.UpdateMemberRequest) (*idl.UpdateMemberResponse, error) {
	if req.Member == nil || req.GetMember().GetSid() == nil {
		return new(idl.UpdateMemberResponse), status.Error(codes.Unauthenticated, "auth failed")
	}
	sid, ok := token.Auth(ctx)
	if !ok {
		return new(idl.UpdateMemberResponse), status.Error(codes.Unauthenticated, "auth failed")
	}
	if sid != "admin" && sid != req.GetMember().GetSid().GetValue() {
		return new(idl.UpdateMemberResponse), status.Error(codes.Unauthenticated, "auth failed")
	}
	isAdmin := sid == "admin"
	db := model.GetDB()
	member := &model.Member{}
	err := db.Model(&model.Member{}).Where("sid = ?", req.GetMember().GetSid().GetValue()).First(member).Error
	if err != nil {
		return new(idl.UpdateMemberResponse), status.Error(codes.Internal, "database error")
	}
	if sid == member.Sid && util.Hashing(req.Pwd) != member.Pwd {
		return new(idl.UpdateMemberResponse), status.Error(codes.Unauthenticated, "The password is incorrect")
	}
	doUpdateMember(isAdmin, req.Member, member)
	if req.NewPwd != nil {
		member.Pwd = util.Hashing(req.NewPwd.GetValue())
	}
	db = model.GetDB()
	db.Save(member)
	return new(idl.UpdateMemberResponse), nil
}

func doUpdateMember(isAdmin bool, reqMember *idl.Member, member *model.Member) {
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
	if reqMember.SubscribeStatus.GetEmail() != nil {
		member.SubscribeStatus.Email = reqMember.SubscribeStatus.GetEmail().GetValue()
	}
	if reqMember.SelfingMode != nil {
		member.SelfingMode = reqMember.SelfingMode.GetValue()
	}
	if reqMember.Smtp != nil {
		member.SMTP = model.SMTP{
			Host: reqMember.Smtp.Host,
			Port: int(reqMember.Smtp.Port),
			Pwd:  reqMember.Smtp.Pwd,
		}
	}
	if isAdmin && reqMember.IsOfficial != nil {
		member.IsOfficial = reqMember.IsOfficial.GetValue()
	}
	if isAdmin && reqMember.TeamName != nil {
		member.TeamName = reqMember.TeamName.GetValue()
	}
}
