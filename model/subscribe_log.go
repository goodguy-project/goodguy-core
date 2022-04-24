package model

import (
	"time"

	"gorm.io/gorm"
)

type SubscribeLogStatus int32

const (
	SubscribeLogStatus_OK     SubscribeLogStatus = 1
	SubscribeLogStatus_Failed SubscribeLogStatus = 2
)

type SubscribeLog struct {
	gorm.Model
	Sid           string `gorm:"index"`
	SubscribeBit  int64
	SubscribeTime time.Time `gorm:"index"`
	Status        SubscribeLogStatus
}
