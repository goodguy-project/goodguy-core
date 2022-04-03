package util

import "encoding/json"

func JsonToString(x interface{}) string {
	s, _ := json.Marshal(x)
	return string(s)
}
