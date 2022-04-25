package util

import (
	"encoding/json"
	"log"
)

var Json func(x interface{}) string = nil

func init() {
	if Json == nil {
		Json = func(x interface{}) string {
			s, err := json.Marshal(x)
			if err != nil {
				log.Printf("json marshal error, err: %v\n", err)
				return ""
			}
			return string(s)
		}
	}
}
