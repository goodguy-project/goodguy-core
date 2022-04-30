package model

import (
	"gorm.io/gorm"

	"github.com/goodguy-project/goodguy-core/idl"
)

type SubscribeLogStatus int32

const (
	SubscribeLogStatus_Undone SubscribeLogStatus = 1
	SubscribeLogStatus_OK     SubscribeLogStatus = 2
	SubscribeLogStatus_Failed SubscribeLogStatus = 3
)

type ContestMsg struct {
	Contest  []*idl.RecentContest_ContestMessage
	Member   *Member
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
	HashTag       string `gorm:"uniqueIndex"`
	ContestMsg    *ContestMsg
}
