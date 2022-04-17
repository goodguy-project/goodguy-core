package statistics

import (
	"time"

	"github.com/robfig/cron/v3"
)

func Serve() {
	shc, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithSeconds(), cron.WithLocation(shc))
	// 北京时间凌晨4点启动
	_, err := c.AddFunc("0 0 4 * * *", gao)
	if err != nil {
		panic(err)
	}
	c.Run()
}
