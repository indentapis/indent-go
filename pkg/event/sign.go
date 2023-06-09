package event

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
)

var (
	// testHashFunc allows using a different hash function for testing.
	testHashFunc func() hash.Hash
)

// Sign an event using a secret.
func Sign(secret []byte, timestamp string, eventData []byte) (sig string, err error) {
	h := newHash(secret)
	sigTxt := []byte(fmt.Sprintf("v0:%s:%s", timestamp, eventData))
	if _, err = h.Write(sigTxt); err != nil {
		err = fmt.Errorf("failed to sign message: %w", err)
		return
	}
	sig = hex.EncodeToString(h.Sum(nil))
	return
}

// newHash instantiates a new hash function with secret.
func newHash(secret []byte) hash.Hash {
	hashFunc := sha256.New
	if testHashFunc != nil {
		hashFunc = testHashFunc
	}
	return hmac.New(hashFunc, secret)
}
