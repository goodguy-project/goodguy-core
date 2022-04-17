package main

import (
	"github.com/goodguy-project/goodguy-core/initialize"
	"github.com/goodguy-project/goodguy-core/timing/email"
	"github.com/goodguy-project/goodguy-core/timing/statistics"
	"github.com/goodguy-project/goodguy-core/web"
)

func main() {
	initialize.MustInit()
	go statistics.Serve()
	go email.Serve()
	go web.Serve()
	select {} // wait for shutdown
}
