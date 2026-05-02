package garmin

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is the high-level entry point for the Garmin Connect API. It wraps
// a token-refreshing Context and exposes one method per documented API call.
//
// Construct a Client with NewClient, supply credentials via the Context, and
// pass a context.Context to every method for cancellation/timeout control.
type Client struct {
	ctx *Context
}

// NewClient builds a Client from a Context. The Context owns the http.Client
// and the auth state — share it across multiple Clients to coalesce token
// refreshes.
func NewClient(ctx *Context) *Client { return &Client{ctx: ctx} }

// Endpoints — kept as constants for readability and easy diff against the C#
// reference. Garmin's URL hierarchy is documented at
// https://github.com/cyberjunky/python-garminconnect.
const (
	urlUserProfile          = "/userprofile-service/socialProfile"
	urlUserSettings         = "/userprofile-service/userprofile/user-settings"
	urlUserSummary          = "/usersummary-service/usersummary/daily/"
	urlUserSummaryHydration = "/usersummary-service/usersummary/hydration/daily/"
	urlUserSummaryChart     = "/wellness-service/wellness/dailySummaryChart/"
	urlHeartRates           = "/wellness-service/wellness/dailyHeartRate/"
	urlSleepData            = "/wellness-service/wellness/dailySleepData/"
	urlWeight               = "/weight-service/weight"
	urlBodyComposition      = "/weight-service/weight/daterangesnapshot"
	urlActivities           = "/activitylist-service/activities/search/activities"
	urlActivity             = "/activity-service/activity/"
	urlPersonalRecord       = "/personalrecord-service/personalrecord/"
	urlTcxDownload          = "/download-service/export/tcx/activity/"
	urlGpxDownload          = "/download-service/export/gpx/activity/"
	urlKmlDownload          = "/download-service/export/kml/activity/"
	urlFitDownload          = "/download-service/files/activity/"
	urlCsvDownload          = "/download-service/export/csv/activity/"
	urlDeviceList           = "/device-service/deviceregistration/devices"
	urlDeviceService        = "/device-service/deviceservice/"
	urlDeviceMessage        = "/device-service/devicemessage/messages"
	urlGear                 = "/gear-service/gear/"
	urlWorkout              = "/workout-service/workout/"
	urlWorkoutSchedule      = "/workout-service/schedule/"
	urlWorkouts             = "/workout-service/workouts"
	urlReportHrvStatus      = "/hrv-service/hrv/daily/"
	urlCalendarYear         = "/calendar-service/year/"
	urlBodyBattery          = "/wellness-service/wellness/bodyBattery/reports/daily/"
	urlBloodPressure        = "/bloodpressure-service/bloodpressure"
	urlUpload               = "/upload-service/upload"
)

// Activities ----------------------------------------------------------------

// GetActivities returns up to limit activities starting from start.
func (c *Client) GetActivities(ctx context.Context, start, limit int) ([]GarminActivity, error) {
	url := fmt.Sprintf("%s?start=%d&limit=%d", urlActivities, start, limit)
	return GetAndDeserialize[[]GarminActivity](ctx, c.ctx, url)
}

// GetActivitiesByDate paginates through every activity in [startDate, endDate].
// activityType is optional (one of cycling/running/swimming/multi_sport/
// fitness_equipment/hiking/walking/other).
func (c *Client) GetActivitiesByDate(ctx context.Context, startDate, endDate time.Time, activityType string) ([]GarminActivity, error) {
	const pageSize = 20
	activitySlug := ""
	if activityType != "" {
		activitySlug = "&activityType=" + activityType
	}
	var result []GarminActivity
	start := 0
	for {
		url := fmt.Sprintf("%s?startDate=%s&endDate=%s&start=%d&limit=%d%s",
			urlActivities,
			startDate.Format(dateLayout),
			endDate.Format(dateLayout),
			start, pageSize, activitySlug)
		page, err := GetAndDeserialize[[]GarminActivity](ctx, c.ctx, url)
		if err != nil {
			return nil, err
		}
		if len(page) == 0 {
			break
		}
		result = append(result, page...)
		start += pageSize
	}
	return result, nil
}

// GetActivityExerciseSets returns exercise-set details for a strength activity.
func (c *Client) GetActivityExerciseSets(ctx context.Context, activityID int64) (GarminExerciseSets, error) {
	return GetAndDeserialize[GarminExerciseSets](ctx, c.ctx, fmt.Sprintf("%s%d", urlActivity, activityID))
}

// GetActivitySplits returns the per-lap breakdown for an activity.
func (c *Client) GetActivitySplits(ctx context.Context, activityID int64) (GarminActivitySplits, error) {
	return GetAndDeserialize[GarminActivitySplits](ctx, c.ctx, fmt.Sprintf("%s%d/splits", urlActivity, activityID))
}

// GetActivitySplitSummaries returns the split-summaries record.
func (c *Client) GetActivitySplitSummaries(ctx context.Context, activityID int64) (GarminSplitSummary, error) {
	return GetAndDeserialize[GarminSplitSummary](ctx, c.ctx, fmt.Sprintf("%s%d/split_summaries", urlActivity, activityID))
}

// GetActivityWeather returns the weather snapshot for an activity.
func (c *Client) GetActivityWeather(ctx context.Context, activityID int64) (GarminActivityWeather, error) {
	return GetAndDeserialize[GarminActivityWeather](ctx, c.ctx, fmt.Sprintf("%s%d/weather", urlActivity, activityID))
}

// GetActivityHrInTimezones returns the time-in-zone breakdown for an activity.
func (c *Client) GetActivityHrInTimezones(ctx context.Context, activityID int64) ([]GarminHrTimeInZones, error) {
	return GetAndDeserialize[[]GarminHrTimeInZones](ctx, c.ctx, fmt.Sprintf("%s%d/hrTimeInZones", urlActivity, activityID))
}

// GetActivityDetails returns the dense per-sample metric series for an
// activity. maxChartSize/maxPolylineSize control the down-sampling Garmin
// applies (default values mirror the C# library: 2000 / 4000).
func (c *Client) GetActivityDetails(ctx context.Context, activityID int64, maxChartSize, maxPolylineSize int) (GarminActivityDetails, error) {
	if maxChartSize <= 0 {
		maxChartSize = 2000
	}
	if maxPolylineSize <= 0 {
		maxPolylineSize = 4000
	}
	url := fmt.Sprintf("%s%d/details?maxChartSize=%d&maxPolylineSize=%d", urlActivity, activityID, maxChartSize, maxPolylineSize)
	return GetAndDeserialize[GarminActivityDetails](ctx, c.ctx, url)
}

// DownloadActivity retrieves an activity in the requested format. For
// FormatOriginal the bytes are a zip — callers must extract.
func (c *Client) DownloadActivity(ctx context.Context, activityID int64, format ActivityDownloadFormat) ([]byte, error) {
	var url string
	switch format {
	case FormatOriginal:
		url = fmt.Sprintf("%s%d", urlFitDownload, activityID)
	case FormatTCX:
		url = fmt.Sprintf("%s%d", urlTcxDownload, activityID)
	case FormatGPX:
		url = fmt.Sprintf("%s%d", urlGpxDownload, activityID)
	case FormatKML:
		url = fmt.Sprintf("%s%d", urlKmlDownload, activityID)
	case FormatCSV:
		url = fmt.Sprintf("%s%d", urlCsvDownload, activityID)
	default:
		return nil, fmt.Errorf("unexpected download format %d", format)
	}
	resp, err := c.ctx.MakeHTTPGet(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// Calendar ------------------------------------------------------------------

// GetCalendarByYear returns the year-overview heatmap.
func (c *Client) GetCalendarByYear(ctx context.Context, year int) (GarminCalendarYear, error) {
	return GetAndDeserialize[GarminCalendarYear](ctx, c.ctx, fmt.Sprintf("%s%d", urlCalendarYear, year))
}

// GetCalendarByMonth returns the month overview. month is the 0-based index
// Garmin uses (use NewGarminMonth to convert from time.Month).
func (c *Client) GetCalendarByMonth(ctx context.Context, year int, month GarminMonth) (GarminCalendarMonth, error) {
	url := fmt.Sprintf("%s%d/month/%d", urlCalendarYear, year, int(month))
	return GetAndDeserialize[GarminCalendarMonth](ctx, c.ctx, url)
}

// GetCalendarByWeek returns the calendar week containing the given day.
// Pass any day from the desired week.
func (c *Client) GetCalendarByWeek(ctx context.Context, day time.Time) (GarminCalendarWeek, error) {
	url := fmt.Sprintf("%s%d/month/%d/day/%d/start/1",
		urlCalendarYear, day.Year(), int(day.Month())-1, day.Day())
	return GetAndDeserialize[GarminCalendarWeek](ctx, c.ctx, url)
}

// Devices -------------------------------------------------------------------

// GetDevices returns every device registered to the account.
func (c *Client) GetDevices(ctx context.Context) ([]GarminDevice, error) {
	return GetAndDeserialize[[]GarminDevice](ctx, c.ctx, urlDeviceList)
}

// GetDeviceSettings returns settings for one device.
func (c *Client) GetDeviceSettings(ctx context.Context, deviceID int64) (GarminDeviceSettings, error) {
	return GetAndDeserialize[GarminDeviceSettings](ctx, c.ctx, fmt.Sprintf("%sdevice-info/settings/%d", urlDeviceService, deviceID))
}

// GetDeviceLastUsed returns the most recently synced device summary.
func (c *Client) GetDeviceLastUsed(ctx context.Context) (GarminDeviceLastUsed, error) {
	return GetAndDeserialize[GarminDeviceLastUsed](ctx, c.ctx, urlDeviceService+"mylastused")
}

// GetDeviceMessages returns the device message inbox.
func (c *Client) GetDeviceMessages(ctx context.Context) (GarminDeviceMessages, error) {
	return GetAndDeserialize[GarminDeviceMessages](ctx, c.ctx, urlDeviceMessage)
}

// Gear ----------------------------------------------------------------------

// GetGearTypes returns the gear-type lookup table.
func (c *Client) GetGearTypes(ctx context.Context) ([]GarminGearType, error) {
	return GetAndDeserialize[[]GarminGearType](ctx, c.ctx, urlGear+"types")
}

// GetUserGears returns gears registered to the account.
func (c *Client) GetUserGears(ctx context.Context, userID int64) ([]GarminGear, error) {
	return GetAndDeserialize[[]GarminGear](ctx, c.ctx, fmt.Sprintf("%sfilterGear?userProfilePk=%d", urlGear, userID))
}

// GetActivityGears returns the gears used during an activity.
func (c *Client) GetActivityGears(ctx context.Context, activityID int64) ([]GarminGear, error) {
	return GetAndDeserialize[[]GarminGear](ctx, c.ctx, fmt.Sprintf("%sfilterGear?activityId=%d", urlGear, activityID))
}

// Owner / profile -----------------------------------------------------------

// GetSocialProfile returns the cached profile, fetching it on the first call.
// Wellness endpoints depend on this being populated because they key off
// `displayName` instead of user ID.
func (c *Client) GetSocialProfile(ctx context.Context) (GarminSocialProfile, error) {
	c.ctx.profileMu.RLock()
	if p := c.ctx.profile; p != nil {
		c.ctx.profileMu.RUnlock()
		return *p, nil
	}
	c.ctx.profileMu.RUnlock()

	p, err := GetAndDeserialize[GarminSocialProfile](ctx, c.ctx, urlUserProfile)
	if err != nil {
		return GarminSocialProfile{}, err
	}
	c.ctx.profileMu.Lock()
	c.ctx.profile = &p
	c.ctx.profileMu.Unlock()
	return p, nil
}

// GetUserSettings returns the user settings blob.
func (c *Client) GetUserSettings(ctx context.Context) (GarminUserSettings, error) {
	return GetAndDeserialize[GarminUserSettings](ctx, c.ctx, urlUserSettings)
}

// GetPersonalRecord returns personal records for the named owner. ownerDisplay
// name is the value of GarminSocialProfile.DisplayName.
func (c *Client) GetPersonalRecord(ctx context.Context, ownerDisplayName string) ([]GarminPersonalRecord, error) {
	return GetAndDeserialize[[]GarminPersonalRecord](ctx, c.ctx, fmt.Sprintf("%sprs/%s", urlPersonalRecord, ownerDisplayName))
}

// Wellness ------------------------------------------------------------------

// GetWellnessStepsData returns a 15-minute granularity step series for one day.
func (c *Client) GetWellnessStepsData(ctx context.Context, day time.Time) ([]GarminStepsData, error) {
	profile, err := c.GetSocialProfile(ctx)
	if err != nil {
		return nil, err
	}
	if profile.DisplayName == "" {
		return nil, errors.New("social profile could not be loaded")
	}
	url := fmt.Sprintf("%s%s?date=%s", urlUserSummaryChart, profile.DisplayName, day.Format(dateLayout))
	return GetAndDeserialize[[]GarminStepsData](ctx, c.ctx, url)
}

// GetUserSummary returns the daily activity/wellness summary.
func (c *Client) GetUserSummary(ctx context.Context, day time.Time) (GarminStats, error) {
	profile, err := c.GetSocialProfile(ctx)
	if err != nil {
		return GarminStats{}, err
	}
	if profile.DisplayName == "" {
		return GarminStats{}, errors.New("social profile could not be loaded")
	}
	url := fmt.Sprintf("%s%s?calendarDate=%s", urlUserSummary, profile.DisplayName, day.Format(dateLayout))
	return GetAndDeserialize[GarminStats](ctx, c.ctx, url)
}

// GetWellnessHeartRates returns the daily heart-rate series.
func (c *Client) GetWellnessHeartRates(ctx context.Context, day time.Time) (GarminHr, error) {
	profile, err := c.GetSocialProfile(ctx)
	if err != nil {
		return GarminHr{}, err
	}
	if profile.DisplayName == "" {
		return GarminHr{}, errors.New("social profile could not be loaded")
	}
	url := fmt.Sprintf("%s%s?date=%s", urlHeartRates, profile.DisplayName, day.Format(dateLayout))
	return GetAndDeserialize[GarminHr](ctx, c.ctx, url)
}

// GetWellnessSleepData returns the daily sleep summary.
func (c *Client) GetWellnessSleepData(ctx context.Context, day time.Time) (GarminSleepData, error) {
	profile, err := c.GetSocialProfile(ctx)
	if err != nil {
		return GarminSleepData{}, err
	}
	if profile.DisplayName == "" {
		return GarminSleepData{}, errors.New("social profile could not be loaded")
	}
	url := fmt.Sprintf("%s%s?date=%s", urlSleepData, profile.DisplayName, day.Format(dateLayout))
	return GetAndDeserialize[GarminSleepData](ctx, c.ctx, url)
}

// GetBodyComposition returns the multi-day body-composition summary.
func (c *Client) GetBodyComposition(ctx context.Context, startDate, endDate time.Time) (GarminBodyComposition, error) {
	url := fmt.Sprintf("%s?startDate=%s&endDate=%s", urlBodyComposition, startDate.Format(dateLayout), endDate.Format(dateLayout))
	return GetAndDeserialize[GarminBodyComposition](ctx, c.ctx, url)
}

// GetHydrationData returns the daily hydration summary.
func (c *Client) GetHydrationData(ctx context.Context, day time.Time) (GarminHydrationData, error) {
	url := fmt.Sprintf("%s%s", urlUserSummaryHydration, day.Format(dateLayout))
	return GetAndDeserialize[GarminHydrationData](ctx, c.ctx, url)
}

// GetWellnessBodyBatteryData returns the multi-day body-battery series.
func (c *Client) GetWellnessBodyBatteryData(ctx context.Context, startDate, endDate time.Time) ([]GarminBodyBatteryData, error) {
	url := fmt.Sprintf("%s?startDate=%s&endDate=%s", urlBodyBattery, startDate.Format(dateLayout), endDate.Format(dateLayout))
	return GetAndDeserialize[[]GarminBodyBatteryData](ctx, c.ctx, url)
}

// Blood pressure ------------------------------------------------------------

// GetBloodPressureDaily returns blood-pressure entries for a single day.
func (c *Client) GetBloodPressureDaily(ctx context.Context, day time.Time) (GarminBloodPressureDaily, error) {
	url := fmt.Sprintf("%s/dayview/%s", urlBloodPressure, day.Format(dateLayout))
	return GetAndDeserialize[GarminBloodPressureDaily](ctx, c.ctx, url)
}

// GetBloodPressureRange returns blood-pressure entries for [startDate, endDate].
func (c *Client) GetBloodPressureRange(ctx context.Context, startDate, endDate time.Time) ([]GarminBloodPressureMeasurement, error) {
	url := fmt.Sprintf("%s/daily/last/%s/%s", urlBloodPressure, startDate.Format(dateLayout), endDate.Format(dateLayout))
	return GetAndDeserialize[[]GarminBloodPressureMeasurement](ctx, c.ctx, url)
}

// Workouts ------------------------------------------------------------------

// GetWorkout returns one workout.
func (c *Client) GetWorkout(ctx context.Context, workoutID int64) (GarminWorkout, error) {
	return GetAndDeserialize[GarminWorkout](ctx, c.ctx, fmt.Sprintf("%s%d", urlWorkout, workoutID))
}

// GetWorkoutTypes returns the workout-types lookup tables.
func (c *Client) GetWorkoutTypes(ctx context.Context) (GarminWorkoutTypes, error) {
	return GetAndDeserialize[GarminWorkoutTypes](ctx, c.ctx, urlWorkout+"types")
}

// GetWorkouts paginates through workouts according to params.
func (c *Client) GetWorkouts(ctx context.Context, params WorkoutsParameters) ([]GarminWorkout, error) {
	url := fmt.Sprintf("%s?%s", urlWorkouts, params.queryString())
	return GetAndDeserialize[[]GarminWorkout](ctx, c.ctx, url)
}

// Reports -------------------------------------------------------------------

// GetReportHrvStatus returns the HRV-status report between two dates.
func (c *Client) GetReportHrvStatus(ctx context.Context, startDate, endDate time.Time) (GarminReportHrvStatus, error) {
	url := fmt.Sprintf("%s%s/%s", urlReportHrvStatus, startDate.Format(dateLayout), endDate.Format(dateLayout))
	return GetAndDeserialize[GarminReportHrvStatus](ctx, c.ctx, url)
}

// Weight --------------------------------------------------------------------

// GetWeightRange returns the multi-day weight summary.
func (c *Client) GetWeightRange(ctx context.Context, startDate, endDate time.Time) (GarminWeightRange, error) {
	url := fmt.Sprintf("%s/range/%s/%s?includeAll=true", urlWeight, startDate.Format(dateLayout), endDate.Format(dateLayout))
	return GetAndDeserialize[GarminWeightRange](ctx, c.ctx, url)
}

// drainResponse closes the body and discards anything left so connections can
// be reused.
func drainResponse(resp *http.Response) {
	if resp == nil || resp.Body == nil {
		return
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
}
