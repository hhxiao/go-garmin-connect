package auth

import "fmt"

// AuthErrorCode tags an AuthenticationError with the step that failed. Mirrors
// the C# `Code` enum so log messages and tests on either side line up.
type AuthErrorCode uint8

const (
	CodeNone                 AuthErrorCode = 0
	CodeCookiesNotFound      AuthErrorCode = 1
	CodeCsrfTokenNotFound    AuthErrorCode = 2
	CodeOAuth1TicketNotFound AuthErrorCode = 3
	CodeOAuth1TokenNotFound  AuthErrorCode = 4
	CodeOAuth2TokenNotFound  AuthErrorCode = 5
	CodeMfaBlockedCloudflare AuthErrorCode = 6
	CodeMfaInvalidCode       AuthErrorCode = 7
)

// AuthenticationError describes a failure during the Garmin SSO flow.
// Wrap-aware: callers can `errors.As` to inspect Code.
type AuthenticationError struct {
	Code    AuthErrorCode
	Message string
	Err     error
}

// Error satisfies the error interface.
func (e *AuthenticationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap exposes the underlying cause for errors.Is/As.
func (e *AuthenticationError) Unwrap() error { return e.Err }
