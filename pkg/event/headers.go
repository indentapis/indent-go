package event

import (
	"strings"

	"go.indent.com/indent-go/pkg/common"
)

// SignatureFromHeaders returns the signature and timestamp from headers.
func SignatureFromHeaders(headers map[string]string) (sig, timestamp string) {
	for k, v := range headers {
		header := strings.ToLower(k)
		switch header {
		case strings.ToLower(common.HeaderSignature):
			sig = v
		case strings.ToLower(common.HeaderTimestamp):
			timestamp = v
		}
	}
	return sig, timestamp
}
