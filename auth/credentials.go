package auth

// ConsumerCredentials are the OAuth1 consumer key/secret used by the Garmin
// mobile apps. Callers can override DefaultConsumer if Garmin rotates them
// (the values are public and change every few years; the original source is
// `https://thegarth.s3.amazonaws.com/oauth_consumer.json`).
type ConsumerCredentials struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

// DefaultConsumer is the OAuth1 consumer pair as of the time this library was
// written. Mirrors the value embedded in the C# implementation.
var DefaultConsumer = ConsumerCredentials{
	ConsumerKey:    "fc3e99d2-118c-44b8-8ae3-03370dde24c0",
	ConsumerSecret: "E08WAR897WEy2knn7aFBrvegVAf0AFdWBBF",
}
