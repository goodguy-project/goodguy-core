package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
)

func GetMember(ctx context.Context, req *idl.GetMemberRequest) (*idl.GetMemberResponse, error) {
	response := &idl.GetMemberResponse{}
	if req == nil {
		return response, status.Error(codes.DataLoss, "login failed")
	}
	pageNo := req.PageNo
	pageSize := req.PageSize
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 || pageSize > 500 {
		pageSize = 10
	}
	db := model.GetDB().Model(&model.Member{})
	if len(req.Id) > 0 {
		db = db.Where("id IN ?", req.Id)
	}
	if len(req.Sid) > 0 {
		db = db.Where("sid IN ?", req.Sid)
	}
	if len(req.IsOfficial) > 0 {
		db = db.Where("is_official IN ?", req.IsOfficial)
	}
	if len(req.Name) > 0 {
		db = db.Where("name IN ?", req.Name)
	}
	if len(req.School) > 0 {
		db = db.Where("school IN ?", req.School)
	}
	if len(req.Grade) > 0 {
		db = db.Where("grade IN ?", req.Grade)
	}
	if len(req.Clazz) > 0 {
		db = db.Where("clazz IN ?", req.Clazz)
	}
	if len(req.TeamName) > 0 {
		db = db.Where("team_name IN ?", req.TeamName)
	}
	if len(req.CodeforcesId) > 0 {
		db = db.Where("codeforces_id IN ?", req.CodeforcesId)
	}
	if len(req.AtcoderId) > 0 {
		db = db.Where("atcoder_id IN ?", req.AtcoderId)
	}
	if len(req.CodechefId) > 0 {
		db = db.Where("codechef_id IN ?", req.CodechefId)
	}
	if len(req.NowcoderId) > 0 {
		db = db.Where("nowcoder_id IN ?", req.NowcoderId)
	}
	if len(req.VjudgeId) > 0 {
		db = db.Where("vjudge_id IN ?", req.VjudgeId)
	}
	if len(req.LeetcodeId) > 0 {
		db = db.Where("leetcode_id IN ?", req.LeetcodeId)
	}
	if len(req.LuoguId) > 0 {
		db = db.Where("luogu_id IN ?", req.LuoguId)
	}
	var members []*model.Member
	db.Order("id desc").Offset(int(pageSize * (pageNo - 1))).Limit(int(pageSize)).Find(&members)
	for _, member := range members {
		response.Member = append(response.Member, member.ToProtoMember())
	}
	count := int64(0)
	model.GetDB().Model(&model.Member{}).Count(&count)
	response.Size = int32(count)
	return response, nil
}
