package event

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"go.indent.com/indent-go/pkg/common"
)

const (
	expectedSig       = "c27f187960625ebccc33336030b071dac1d54b1e98d9e27fe5452ee00c935fb0"
	expectedTimestamp = "2020-05-01T07:00:00Z"
)

func TestSignatureFromHeaders(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	sig, timestamp := SignatureFromHeaders(map[string]string{
		common.HeaderSignature: expectedSig,
		common.HeaderTimestamp: expectedTimestamp,
		"otherData":            "otherValue",
	})
	r.Equal(expectedSig, sig)
	r.Equal(expectedTimestamp, timestamp)
}

func TestSignatureFromHeadersNoSignature(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	sig, timestamp := SignatureFromHeaders(map[string]string{
		"X-Indent-TimeStamp": expectedTimestamp,
	})
	r.Empty(sig)
	r.Equal(expectedTimestamp, timestamp)
}

func TestSignatureFromHeadersNoTimestamp(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	sig, timestamp := SignatureFromHeaders(map[string]string{
		strings.ToUpper(common.HeaderSignature): expectedSig,
	})
	r.Equal(expectedSig, sig)
	r.Empty(timestamp)
}

func TestSignatureFromHeadersEmpty(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	sig, timestamp := SignatureFromHeaders(nil)
	r.Empty(sig)
	r.Empty(timestamp)
}
