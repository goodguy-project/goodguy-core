package model

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/gorm"

	"github.com/goodguy-project/goodguy-core/idl"
)

type SubscribeStatus struct {
	IsSubscribe bool `gorm:"index"`
	EmailBit    uint64
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
	IsOfficial   bool
	CodeforcesId string `gorm:"index"`
	AtcoderId    string `gorm:"index"`
	CodechefId   string `gorm:"index"`
	NowcoderId   string `gorm:"index"`
	VjudgeId     string `gorm:"index"`
	LeetcodeId   string `gorm:"index"`
	LuoguId      string `gorm:"index"`
	Email        string `gorm:"index"`
	IsAdmin      bool
	Pwd          string
	SelfingMode  bool
	SMTP
	SubscribeStatus
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
		SelfingMode:  wrapperspb.Bool(m.SelfingMode),
		Smtp: &idl.SMTP{
			Host: m.Host,
			Port: int32(m.Port),
			Pwd:  m.Pwd,
		},
		SubscribeStatus: &idl.SubscribeStatus{
			IsSubscribe: wrapperspb.Bool(m.IsSubscribe),
			EmailBit:    wrapperspb.UInt64(m.EmailBit),
		},
	}
	return r
}
