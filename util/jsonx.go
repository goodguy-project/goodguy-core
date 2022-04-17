package util

import (
	"encoding/json"
	"log"
)

func Json(x interface{}) string {
	s, err := json.Marshal(x)
	if err != nil {
		log.Printf("json marshal error, err: %v\n", err)
		return ""
	}
	return string(s)
}
