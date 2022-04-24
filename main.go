package main

import (
	"github.com/goodguy-project/goodguy-core/core/timing/statistics"
	"github.com/goodguy-project/goodguy-core/core/timing/subscribe"
	"github.com/goodguy-project/goodguy-core/core/web"
	"github.com/goodguy-project/goodguy-core/initialize"
)

func main() {
	initialize.MustInit()
	go statistics.Serve()
	go subscribe.Serve()
	go web.Serve()
	select {} // wait for shutdown
}
