// Package oauthutil provides helpers for working with OAuth 2.
package oauthutil

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"golang.org/x/oauth2"

	"go.indent.com/indent-go/pkg/common"
	"go.indent.com/indent-go/pkg/rand"
)

var (
	// ErrMissingOAuthConfig is returned when OAuth configuration has not been provided.
	ErrMissingOAuthConfig = errors.New("missing OAuth config")
)

// LoginOptions contain the configuration for the Authorize command.
type LoginOptions struct {
	ListenAddr     string
	OAuth          *oauth2.Config
	NoRefreshToken bool
}

// NewLoginOptions returns LoginOptions with defaults set.
func NewLoginOptions() *LoginOptions {
	listenAddr := "localhost:4105"
	return &LoginOptions{
		ListenAddr: listenAddr,
	}
}

// Login prompts users to authenticate with a browser.
func Login(opts *LoginOptions) (code string, err error) {
	if opts == nil || opts.OAuth == nil {
		return "", ErrMissingOAuthConfig
	}

	// set redirect URL
	opts.OAuth.RedirectURL = "http://" + opts.ListenAddr + "/auth"

	// setup handler for specified state
	state := rand.Str(common.StateLen)
	handler, codeChan := newLoginHandler(state)

	// start server
	server := &http.Server{Addr: opts.ListenAddr, Handler: handler, ReadHeaderTimeout: 2 * time.Second}
	go startServer(server)
	defer func() {
		if closeErr := server.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close OAuth 2 redirect server: %w", closeErr)
		}
	}()

	// default to getting a refresh token
	authCodeOpt := oauth2.AccessTypeOffline
	if opts.NoRefreshToken {
		authCodeOpt = oauth2.AccessTypeOnline
	}

	// print URL and open browser for user to authorize
	authURL := opts.OAuth.AuthCodeURL(state, authCodeOpt, oauth2.SetAuthURLParam("prompt", "consent"))
	log.Printf("Click link to authorize Indent Command Line: %s", authURL)
	go openBrowser(authURL)

	// wait to receive code
	code = <-codeChan
	return code, err
}

func newLoginHandler(state string) (_ http.Handler, _ chan string) {
	codeChan := make(chan string)
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/favicon.ico" {
			http.Error(rw, "Not Found", http.StatusNotFound)
			return
		}

		if err := req.ParseForm(); err != nil {
			log.Printf("Failed to parse request form: %v", err)
			http.Error(rw, "Failed to parse request form", http.StatusInternalServerError)
		} else if reqState := req.FormValue("state"); reqState != state {
			log.Printf("Unexpected OAuth state: wanted '%s', received '%s'", state, reqState)
			http.Error(rw, "Unexpected state", http.StatusInternalServerError)
		} else if code := req.FormValue("code"); code != "" {
			if _, err = fmt.Fprint(rw, "<h1>Authentication successful</h1>"); err != nil {
				log.Printf("Failed to write success response: %v", err)
			}
			flush(rw)
			codeChan <- code
			return
		}
		log.Println("request missing code")
		http.Error(rw, "request missing code", http.StatusInternalServerError)
	}), codeChan
}

func startServer(server *http.Server) {
	log.Println("Starting redirect server")
	if err := server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
		log.Printf("Failed to listen for redirects: %v", err)
	}
}

func openBrowser(url string) {
	execOpts := []string{"xdg-open", "open", "google-chrome"}
	for _, execOpt := range execOpts {
		// Should only be used by clients
		// #nosec G204
		err := exec.Command(execOpt, url).Run()
		if err == nil {
			return
		}
	}
	log.Println("Failed launching browser")
}

func flush(rw http.ResponseWriter) {
	flusher, ok := rw.(http.Flusher)
	if !ok {
		panic("ResponseWriter does not have Flush()")
	}
	flusher.Flush()
}
