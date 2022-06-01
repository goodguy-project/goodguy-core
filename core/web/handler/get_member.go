package handler

import (
	"context"

	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
)

func GetMember(ctx context.Context, req *idl.GetMemberRequest) (*idl.GetMemberResponse, error) {
	response := &idl.GetMemberResponse{}
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
