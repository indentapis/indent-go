package event

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go.indent.com/indent-go/pkg/common"
)

func TestVerify(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	secret, timestamp, msgData, sig := testSign(r)
	err := Verify(secret, timestamp, msgData, sig)
	r.NoError(err)
}

func TestVerifyFail(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	secret, timestamp, msgData, sig := testSign(r)
	var extraByte byte
	msgData = append(msgData, extraByte)
	err := Verify(secret, timestamp, msgData, sig)
	r.Error(err)
}

func TestVerifyHeaders(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	secret, timestamp, msgData, sig := testSign(r)
	headers := map[string]string{
		common.HeaderSignature: sig,
		common.HeaderTimestamp: timestamp,
	}
	err := VerifyHeaders(secret, headers, msgData)
	r.NoError(err)
}

func TestVerifyHeadersEmpty(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	secret, _, msgData, _ := testSign(r)
	err := VerifyHeaders(secret, nil, msgData)
	r.Error(err)
}
