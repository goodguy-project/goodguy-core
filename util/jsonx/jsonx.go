package jsonx

import (
	"encoding/json"
	"log"
)

var Json func(x interface{}) string = nil
var Marshal func(x interface{}) (string, error) = nil
var Unmarshal func(b []byte, d interface{}) error = nil

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
	if Marshal == nil {
		Marshal = func(x interface{}) (string, error) {
			s, err := json.Marshal(x)
			return string(s), err
		}
	}
	if Unmarshal == nil {
		Unmarshal = func(b []byte, d interface{}) error {
			return json.Unmarshal(b, d)
		}
	}
}
