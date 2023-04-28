package oauthutil

import (
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/authhandler"

	"go.indent.com/indent-go/pkg/rand"
)

const (
	// pkceVerifierLen is the length of the OAuth2 PKCE code verifier.
	pkceVerifierLen = 32

	// pkceChallengeMethod is the method used to generate the PKCE challenge.
	pkceChallengeMethod = "S256"

	// PKCE auth URL params
	authURLParamPKCEChallenge       = "code_challenge"
	authURLParamPKCEChallengeMethod = "code_challenge_method"
	authURLParamPKCEVerifier        = "code_verifier"
)

// PKCEVerifier is a OAuth2 PKCE code verifier.
type PKCEVerifier authhandler.PKCEParams

// NewPKCEVerifier returns a new code verifier.
func NewPKCEVerifier() *PKCEVerifier {
	verifier := rand.Hex(pkceVerifierLen)
	challenge := sha256.Sum256([]byte(verifier))
	return &PKCEVerifier{
		Challenge:       base64.RawURLEncoding.EncodeToString(challenge[:]),
		ChallengeMethod: pkceChallengeMethod,
		Verifier:        verifier,
	}
}

// AuthOpts returns the options to be passed to the OAuth2 authorization code flow from params.
func (v *PKCEVerifier) AuthOpts() []oauth2.AuthCodeOption {
	if v == nil {
		return nil
	}
	return []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam(authURLParamPKCEChallenge, v.Challenge),
		oauth2.SetAuthURLParam(authURLParamPKCEChallengeMethod, v.ChallengeMethod),
	}
}

// TokenOpt returns the options to be passed in the OAuth2 token exchange.
func (v *PKCEVerifier) TokenOpt() oauth2.AuthCodeOption {
	if v == nil {
		return nil
	}
	return oauth2.SetAuthURLParam(authURLParamPKCEVerifier, v.Verifier)
}
