package util

import (
	"log"

	"github.com/bytedance/sonic"
)

func init() {
	Json = func(x interface{}) string {
		s, err := sonic.Marshal(x)
		if err != nil {
			log.Printf("json marshal error, err: %v\n", err)
			return ""
		}
		return string(s)
	}
}
