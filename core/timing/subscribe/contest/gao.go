package contest

import (
	"context"
	"log"

	"github.com/goodguy-project/goodguy-core/client/crawl"
	"github.com/goodguy-project/goodguy-core/core/oj"
	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/util/conf"
)

func getSubscribeContest(contests map[*oj.OnlineJudge][]*idl.RecentContest_ContestMessage, bit uint64) []*model.Contest {
	var r []*model.Contest
	for onlineJudge, contest := range contests {
		if onlineJudge.Contain(bit) {
			for _, c := range contest {
				r = append(r, &model.Contest{
					ContestMessage: c,
					OnlineJudge:    onlineJudge,
				})
			}
		}
	}
	return r
}

func doCrawl(ctx context.Context) map[*oj.OnlineJudge][]*idl.RecentContest_ContestMessage {
	resp, err := crawl.Client.MGetRecentContest(ctx, &idl.MGetRecentContestRequest{})
	r := make(map[*oj.OnlineJudge][]*idl.RecentContest_ContestMessage)
	if err != nil {
		log.Printf("MGetRecentContest error, err: %v\n", err)
		return r
	}
	for _, rc := range resp.GetRecentContest() {
		onlineJudge, ok := oj.OJMap[rc.Platform]
		if !ok {
			log.Printf("not support online judge %s", rc.Platform)
		}
		r[onlineJudge] = rc.RecentContest
	}
	return r
}

func gao() {
	var err error
	db := model.GetDB()
	count := int64(0)
	err = db.Model(&model.Member{}).Count(&count).Error
	if err != nil {
		log.Printf("database error, err: %v\n", err)
		return
	}
	buffer := conf.Viper().GetInt64("subscribe.buffer")
	if buffer <= 0 {
		buffer = 100
	}
	times := (count + buffer - 1) / buffer
	for no := int64(0); no < times; no += 1 {
		var members []*model.Member
		err = db.Model(&model.Member{}).Offset(int(buffer * no)).Limit(int(buffer)).
			Find(&members).Error
		if err != nil {
			log.Printf("database error, err: %v\n", err)
		}
		contests := doCrawl(context.Background())
		doEmail(members, contests)
	}
}
