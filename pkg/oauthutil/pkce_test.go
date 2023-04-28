package oauthutil

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPKCEVerifier(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	verifier := NewPKCEVerifier()
	r.Len(verifier.Verifier, hex.EncodedLen(pkceVerifierLen))
	r.Len(verifier.Challenge, base64.RawURLEncoding.EncodedLen(sha256.Size))
	r.Equal("S256", verifier.ChallengeMethod)

	authOpts := verifier.AuthOpts()
	r.Len(authOpts, 2)
	r.NotNil(authOpts[0])
	r.NotNil(authOpts[1])

	tokenOpt := verifier.TokenOpt()
	r.NotNil(tokenOpt)
}
