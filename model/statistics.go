package model

import "gorm.io/gorm"

type MemberContestRecord struct {
	gorm.Model
	Sid      string `gorm:"index:member_contest_record_key"` // 用户名
	Platform string `gorm:"index:member_contest_record_key"` // 平台
	Handle   string `gorm:"index:member_contest_record_key"` // 平台ID
	Rating   int32  // rating
	Length   int32  // 场次
}
