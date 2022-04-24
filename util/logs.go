package util

import (
	"log"

	"github.com/goodguy-project/goodguy-core/util/conf"
)

func Debug(format string, args ...interface{}) {
	if conf.Viper().GetBool("debug") {
		log.Printf(format, args...)
	}
}
