package auth

import (
	"crypto/sha1"
	"fmt"
)

func GenerateHasPassword(password string, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
