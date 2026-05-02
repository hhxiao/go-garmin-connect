package garmin

import "fmt"

// RequestError is returned when an authenticated API call returns a non-2xx
// status that isn't 429 (rate limit) or 401/403 (which trigger a token refresh).
type RequestError struct {
	URL     string
	Status  int
	Details string
}

// Error satisfies the error interface.
func (e *RequestError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("request [%s] returned status %d: %s", e.URL, e.Status, e.Details)
	}
	return fmt.Sprintf("request [%s] returned status %d", e.URL, e.Status)
}

// TooManyRequestsError is returned for HTTP 429. Callers should back off and
// retry; the original library does not embed any Retry-After hint because
// Garmin doesn't send one.
type TooManyRequestsError struct{}

// Error satisfies the error interface.
func (TooManyRequestsError) Error() string { return "too many requests; try again later" }

// UnexpectedError is returned when a model invariant fails — usually because
// Garmin changed a response shape and a required field is missing.
type UnexpectedError struct{ Property string }

// Error satisfies the error interface.
func (e *UnexpectedError) Error() string {
	return fmt.Sprintf("model changed: %s not found", e.Property)
}
