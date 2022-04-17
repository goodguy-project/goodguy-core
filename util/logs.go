package util

import (
	"log"

	"github.com/spf13/viper"
)

func Debug(format string, args ...interface{}) {
	if viper.GetBool("debug") {
		log.Printf(format, args...)
	}
}
