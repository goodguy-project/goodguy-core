package util

import (
	"crypto/sha512"
	"fmt"
)

func Hashing(pwd string) string {
	b := []byte(fmt.Sprintf("%s>_<%s>_<%s", pwd, pwd, pwd))
	return fmt.Sprintf("%x", sha512.Sum512(b))
}
