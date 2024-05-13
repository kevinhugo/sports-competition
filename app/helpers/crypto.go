package helpers

import (
	"crypto/sha512"
	"fmt"
)

func SHA512(plainText string) string {
	newSHAHasher := sha512.New()
	newSHAHasher.Write([]byte(plainText))

	return fmt.Sprintf("%x", newSHAHasher.Sum(nil))
}
