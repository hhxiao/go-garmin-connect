//go:build integration
// +build integration

// Package integration holds the live-credential test suite for go-garmin-connect.
//
// All tests in this package are gated behind the `integration` build tag so
// `go test ./...` stays fast and offline. To run them:
//
//	GARMIN_LOGIN=you@example.com GARMIN_PASSWORD=… \
//	  go test -tags=integration ./integration/...
//
// Tokens are cached in `$HOME/.garmin_token.json` to keep the SSO flow off
// the hot path between runs. A few tests are skipped by default because they
// mutate device state — set GARMIN_RUN_DESTRUCTIVE=1 to opt in.
package integration

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	garmin "github.com/sealbro/go-garmin-connect"
	"github.com/sealbro/go-garmin-connect/auth"
)

var (
	clientOnce sync.Once
	cachedC    *garmin.Client
	clientErr  error
)

// lazyClient lazily builds a singleton Garmin client backed by env-var
// credentials and a file token cache. Failure to construct fails the test
// once and short-circuits subsequent tests with the same error.
func lazyClient(t *testing.T) *garmin.Client {
	t.Helper()
	clientOnce.Do(func() {
		login := os.Getenv("GARMIN_LOGIN")
		password := os.Getenv("GARMIN_PASSWORD")
		if login == "" || password == "" {
			clientErr = errMissingCreds
			return
		}
		params, err := auth.NewBasicAuth(login, password)
		if err != nil {
			clientErr = err
			return
		}
		home, _ := os.UserHomeDir()
		cache := auth.NewFileTokenCache(filepath.Join(home, ".garmin_token.json"))
		ctx := garmin.NewContext(nil, params, garmin.WithTokenCache(cache))
		cachedC = garmin.NewClient(ctx)
	})
	if clientErr != nil {
		t.Skipf("integration client unavailable: %v", clientErr)
	}
	return cachedC
}

// runDestructive returns true when the caller has explicitly opted in to
// tests that mutate state (sleep window, send-to-device, file upload).
func runDestructive() bool { return os.Getenv("GARMIN_RUN_DESTRUCTIVE") == "1" }

type stringErr string

func (s stringErr) Error() string { return string(s) }

const errMissingCreds = stringErr("set GARMIN_LOGIN and GARMIN_PASSWORD to run integration tests")
