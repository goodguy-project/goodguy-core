package model

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/gorm"

	"github.com/goodguy-project/goodguy-core/idl"
)

type Member struct {
	gorm.Model
	Sid          string `gorm:"uniqueIndex,size:128"`
	Name         string `gorm:"index:member_name_index"`
	School       string
	Grade        int32
	Clazz        string
	IsOfficial   bool
	CodeforcesId string `gorm:"index:member_codeforces_id_index"`
	AtcoderId    string `gorm:"index:member_atcoder_id_index"`
	CodechefId   string `gorm:"index:member_codechef_id_index"`
	NowcoderId   string `gorm:"index:member_nowcoder_id_index"`
	VjudgeId     string `gorm:"index:member_vjudge_id_index"`
	LeetcodeId   string `gorm:"index:member_leetcode_id_index"`
	LuoguId      string `gorm:"index:member_luogu_id_index"`
	Email        string `gorm:"index:member_email_index"`
	IsSubscribe  bool   `gorm:"index:member_is_subscribe_index"`
	IsAdmin      bool
	Pwd          string
	SubscribeBit int64 // TODO
}

func (m *Member) ToProtoMember() *idl.Member {
	r := &idl.Member{
		Id:           int64(m.ID),
		Sid:          wrapperspb.String(m.Sid),
		Name:         wrapperspb.String(m.Name),
		School:       wrapperspb.String(m.School),
		Grade:        wrapperspb.Int32(m.Grade),
		Clazz:        wrapperspb.String(m.Clazz),
		IsOfficial:   wrapperspb.Bool(m.IsOfficial),
		CodeforcesId: wrapperspb.String(m.CodeforcesId),
		AtcoderId:    wrapperspb.String(m.AtcoderId),
		CodechefId:   wrapperspb.String(m.CodechefId),
		NowcoderId:   wrapperspb.String(m.NowcoderId),
		VjudgeId:     wrapperspb.String(m.VjudgeId),
		LeetcodeId:   wrapperspb.String(m.LeetcodeId),
		LuoguId:      wrapperspb.String(m.LuoguId),
		Email:        wrapperspb.String(m.Email),
		IsSubscribe:  wrapperspb.Bool(m.IsSubscribe),
	}
	return r
}
