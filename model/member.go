package model

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/gorm"

	"github.com/goodguy-project/goodguy-core/idl"
)

type SubscribeStatus struct {
	Email uint64
}

type SMTP struct {
	Host string
	Port int
	Pwd  string
}

type Member struct {
	gorm.Model
	Sid          string `gorm:"uniqueIndex,size:128"`
	Name         string `gorm:"index"`
	School       string
	Grade        int32
	Clazz        string
	IsOfficial   bool `gorm:"index"`
	CodeforcesId string
	AtcoderId    string
	CodechefId   string
	NowcoderId   string
	VjudgeId     string
	LeetcodeId   string
	LuoguId      string
	Email        string
	IsAdmin      bool
	Pwd          string
	SelfingMode  bool
	TeamName     string `gorm:"index"`
	SMTP
	SubscribeStatus
}

// ToProtoMember 只能传可以公开的信息
func (m *Member) ToProtoMember() *idl.Member {
	r := &idl.Member{
		Id:           int64(m.ID),
		Sid:          wrapperspb.String(m.Sid),
		Name:         wrapperspb.String(m.Name),
		School:       wrapperspb.String(m.School),
		Grade:        wrapperspb.Int32(m.Grade),
		Clazz:        wrapperspb.String(m.Clazz),
		IsOfficial:   wrapperspb.Bool(m.IsOfficial),
		TeamName:     wrapperspb.String(m.TeamName),
		CodeforcesId: wrapperspb.String(m.CodeforcesId),
		AtcoderId:    wrapperspb.String(m.AtcoderId),
		CodechefId:   wrapperspb.String(m.CodechefId),
		NowcoderId:   wrapperspb.String(m.NowcoderId),
		VjudgeId:     wrapperspb.String(m.VjudgeId),
		LeetcodeId:   wrapperspb.String(m.LeetcodeId),
		LuoguId:      wrapperspb.String(m.LuoguId),
	}
	return r
}
