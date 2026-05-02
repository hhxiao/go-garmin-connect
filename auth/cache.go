package auth

import (
	"context"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// TokenCache persists OAuth2 tokens between calls. The store may be local
// memory, a file, a keychain, or a shared store — it just needs to honor the
// expiry encoded by the implementation.
type TokenCache interface {
	// Get returns nil (no token) when the cache is empty or the cached token
	// has expired. A non-nil error terminates the auth flow.
	Get(ctx context.Context) (*OAuth2Token, error)
	// Set stores token and starts a fresh expiry window.
	Set(ctx context.Context, token *OAuth2Token) error
}

// InMemoryTokenCache holds the token in process memory. Tokens are lost when
// the program exits — use FileTokenCache for long-lived clients.
type InMemoryTokenCache struct {
	mu        sync.Mutex
	token     *OAuth2Token
	expiresAt time.Time
}

// NewInMemoryTokenCache returns an empty cache.
func NewInMemoryTokenCache() *InMemoryTokenCache { return &InMemoryTokenCache{} }

// Get satisfies TokenCache.
func (c *InMemoryTokenCache) Get(_ context.Context) (*OAuth2Token, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.token != nil && time.Now().UTC().Before(c.expiresAt) {
		return c.token, nil
	}
	return nil, nil
}

// Set satisfies TokenCache.
func (c *InMemoryTokenCache) Set(_ context.Context, token *OAuth2Token) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.token = token
	c.expiresAt = time.Now().UTC().Add(time.Duration(token.ExpiresIn) * time.Second)
	return nil
}

// FileTokenCache stores the token as JSON at FilePath. The file is written
// with mode 0600 and contains the token plus the absolute UTC expiry.
type FileTokenCache struct {
	FilePath string
}

// NewFileTokenCache constructs a cache backed by path. The file is created on
// first Set call; missing parent directories are created with mode 0700.
func NewFileTokenCache(path string) *FileTokenCache { return &FileTokenCache{FilePath: path} }

type cachedToken struct {
	Token     *OAuth2Token `json:"oauth2_token"`
	ExpiresAt time.Time    `json:"expire_at"`
}

// Get satisfies TokenCache.
func (c *FileTokenCache) Get(_ context.Context) (*OAuth2Token, error) {
	raw, err := os.ReadFile(c.FilePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	var cached cachedToken
	if err := json.Unmarshal(raw, &cached); err != nil {
		// Treat corrupt files as cache misses — the next request will refresh.
		return nil, nil
	}
	if cached.Token != nil && time.Now().UTC().Before(cached.ExpiresAt) {
		return cached.Token, nil
	}
	return nil, nil
}

// Set satisfies TokenCache.
func (c *FileTokenCache) Set(_ context.Context, token *OAuth2Token) error {
	if dir := filepath.Dir(c.FilePath); dir != "" {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return err
		}
	}
	cached := cachedToken{
		Token:     token,
		ExpiresAt: time.Now().UTC().Add(time.Duration(token.ExpiresIn) * time.Second),
	}
	raw, err := json.Marshal(cached)
	if err != nil {
		return err
	}
	return os.WriteFile(c.FilePath, raw, 0o600)
}
