package model

import (
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
}

func (m *Member) ToProtoMember() *idl.Member {
	r := &idl.Member{
		Id:           int64(m.ID),
		Sid:          m.Sid,
		Name:         m.Name,
		School:       m.School,
		Grade:        m.Grade,
		Clazz:        m.Clazz,
		IsOfficial:   m.IsOfficial,
		CodeforcesId: m.CodeforcesId,
		AtcoderId:    m.AtcoderId,
		CodechefId:   m.CodechefId,
		NowcoderId:   m.NowcoderId,
		VjudgeId:     m.VjudgeId,
		LeetcodeId:   m.LeetcodeId,
		LuoguId:      m.LuoguId,
		Email:        m.Email,
	}
	if m.IsSubscribe {
		r.IsSubscribe = idl.Bool_Bool_True
	} else {
		r.IsSubscribe = idl.Bool_Bool_False
	}
	return r
}
