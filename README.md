# go-garmin-connect

Unofficial Go client for the Garmin Connect web API. Ported from
[`dotnet.garmin.connect`](https://github.com/sealbro/dotnet.garmin.connect)
and supports the same OAuth1 + OAuth2 hybrid SSO flow used by the Garmin
mobile apps, including MFA.

## Install

```bash
go get github.com/sealbro/go-garmin-connect@latest
```

The module is versioned with semver — pin a tag (`@v0.1.0`) for reproducible
builds.

## Quick start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    garmin "github.com/sealbro/go-garmin-connect"
    "github.com/sealbro/go-garmin-connect/auth"
)

func main() {
    params, err := auth.NewBasicAuth(os.Getenv("GARMIN_LOGIN"), os.Getenv("GARMIN_PASSWORD"))
    if err != nil { log.Fatal(err) }

    cache := auth.NewFileTokenCache(os.ExpandEnv("$HOME/.garmin_token.json"))
    client := garmin.NewClient(garmin.NewContext(nil, params, garmin.WithTokenCache(cache)))

    ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
    defer cancel()

    profile, err := client.GetSocialProfile(ctx)
    if err != nil { log.Fatal(err) }
    fmt.Println(profile.DisplayName)
}
```

A complete example is in `examples/basic/main.go`.

## MFA

Pass an MFA provider to `NewContext`:

```go
mfa := auth.MfaCodeFunc(func(ctx context.Context) (string, error) {
    fmt.Print("MFA code: ")
    var code string
    _, err := fmt.Scanln(&code)
    return code, err
})

ctx := garmin.NewContext(nil, params, garmin.WithMfaProvider(mfa))
```

For non-interactive runs supply `auth.StaticMfa{Code: "123456"}`. The default
is `auth.NotImplementedMfa`, which fails fast with `auth.ErrMfaNotImplemented`
if Garmin asks for a code.

## Token caching

`auth.InMemoryTokenCache` (default) keeps tokens for the lifetime of the
process. `auth.NewFileTokenCache(path)` persists them to JSON so subsequent
runs skip the SSO flow until the token expires. Implement `auth.TokenCache`
to plug in a keychain, Vault, or shared store.

## Endpoints

The `Client` exposes the same surface as the C# library:

- Activities — `GetActivities`, `GetActivitiesByDate`, `GetActivityDetails`,
  `GetActivitySplits`, `GetActivitySplitSummaries`, `GetActivityWeather`,
  `GetActivityHrInTimezones`, `GetActivityExerciseSets`, `DownloadActivity`,
  `UploadFile`, `UploadFileFromReader`
- Calendar — `GetCalendarByYear`, `GetCalendarByMonth`, `GetCalendarByWeek`
- Devices — `GetDevices`, `GetDeviceSettings`, `GetDeviceLastUsed`,
  `GetDeviceMessages`, `SendWorkoutToDevices`
- Gear — `GetGearTypes`, `GetUserGears`, `GetActivityGears`
- Owner — `GetSocialProfile`, `GetUserSettings`, `GetPersonalRecord`
- Wellness — `GetWellnessStepsData`, `GetUserSummary`,
  `GetWellnessHeartRates`, `GetWellnessSleepData`, `GetBodyComposition`,
  `GetHydrationData`, `GetWellnessBodyBatteryData`
- Workouts — `GetWorkout`, `GetWorkouts`, `GetWorkoutTypes`, `UpdateWorkout`,
  `ScheduleWorkout`, `RemoveScheduledWorkout`
- HRV — `GetReportHrvStatus`
- Weight — `GetWeightRange`, `AddWeight`, `RemoveWeight`, `SetUserWeight`
  (deprecated)
- Blood pressure — `GetBloodPressureDaily`, `GetBloodPressureRange`,
  `AddBloodPressure`, `RemoveBloodPressure`
- Sleep — `SetUserSleepTimes`

## Errors

- `*garmin.RequestError` — non-2xx HTTP response (with URL + status + body)
- `garmin.TooManyRequestsError` — HTTP 429
- `*auth.AuthenticationError` — SSO flow failure (`Code` identifies which
  step broke)

Use `errors.As` to inspect.

## Testing

```bash
# Fast unit tests — no credentials needed.
go test ./...

# Live integration tests against Garmin Connect. Requires valid creds.
GARMIN_LOGIN=you@example.com GARMIN_PASSWORD=… \
  go test -tags=integration ./integration/...

# Include destructive tests that mutate account state (sleep window,
# update workout, schedule workout, add/remove weight & blood pressure).
GARMIN_RUN_DESTRUCTIVE=1 GARMIN_LOGIN=… GARMIN_PASSWORD=… \
  go test -tags=integration ./integration/...
```

The `integration` build tag mirrors the C# `[Collection("Garmin Integrations")]`
suite. Tokens are cached at `$HOME/.garmin_token.json` so re-runs skip the
SSO flow until expiry.

## License

MIT — see [LICENSE](LICENSE).
