package oj

import "github.com/goodguy-project/goodguy-core/idl"

type OnlineJudge struct {
	platform string
	enum     int
}

func (oj *OnlineJudge) Name() string {
	return oj.platform
}

func (oj *OnlineJudge) Enum() int {
	return oj.enum
}

func (oj *OnlineJudge) Bit() uint64 {
	return uint64(1) << oj.enum
}

func (oj *OnlineJudge) Contain(b uint64) bool {
	return (b & (uint64(1) << oj.enum)) > 0
}

var (
	Codeforces = &OnlineJudge{platform: "codeforces", enum: int(idl.OnlineJudge_OnlineJudge_Codeforces)}
	Atcoder    = &OnlineJudge{platform: "atcoder", enum: int(idl.OnlineJudge_OnlineJudge_Atcoder)}
	Codechef   = &OnlineJudge{platform: "codechef", enum: int(idl.OnlineJudge_OnlineJudge_Codechef)}
	Nowcoder   = &OnlineJudge{platform: "nowcoder", enum: int(idl.OnlineJudge_OnlineJudge_Nowcoder)}
	Leetcode   = &OnlineJudge{platform: "leetcode", enum: int(idl.OnlineJudge_OnlineJudge_Leetcode)}
	Luogu      = &OnlineJudge{platform: "luogu", enum: int(idl.OnlineJudge_OnlineJudge_Luogu)}
	Vjudge     = &OnlineJudge{platform: "vjudge", enum: int(idl.OnlineJudge_OnlineJudge_Vjudge)}
	OJMap      = map[string]*OnlineJudge{
		"codeforces": Codeforces,
		"atcoder":    Atcoder,
		"codechef":   Codechef,
		"nowcoder":   Nowcoder,
		"leetcode":   Leetcode,
		"luogu":      Luogu,
		"vjudge":     Vjudge,
	}
)
