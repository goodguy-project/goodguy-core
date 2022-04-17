package email

import (
	"time"

	"github.com/robfig/cron/v3"
)

func Serve() {
	shc, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithSeconds(), cron.WithLocation(shc))
	_, err := c.AddFunc("0 */45 * * * *", gao)
	if err != nil {
		panic(err)
	}
	c.Run()
}
