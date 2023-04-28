package rand

import (
	"crypto/rand"
	"encoding/hex"
)

// Bytes returns a random byte slice of the specified length. It panics if it fails to read.
func Bytes(length int) []byte {
	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	return buf
}

// Hex returns random bytes of the specified length represented as hex. It panics if it fails to read.
func Hex(length int) string {
	return hex.EncodeToString(Bytes(length))
}
