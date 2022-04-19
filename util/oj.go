package util

import "github.com/goodguy-project/goodguy-core/idl"

type OnlineJudge struct {
	platform string
	bit      int64
}

func (oj *OnlineJudge) Name() string {
	return oj.platform
}

func (oj *OnlineJudge) Bit() int64 {
	return oj.bit
}

func (oj *OnlineJudge) Contain(b int64) bool {
	return (b & oj.bit) > 0
}

var (
	Codeforces = &OnlineJudge{platform: "codeforces", bit: 1 << idl.OnlineJudge_OnlineJudge_Codeforces}
	Atcoder    = &OnlineJudge{platform: "atcoder", bit: 1 << idl.OnlineJudge_OnlineJudge_Atcoder}
	Codechef   = &OnlineJudge{platform: "codechef", bit: 1 << idl.OnlineJudge_OnlineJudge_Codechef}
	Nowcoder   = &OnlineJudge{platform: "nowcoder", bit: 1 << idl.OnlineJudge_OnlineJudge_Nowcoder}
	Leetcode   = &OnlineJudge{platform: "leetcode", bit: 1 << idl.OnlineJudge_OnlineJudge_Leetcode}
	Luogu      = &OnlineJudge{platform: "luogu", bit: 1 << idl.OnlineJudge_OnlineJudge_Luogu}
	Vjudge     = &OnlineJudge{platform: "vjudge", bit: 1 << idl.OnlineJudge_OnlineJudge_Vjudge}
)
