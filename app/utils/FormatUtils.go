package utils

import (
	"crypto/sha1"
	"fmt"
)

func Encrypt(plaintext string) (crypto string) {
	crypto = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
