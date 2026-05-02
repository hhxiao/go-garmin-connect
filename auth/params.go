package auth

import (
	"errors"
	"net/url"
)

// AuthParameters bundles the per-account knobs required to drive the Garmin
// SSO flow: domain, default headers/query/form params for the SSO endpoints,
// and a handle for the rotating cookie + CSRF state.
//
// BasicAuth is the default implementation. Custom implementations can swap
// the domain (e.g. `garmin.cn`) or inject extra headers.
type AuthParameters interface {
	UserAgent() string
	Domain() string
	BaseURL() string

	Cookies() string
	SetCookies(string)
	Csrf() string
	SetCsrf(string)

	ConsumerCredentials() ConsumerCredentials

	// Headers, FormParameters, QueryParameters and MfaParameters are merged
	// into outbound auth requests. Modifying the returned values must not
	// mutate the implementation's internal state.
	Headers() map[string]string
	FormParameters() url.Values
	QueryParameters() url.Values
	MfaParameters() url.Values
}

// BasicAuth implements AuthParameters with username/password credentials. It
// is safe to share between goroutines as long as the caller does not mutate
// it concurrently with auth requests; the auth Service serializes refreshes.
type BasicAuth struct {
	Email     string
	Password  string
	Agent     UserAgent
	Consumer  ConsumerCredentials
	domain    string
	cookies   string
	csrfToken string
}

// NewBasicAuth validates the credentials and returns a ready-to-use param
// set. domain defaults to "garmin.com"; pass an explicit value for the
// `garmin.cn` regional endpoint.
func NewBasicAuth(email, password string) (*BasicAuth, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}
	return &BasicAuth{
		Email:    email,
		Password: password,
		Agent:    NewStaticUserAgent(""),
		Consumer: DefaultConsumer,
		domain:   "garmin.com",
	}, nil
}

// WithDomain overrides the default `garmin.com` domain — use `garmin.cn` for
// users registered against the China region.
func (b *BasicAuth) WithDomain(domain string) *BasicAuth { b.domain = domain; return b }

// WithUserAgent overrides the User-Agent header.
func (b *BasicAuth) WithUserAgent(ua UserAgent) *BasicAuth { b.Agent = ua; return b }

// WithConsumer overrides the OAuth1 consumer credentials.
func (b *BasicAuth) WithConsumer(c ConsumerCredentials) *BasicAuth { b.Consumer = c; return b }

// UserAgent returns the User-Agent string.
func (b *BasicAuth) UserAgent() string { return b.Agent.New() }

// Domain returns the SSO/API domain (e.g. "garmin.com").
func (b *BasicAuth) Domain() string { return b.domain }

// BaseURL returns the host used for authenticated API calls.
func (b *BasicAuth) BaseURL() string { return "https://connect." + b.domain }

// Cookies returns the rotating cookie string captured during auth.
func (b *BasicAuth) Cookies() string { return b.cookies }

// SetCookies updates the cookie jar; called by the auth Service.
func (b *BasicAuth) SetCookies(c string) { b.cookies = c }

// Csrf returns the rotating CSRF token captured during auth.
func (b *BasicAuth) Csrf() string { return b.csrfToken }

// SetCsrf updates the CSRF token; called by the auth Service.
func (b *BasicAuth) SetCsrf(c string) { b.csrfToken = c }

// ConsumerCredentials returns the OAuth1 consumer credentials.
func (b *BasicAuth) ConsumerCredentials() ConsumerCredentials { return b.Consumer }

// Headers returns the static headers attached to every SSO request.
func (b *BasicAuth) Headers() map[string]string {
	h := map[string]string{
		"User-Agent": b.UserAgent(),
		"origin":     "https://sso." + b.domain,
	}
	if b.cookies != "" {
		h["cookie"] = b.cookies
	}
	return h
}

// FormParameters returns the body posted to the SSO signin endpoint.
func (b *BasicAuth) FormParameters() url.Values {
	return url.Values{
		"embed":    {"true"},
		"_csrf":    {b.csrfToken},
		"username": {b.Email},
		"password": {b.Password},
	}
}

// QueryParameters returns the embed parameters that decorate the SSO URLs.
func (b *BasicAuth) QueryParameters() url.Values {
	return url.Values{
		"id":          {"gauth-widget"},
		"embedWidget": {"true"},
	}
}

// MfaParameters returns the body posted to the MFA verification endpoint
// (without the actual code, which is added by the auth Service).
func (b *BasicAuth) MfaParameters() url.Values {
	return url.Values{
		"embed":    {"true"},
		"fromPage": {"setupEnterMfaCode"},
		"_csrf":    {b.csrfToken},
	}
}
