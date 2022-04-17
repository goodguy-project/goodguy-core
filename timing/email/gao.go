package email

import (
	"context"
	"log"

	"github.com/spf13/viper"

	"github.com/goodguy-project/goodguy-core/client/crawl"
	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/util"
)

func doEmailSubscribe(ctx context.Context, subscribe []*model.Member) error {
	return nil
}

func doCrawl(ctx context.Context) (*idl.MGetRecentContestResponse, error) {
	resp, err := crawl.Client.MGetRecentContest(ctx, &idl.MGetRecentContestRequest{})
	log.Printf("resp: %v, err: %v", util.Json(resp), err)
	return resp, err
}

func gao() {
	var err error
	db := model.GetDB()
	count := int64(0)
	err = db.Model(&model.Member{}).Where("is_subscribe = 1").Count(&count).Error
	if err != nil {
		log.Printf("database error, err: %v\n", err)
		return
	}
	buffer := viper.GetInt64("email.buffer")
	if buffer <= 0 {
		buffer = 100
	}
	times := (count + buffer - 1) / buffer
	for no := int64(0); no < times; no += 1 {
		var members []*model.Member
		err = db.Model(&model.Member{}).Where("is_subscribe = 1").Offset(int(buffer * no)).Limit(int(buffer)).
			Find(&members).Error
		if err != nil {
			log.Printf("database error, err: %v\n", err)
		}
		_ = doEmailSubscribe(context.Background(), members)
	}
}
