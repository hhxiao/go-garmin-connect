package oauth1

import (
	"strings"
	"testing"
)

// TestEncodeRelaxed verifies the parens / space encoding match the C#
// reference. Garmin's signature verification rejects values that diverge.
func TestEncodeRelaxed(t *testing.T) {
	cases := map[string]string{
		"hello world":  "hello%20world",
		"foo(bar)":     "foo%28bar%29",
		"a/b":          "a%2Fb",
		"a+b":          "a%2Bb",
		"http://x.y/z": "http%3A%2F%2Fx.y%2Fz",
	}
	for in, want := range cases {
		got := encodeRelaxed(in)
		if got != want {
			t.Errorf("encodeRelaxed(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestEncodeStrict checks that only the unreserved set passes through.
func TestEncodeStrict(t *testing.T) {
	if encodeStrict("AZaz09-._~") != "AZaz09-._~" {
		t.Errorf("unreserved set should pass through")
	}
	if encodeStrict("a b") != "a%20b" {
		t.Errorf("space should be %%20")
	}
	if encodeStrict("(") != "%28" || encodeStrict(")") != "%29" {
		t.Errorf("parens should be %%28/%%29")
	}
}

// TestAuthorizationHeaderShape sanity-checks the OAuth header that goes to
// the Garmin preauthorized endpoint. We can't validate the signature value
// itself without a fixed nonce/timestamp, so we just check the wrapper.
func TestAuthorizationHeaderShape(t *testing.T) {
	header, err := AuthorizationHeader("GET",
		"https://connectapi.garmin.com/oauth-service/oauth/preauthorized?ticket=abc&login-url=https://sso.garmin.com/sso/embed&accepts-mfa-tokens=true",
		Credentials{ConsumerKey: "ck", ConsumerSecret: "cs"}, false)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"OAuth ",
		`oauth_consumer_key="ck"`,
		`oauth_signature_method="HMAC-SHA1"`,
		`oauth_signature="`,
		`oauth_nonce="`,
		`oauth_timestamp="`,
		`oauth_version="1.0"`,
	} {
		if !strings.Contains(header, want) {
			t.Errorf("header missing %q\nfull: %s", want, header)
		}
	}
}
