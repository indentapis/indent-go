package event

import (
	"fmt"
)

// Verify uses a secret and timestamp to verify data.
func Verify(secret []byte, timestamp string, data []byte, sig string) error {
	actual, err := Sign(secret, timestamp, data)
	if err == nil && actual != sig {
		return fmt.Errorf("signature mismatch: %s != %s", actual, sig)
	}
	return err
}

// VerifyHeaders uses a secret and headers to verify data.
func VerifyHeaders(secret []byte, headers map[string]string, data []byte) error {
	sig, timestamp := SignatureFromHeaders(headers)
	return Verify(secret, timestamp, data, sig)
}
