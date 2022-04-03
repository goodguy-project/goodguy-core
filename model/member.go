package model

import (
	"gorm.io/gorm"

	"github.com/goodguy-project/goodguy-core/idl"
)

type Member struct {
	gorm.Model
	Sid          string `gorm:"uniqueIndex,size:128"`
	Name         string
	School       string
	Grade        int32
	Clazz        string
	IsOfficial   bool
	CodeforcesId string
	AtcoderId    string
	CodechefId   string
	NowcoderId   string
	VjudgeId     string
	LeetcodeId   string
	LuoguId      string
	Email        string
	IsSubscribe  bool
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
