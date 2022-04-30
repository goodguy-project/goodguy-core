package model

import (
	"github.com/goodguy-project/goodguy-core/core/oj"
	"github.com/goodguy-project/goodguy-core/idl"
)

type Contest struct {
	ContestMessage *idl.RecentContest_ContestMessage
	OnlineJudge    *oj.OnlineJudge
}
