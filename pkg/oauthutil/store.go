package oauthutil

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"

	"go.indent.com/indent-go/pkg/oidcclaims"
)

const (
	fileModeDir = os.FileMode(0o700)
)

var (
	// ErrMissingOAuthToken is returned when OAuth token has not been provided.
	ErrMissingOAuthToken = errors.New("missing OAuth token")
)

// Should act as TokenSource
var _ Store = new(FileStore)

// Store allows for login and accessing Tokens.
type Store interface {
	oauth2.TokenSource
	Name() string
	Login(loginOpts *LoginOptions) error
	Claims() (*oidcclaims.Claims, error)
	UpdateUserInfo() error
}

// NewStore returns a FileStore with defaults set.
func NewStore(ctx context.Context, headless bool, credentialDir, envName string) *FileStore {
	return &FileStore{
		ctx:       ctx,
		headless:  headless,
		Directory: credentialDir,
		Filename:  envName,
	}
}

// FileStore manages credentials locally.
type FileStore struct {
	// ctx used to create token source.
	ctx context.Context

	// headless indicates that the login process should not open a browser.
	headless bool

	// Directory credentials are stored.
	Directory string

	// Filename of credential being used.
	Filename string
}

type tokenFile struct {
	OAuth  *oauth2.Config
	Token  *oauth2.Token
	Claims *oidcclaims.Claims
}

// Token reads a token from the FileStore.
func (s *FileStore) Token() (*oauth2.Token, error) {
	t, err := s.readToken()
	if err != nil {
		return nil, err
	} else if !t.Token.Valid() {
		if s.headless {
			return nil, fmt.Errorf("token is invalid, expiry %s", t.Token.Expiry)
		}

		opts := NewLoginOptions()
		opts.OAuth = t.OAuth

		if err = s.Login(opts); err != nil {
			return nil, fmt.Errorf("failed to get new refresh token: %w", err)
		} else if t, err = s.readToken(); err != nil {
			return nil, fmt.Errorf("failed to read token after successful login: %w", err)
		}
	}
	return t.Token, nil
}

// Name of the FileStore.
func (s *FileStore) Name() string {
	return s.Filename
}

// Login authorizes the user and writes their token to the FileStore.
func (s *FileStore) Login(loginOpts *LoginOptions) error {
	if code, verifier, err := Login(loginOpts); err != nil {
		return fmt.Errorf("failed to get new token: %w", err)
	} else if token, err := loginOpts.OAuth.Exchange(s.ctx, code, verifier.TokenOpt()); err != nil {
		return fmt.Errorf("failed to exchange code for access token: %w", err)
	} else if err = s.writeToken(&tokenFile{Token: token, OAuth: loginOpts.OAuth}); err != nil {
		return fmt.Errorf("failed to write token: %w", err)
	}
	return nil
}

// Claims related to current token.
func (s *FileStore) Claims() (*oidcclaims.Claims, error) {
	t, err := s.readToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	return t.Claims, nil
}

// UpdateUserInfo sets data returned by userinfo as claims.
func (s *FileStore) UpdateUserInfo() error {
	if t, err := s.readToken(); err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	} else if t.Claims, err = UserInfo(s.ctx, t.OAuth, t.Token); err != nil {
		return fmt.Errorf("failed retrieve UserInfo: %w", err)
	} else if err = s.writeToken(t); err != nil {
		return fmt.Errorf("failed to write token: %w", err)
	}
	return nil
}

func (s *FileStore) readToken() (*tokenFile, error) {
	credFile := s.credFile()
	f, err := os.Open(credFile) // #nosec G304 // Should only be used by clients
	if err != nil {
		return nil, fmt.Errorf("failed to read token file '%s': %w", credFile, err)
	}
	defer closeFile(f)

	t := new(tokenFile)
	if err = gob.NewDecoder(f).Decode(t); err != nil {
		return nil, fmt.Errorf("failed to decode token: %w", err)
	}
	return t, nil
}

func (s *FileStore) writeToken(t *tokenFile) error {
	if _, err := os.Stat(s.Directory); os.IsNotExist(err) {
		if err = os.MkdirAll(s.Directory, fileModeDir); err != nil {
			return fmt.Errorf("failed to create directory '%s': %w", s.Directory, err)
		}
	}

	credFile := s.credFile()
	f, err := os.Create(credFile) // #nosec G304 // only used by clients and tests
	if err != nil {
		return fmt.Errorf("failed to write token file '%s': %w", credFile, err)
	}
	defer closeFile(f)

	if err = gob.NewEncoder(f).Encode(t); err != nil {
		return fmt.Errorf("failed to encode token: %w", err)
	}
	return nil
}

func (s *FileStore) credFile() string {
	return filepath.Join(s.Directory, s.Filename)
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Printf("Failed to close file '%s': %v", file.Name(), err)
	}
}
