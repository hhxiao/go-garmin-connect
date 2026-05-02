package auth

// UserAgent supplies the `User-Agent` header sent for every Garmin auth
// request. Garmin SSO rejects values that don't look like one of their first
// party clients — DefaultUserAgent matches the iOS Connect Mobile app.
type UserAgent interface {
	New() string
}

// DefaultUserAgent matches the value embedded in the C# library. Garmin can
// add new server-side blocks if they decide to; rotate this if SSO starts
// rejecting requests.
const DefaultUserAgent = "GCM-iOS-5.7.2.1"

// StaticUserAgent returns the same string every time.
type StaticUserAgent struct {
	Value string
}

// NewStaticUserAgent uses DefaultUserAgent if v is empty.
func NewStaticUserAgent(v string) StaticUserAgent {
	if v == "" {
		v = DefaultUserAgent
	}
	return StaticUserAgent{Value: v}
}

// New satisfies UserAgent.
func (s StaticUserAgent) New() string { return s.Value }
