package statistics

import (
	"context"
	"log"

	"github.com/spf13/viper"

	"github.com/goodguy-project/goodguy-core/model"
)

func gao() {
	var err error
	db := model.GetDB()
	count := int64(0)
	err = db.Model(&model.Member{}).Count(&count).Error
	if err != nil {
		log.Printf("database error, err: %v\n", err)
		return
	}
	buffer := viper.GetInt64("statistics.buffer")
	if buffer <= 0 {
		buffer = 100
	}
	times := (count + buffer - 1) / buffer
	for no := int64(0); no < times; no += 1 {
		var members []*model.Member
		err = db.Model(&model.Member{}).Offset(int(buffer * no)).Limit(int(buffer)).Find(&members).Error
		if err != nil {
			log.Printf("database error, err: %v\n", err)
		}
		_ = doContestCrawl(context.Background(), members)
	}
}
