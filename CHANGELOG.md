# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.1.0] - 2026-05-02

Initial Go port of `dotnet.garmin.connect`. Supports OAuth1 + OAuth2 hybrid
SSO with MFA, file-based token caching, and the full `IGarminConnectClient`
surface.

### Supported methods

#### Activities
- `GetActivities(start, limit)`
- `GetActivitiesByDate(startDate, endDate, activityType)`
- `GetActivityDetails(activityID, maxChartSize, maxPolylineSize)`
- `GetActivitySplits(activityID)`
- `GetActivitySplitSummaries(activityID)`
- `GetActivityWeather(activityID)`
- `GetActivityHrInTimezones(activityID)`
- `GetActivityExerciseSets(activityID)`
- `DownloadActivity(activityID, format)` — TCX / GPX / KML / CSV / FIT-zip
- `UploadFile(path)` / `UploadFileFromReader(filename, reader)` — `.fit` / `.tcx` / `.gpx` / `.csv` / `.xls` / `.xlsx`

#### Calendar
- `GetCalendarByYear(year)`
- `GetCalendarByMonth(year, month)`
- `GetCalendarByWeek(day)`

#### Devices
- `GetDevices()`
- `GetDeviceSettings(deviceID)`
- `GetDeviceLastUsed()`
- `GetDeviceMessages()`
- `SendWorkoutToDevices(workoutID, deviceIDs)`

#### Gear
- `GetGearTypes()`
- `GetUserGears(userID)`
- `GetActivityGears(activityID)`

#### Owner / profile
- `GetSocialProfile()` (cached after first call)
- `GetUserSettings()`
- `GetPersonalRecord(ownerDisplayName)`
- `SetUserWeight(weight)` (deprecated, prefer `AddWeight`)
- `SetUserSleepTimes(sleep, wake)`

#### Wellness
- `GetUserSummary(day)`
- `GetWellnessStepsData(day)`
- `GetWellnessHeartRates(day)`
- `GetWellnessSleepData(day)`
- `GetWellnessBodyBatteryData(startDate, endDate)`
- `GetHydrationData(day)`
- `GetBodyComposition(startDate, endDate)`

#### Workouts
- `GetWorkouts(params)` — pageable via `WorkoutsParameters`
- `GetWorkout(workoutID)`
- `GetWorkoutTypes()`
- `UpdateWorkout(workout)`
- `ScheduleWorkout(workoutID, date)`
- `RemoveScheduledWorkout(calendarID)`

#### HRV
- `GetReportHrvStatus(startDate, endDate)`

#### Weight
- `GetWeightRange(startDate, endDate)`
- `AddWeight(weight)` — local validation against 1–453 kg before submit
- `RemoveWeight(identifier)`

#### Blood pressure
- `GetBloodPressureRange(startDate, endDate)`
- `GetBloodPressureDaily(day)`
- `AddBloodPressure(measurement)` — local validation: diastolic 30–200, systolic 40–300, pulse 1–300
- `RemoveBloodPressure(identifier)`

### Auth
- `auth.NewBasicAuth(login, password)` with optional `WithDomain` (for
  `garmin.cn`), `WithUserAgent`, `WithConsumer`.
- `auth.MfaCodeProvider` — built-ins: `NotImplementedMfa`, `StaticMfa`,
  `MfaCodeFunc`.
- `auth.TokenCache` — built-ins: `InMemoryTokenCache` (default),
  `FileTokenCache(path)`.
- `*auth.AuthenticationError` carries an `AuthErrorCode` matching every
  step of the SSO flow.

### Errors
- `*garmin.RequestError` — non-2xx HTTP response with URL + status + body.
- `garmin.TooManyRequestsError` — HTTP 429.

### Notes
- Module path `github.com/sealbro/go-garmin-connect`; version exposed as
  `garmin.Version`.
- `go test ./...` runs offline unit tests.
- `go test -tags=integration ./integration/...` runs the live integration
  suite (mirrors C# `[Collection("Garmin Integrations")]`); set
  `GARMIN_RUN_DESTRUCTIVE=1` to opt into state-mutating cases.
