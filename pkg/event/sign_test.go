package event

import (
	"errors"
	"hash"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.indent.com/indent-go/pkg/rand"
)

const (
	testSecretLen = 200
)

func TestSign(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	testSign(r)
}

func TestSignFail(t *testing.T) {
	r := require.New(t)

	testHashFunc = func() hash.Hash {
		return &errHash{}
	}
	defer func() { testHashFunc = nil }()

	_, err := Sign(nil, "", nil)
	r.Error(err)
}

func testSign(r *require.Assertions) (secret []byte, timestamp string, msgData []byte, sig string) {
	secret = rand.Bytes(testSecretLen)
	timestamp = time.Now().UTC().Format(time.RFC3339)
	msgData = []byte("Hello, world!")

	var err error
	sig, err = Sign(secret, timestamp, msgData)
	r.NotEmpty(sig)
	r.NoError(err)
	return
}

// errHash is a hash.Hash that always errors on Write.
type errHash struct {
	hash.Hash
}

func (errHash) Write([]byte) (int, error) {
	return -1, errors.New("always errors")
}
func (errHash) BlockSize() int {
	return 0
}
