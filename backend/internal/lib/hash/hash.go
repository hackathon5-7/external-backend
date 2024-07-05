package lib

import (
	"crypto/md5"
	"fmt"
)

// HashString generates an MD5 hash of the given string.
//
// It takes a string as input and returns a string representing the MD5 hash.
func HashString(data string) string {
	// Compute the MD5 hash of the input string.
	h := md5.Sum([]byte(data))

	// Format the hash bytes as a hexadecimal string and return it.
	return fmt.Sprintf("%x", h)
}
