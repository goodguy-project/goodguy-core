package util

import (
	"fmt"
	"testing"
)

func TestPwdHash(t *testing.T) {
	t.Run("run", func(t *testing.T) {
		fmt.Println(Hashing("admin"))
	})
}
