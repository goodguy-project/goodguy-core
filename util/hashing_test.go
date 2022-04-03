package util

import (
	"fmt"
	"testing"
)

func Test_PwdHash(t *testing.T) {
	t.Run("run", func(t *testing.T) {
		fmt.Println(Hashing("admin"))
	})
}
