package oauthutil

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"

	"go.indent.com/indent-go/pkg/oidcclaims"
)

const (
	testStoreName = "testStore"
)

func TestStoreNew(t *testing.T) {
	store := newTestStore(t, testStoreName)
	defer cleanupStore(t, store)

	// check store setup correctly
	assert.Equal(t, testStoreName, store.Name())

	// remainder should not work without login
	claims, err := store.Claims()
	assert.Nil(t, claims)
	assert.Error(t, err)
	err = store.UpdateUserInfo(context.Background())
	assert.Error(t, err)
}

func TestStoreLoginFail(t *testing.T) {
	store := newTestStore(t, t.Name())
	defer cleanupStore(t, store)

	loginOpts := NewLoginOptions()
	assert.Error(t, store.Login(context.Background(), loginOpts))
}

func TestStoreBadConfig(t *testing.T) {
	store := newTestStore(t, t.Name())
	defer cleanupStore(t, store)

	// write invalid token file
	assert.NoError(t, store.writeToken(new(tokenFile)))

	// remainder should not work
	claims, err := store.Claims()
	assert.Nil(t, claims)
	assert.NoError(t, err)
	err = store.UpdateUserInfo(context.Background())
	assert.Error(t, err)
}

func TestStoreEndpoints(t *testing.T) {
	store := newTestStore(t, t.Name())
	defer cleanupStore(t, store)

	expected := &oidcclaims.Standard{
		Sub:               "248289761001",
		Name:              "Jane Doe",
		GivenName:         "Jane",
		FamilyName:        "Doe",
		MiddleName:        "He",
		Nickname:          "Janie",
		PreferredUsername: "jane.doe",
		Profile:           "https://example.com/jane.doe",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/"+pathUserInfo, req.URL.Path)

		data, err := json.Marshal(expected)
		assert.NoError(t, err)
		_, err = w.Write(data)
		assert.NoError(t, err)
	}))
	defer server.Close()

	// write token file with testing endpoint set
	assert.NoError(t, store.writeToken(&tokenFile{
		OAuth: &oauth2.Config{
			Endpoint: oauth2.Endpoint{
				TokenURL: server.URL + "/a/b",
			},
		},
		Token: &oauth2.Token{
			AccessToken: "someToken",
		},
	}))

	// update userinfo
	assert.NoError(t, store.UpdateUserInfo(context.Background()))

	// check claims
	claims, err := store.Claims()
	assert.NoError(t, err)
	assert.Equal(t, expected, claims.Standard)
}

func newTestStore(t *testing.T, name string) *FileStore {
	t.Helper()
	tmpDir, err := os.MkdirTemp("", testStoreName+name)
	assert.NoError(t, err)
	store := NewStore(tmpDir, name)
	store.Directory = tmpDir
	return store
}

func cleanupStore(t *testing.T, s *FileStore) {
	t.Helper()
	assert.NoError(t, os.RemoveAll(s.Directory))
}
