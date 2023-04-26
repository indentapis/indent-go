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
	Login(ctx context.Context, loginOpts *LoginOptions) error
	Claims() (*oidcclaims.Claims, error)
	UpdateUserInfo(ctx context.Context) error
}

// NewStore returns a FileStore with defaults set.
func NewStore(credentialDir, envName string) *FileStore {
	return &FileStore{
		Directory: credentialDir,
		Filename:  envName,
	}
}

// FileStore manages credentials locally.
type FileStore struct {
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
	}
	return t.Token, nil
}

// Name of the FileStore.
func (s *FileStore) Name() string {
	return s.Filename
}

// Login authorizes the user and writes their token to the FileStore.
func (s *FileStore) Login(ctx context.Context, loginOpts *LoginOptions) error {
	if code, err := Login(loginOpts); err != nil {
		return fmt.Errorf("failed to get new token: %w", err)
	} else if token, err := loginOpts.OAuth.Exchange(ctx, code); err != nil {
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
func (s *FileStore) UpdateUserInfo(ctx context.Context) error {
	if t, err := s.readToken(); err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	} else if t.Claims, err = UserInfo(ctx, t.OAuth, t.Token); err != nil {
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
