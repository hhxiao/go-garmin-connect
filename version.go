// Package garmin is an unofficial Go client for the Garmin Connect web API.
//
// It is a port of the C# library https://github.com/sealbro/dotnet.garmin.connect
// and supports the same OAuth1 + OAuth2 hybrid authentication flow used by the
// official Garmin mobile apps, including MFA.
//
// Versioning follows semver and is exposed as Version. Tag the module with
// `vX.Y.Z` git tags so consumers can `go get` a specific version.
package garmin

// Version is the released version of this module.
//
// Update before tagging a release; the value is exposed for consumers that
// want to log or surface it (it is not used internally).
const Version = "0.1.0"
