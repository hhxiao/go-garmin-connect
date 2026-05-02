package auth

import (
	"context"
	"errors"
)

// MfaCodeProvider supplies a TOTP/SMS code when Garmin's SSO challenges with
// MFA. Implementations may block on user input, read from an env var, or
// look up a Yubikey — they own the policy.
type MfaCodeProvider interface {
	// GetMfaCode is called once per auth attempt that hits an MFA prompt.
	// Returning a non-nil error aborts authentication; an empty string is
	// treated as a missing code and produces an MfaInvalidCode error.
	GetMfaCode(ctx context.Context) (string, error)
}

// ErrMfaNotImplemented is returned by NotImplementedMfa to signal that the
// caller never wired up an MFA provider but the account requires one.
var ErrMfaNotImplemented = errors.New("MfaCodeProvider is not implemented")

// NotImplementedMfa always errors. It is the default for accounts that don't
// have MFA enabled — using it on an MFA-enabled account fails fast.
type NotImplementedMfa struct{}

// GetMfaCode satisfies MfaCodeProvider.
func (NotImplementedMfa) GetMfaCode(context.Context) (string, error) {
	return "", ErrMfaNotImplemented
}

// StaticMfa returns a fixed code. Useful for tests where the code is known
// up front.
type StaticMfa struct{ Code string }

// GetMfaCode satisfies MfaCodeProvider.
func (s StaticMfa) GetMfaCode(context.Context) (string, error) { return s.Code, nil }

// MfaCodeFunc adapts a plain function into a MfaCodeProvider.
type MfaCodeFunc func(ctx context.Context) (string, error)

// GetMfaCode satisfies MfaCodeProvider.
func (f MfaCodeFunc) GetMfaCode(ctx context.Context) (string, error) { return f(ctx) }
