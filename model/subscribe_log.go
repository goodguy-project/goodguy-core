package model

import (
	"gorm.io/gorm"
)

type SubscribeLogStatus int32

const (
	SubscribeLogStatus_Undone SubscribeLogStatus = 1
	SubscribeLogStatus_OK     SubscribeLogStatus = 2
	SubscribeLogStatus_Failed SubscribeLogStatus = 3
)

type SingleContestMsg struct {
	Name      string
	Url       string
	Timestamp int64
	Duration  int32
}

type ContestMsg struct {
	Contest  []*SingleContestMsg `gorm:"type:text[]"`
	Member   *Member             `gorm:"foreignKey:ID;references:MemberId;"`
	MemberId uint
	Platform string
}

type SubscribeLog struct {
	gorm.Model
	Sid           string `gorm:"index"`
	Type          string
	What          string
	Bit           uint64
	SubscribeTime int64
	Status        SubscribeLogStatus
	HashTag       string      `gorm:"uniqueIndex,size:100"`
	ContestMsg    *ContestMsg `gorm:"type:text"`
}
