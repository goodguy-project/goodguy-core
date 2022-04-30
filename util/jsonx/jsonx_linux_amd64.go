package jsonx

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
	Marshal = func(x interface{}) (string, error) {
		s, err := sonic.Marshal(x)
		return string(s), err
	}
	Unmarshal = func(b []byte, d interface{}) error {
		return sonic.Unmarshal(b, d)
	}
}
