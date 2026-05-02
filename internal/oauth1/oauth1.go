// Package oauth1 is a minimal OAuth 1.0a HMAC-SHA1 signer tailored to the
// Garmin Connect SSO flow. Its encoding rules match the C# reference
// implementation (RestSharp-derived `UrlEncodeRelaxed`) byte-for-byte:
//
//   - encoding uses `url.QueryEscape` then percent-encodes "(", ")", and
//     "+" → "%20" so it matches RFC 3986 + the LinkedIn paren tweak the C#
//     library carries for parity with old OAuth providers
//   - parameter normalization uses `UrlEncodeStrict` (which only %-encodes
//     non-unreserved chars) for the value going into the signature base
//
// We do not depend on any third-party OAuth1 module because Garmin's signer
// must reproduce the C# byte sequence exactly or the OAuth1 token exchange
// fails. See dotnet.garmin.connect/Garmin.Connect/OAuth/OAuthTools.cs for
// the original.
package oauth1

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Credentials holds an OAuth1 consumer/token pair. Token + TokenSecret may be
// empty for the initial RequestToken step.
type Credentials struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
}

// AuthorizationHeader builds an `Authorization: OAuth …` header value for
// `method requestURL`. Query-string params from requestURL are folded into
// the signature base per RFC 5849. `protectedResource` selects the header
// shape used by Garmin's `oauth/exchange/user/2.0` call (which keeps
// `oauth_token` even when other oauth_* values are blank).
func AuthorizationHeader(method, requestURL string, c Credentials, protectedResource bool) (string, error) {
	parsed, err := url.Parse(requestURL)
	if err != nil {
		return "", fmt.Errorf("parse url: %w", err)
	}

	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	nonce, err := generateNonce()
	if err != nil {
		return "", err
	}

	authParams := []param{
		{"oauth_consumer_key", c.ConsumerKey},
		{"oauth_nonce", nonce},
		{"oauth_signature_method", "HMAC-SHA1"},
		{"oauth_timestamp", timestamp},
		{"oauth_version", "1.0"},
	}
	if protectedResource {
		// `oauth_token` is always emitted in protected-resource mode, even if empty.
		authParams = append(authParams, param{"oauth_token", c.Token})
	} else if c.Token != "" {
		authParams = append(authParams, param{"oauth_token", c.Token})
	}

	// Build signature-base parameter set: auth params + query-string params,
	// each value strict-encoded, sorted by name then value.
	all := append([]param{}, authParams...)
	for k, vs := range parsed.Query() {
		for _, v := range vs {
			all = append(all, param{k, v})
		}
	}
	for i := range all {
		all[i].value = encodeStrict(all[i].value)
	}
	sort.SliceStable(all, func(i, j int) bool {
		if all[i].name == all[j].name {
			return all[i].value < all[j].value
		}
		return all[i].name < all[j].name
	})

	var sb strings.Builder
	for i, p := range all {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p.name)
		sb.WriteByte('=')
		sb.WriteString(p.value)
	}
	normalizedParams := sb.String()

	baseURL := constructRequestURL(parsed)
	signatureBase := strings.ToUpper(method) + "&" + encodeRelaxed(baseURL) + "&" + encodeRelaxed(normalizedParams)

	key := encodeRelaxed(c.ConsumerSecret) + "&" + encodeRelaxed(c.TokenSecret)
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(signatureBase))
	signature := encodeRelaxed(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	authParams = append(authParams, param{"oauth_signature", signature})
	sort.SliceStable(authParams, func(i, j int) bool {
		return authParams[i].name < authParams[j].name
	})

	var hb strings.Builder
	hb.WriteString("OAuth ")
	first := true
	for _, p := range authParams {
		if p.name == "" || p.value == "" {
			// In protected-resource mode the C# code keeps `oauth_token` even when
			// its value is non-empty (already filtered above) — empty values are
			// dropped here regardless.
			continue
		}
		if !first {
			hb.WriteByte(',')
		}
		hb.WriteString(p.name)
		hb.WriteString(`="`)
		hb.WriteString(p.value)
		hb.WriteByte('"')
		first = false
	}
	return hb.String(), nil
}

type param struct {
	name  string
	value string
}

// encodeRelaxed mirrors C#'s OAuthTools.UrlEncodeRelaxed: standard
// percent-encoding, but additionally percent-encodes "(" and ")".
func encodeRelaxed(s string) string {
	escaped := url.QueryEscape(s)
	// QueryEscape encodes spaces as "+"; OAuth needs "%20".
	escaped = strings.ReplaceAll(escaped, "+", "%20")
	escaped = strings.ReplaceAll(escaped, "(", "%28")
	escaped = strings.ReplaceAll(escaped, ")", "%29")
	return escaped
}

// encodeStrict mirrors C#'s OAuthTools.UrlEncodeStrict: only the unreserved
// set [A-Za-z0-9-._~] passes through; everything else is %-encoded.
func encodeStrict(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') ||
			c == '-' || c == '.' || c == '_' || c == '~' {
			sb.WriteByte(c)
		} else {
			fmt.Fprintf(&sb, "%%%02X", c)
		}
	}
	return sb.String()
}

func constructRequestURL(u *url.URL) string {
	host := u.Host
	scheme := u.Scheme
	port := u.Port()
	if (scheme == "http" && port == "80") || (scheme == "https" && port == "443") {
		// Strip the default port from the host before reconstructing.
		if idx := strings.LastIndex(host, ":"); idx > 0 {
			host = host[:idx]
		}
	}
	return scheme + "://" + host + u.Path
}

func generateNonce() (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	for i, b := range buf {
		buf[i] = chars[int(b)%len(chars)]
	}
	return string(buf), nil
}
