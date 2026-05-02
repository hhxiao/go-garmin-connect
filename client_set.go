package garmin

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// validUploadExtensions matches the extensions the Garmin upload endpoint
// accepts. Mirrors the C# `_validExtensions` list.
var validUploadExtensions = map[string]struct{}{
	".fit": {}, ".tcx": {}, ".gpx": {},
	".csv": {}, ".xls": {}, ".xlsx": {},
}

// SetUserWeight is the deprecated weight setter that piggybacks on the
// user-settings endpoint. Prefer AddWeight.
//
// Deprecated: use AddWeight.
func (c *Client) SetUserWeight(ctx context.Context, weight float64) error {
	body := map[string]any{"userData": map[string]any{"weight": weight}}
	resp, err := c.ctx.MakeHTTPPut(ctx, urlUserSettings, body, nil)
	drainResponse(resp)
	return err
}

// SetUserSleepTimes overrides the configured sleep/wake window. Pass nil to
// reset that side back to the device default.
func (c *Client) SetUserSleepTimes(ctx context.Context, sleepTime, wakeTime *int64) error {
	sleep := UserSleep{
		DefaultSleepTime: sleepTime == nil,
		DefaultWakeTime:  wakeTime == nil,
	}
	if sleepTime != nil {
		sleep.SleepTime = *sleepTime
	}
	if wakeTime != nil {
		sleep.WakeTime = *wakeTime
	}
	body := map[string]any{"userSleep": sleep}
	resp, err := c.ctx.MakeHTTPPut(ctx, urlUserSettings, body, nil)
	drainResponse(resp)
	return err
}

// UpdateWorkout overwrites an existing workout.
func (c *Client) UpdateWorkout(ctx context.Context, workout GarminWorkout) error {
	if workout.WorkoutID == 0 {
		return errors.New("WorkoutId must be from existing workout")
	}
	headers := map[string]string{"X-Http-Method-Override": "PUT"}
	resp, err := c.ctx.MakeHTTPPost(ctx,
		fmt.Sprintf("%s%d", urlWorkout, workout.WorkoutID),
		newUpdateWorkoutPayload(workout),
		headers,
	)
	drainResponse(resp)
	return err
}

// ScheduleWorkout puts an existing workout on the calendar for date.
func (c *Client) ScheduleWorkout(ctx context.Context, workoutID int64, date time.Time) error {
	body := map[string]any{"date": date.Format(dateLayout)}
	resp, err := c.ctx.MakeHTTPPost(ctx, fmt.Sprintf("%s%d", urlWorkoutSchedule, workoutID), body, nil)
	drainResponse(resp)
	return err
}

// RemoveScheduledWorkout deletes a calendar entry by its ID.
func (c *Client) RemoveScheduledWorkout(ctx context.Context, calendarID int64) error {
	resp, err := c.ctx.MakeHTTPDelete(ctx, fmt.Sprintf("%s%d", urlWorkoutSchedule, calendarID), nil)
	drainResponse(resp)
	return err
}

// SendWorkoutToDevices queues a workout for delivery to the listed devices.
// All deviceIDs must belong to the account or the call fails before contact
// with Garmin.
func (c *Client) SendWorkoutToDevices(ctx context.Context, workoutID int64, deviceIDs []int64) error {
	if workoutID == 0 {
		return errors.New("WorkoutId must be from existing workout")
	}
	if len(deviceIDs) == 0 {
		return errors.New("DeviceIds must be not empty")
	}

	workout, err := c.GetWorkout(ctx, workoutID)
	if err != nil {
		return err
	}
	devices, err := c.GetDevices(ctx)
	if err != nil {
		return err
	}
	if workout.WorkoutID != workoutID {
		return errors.New("workout not found")
	}

	deviceByID := make(map[int64]GarminDevice, len(devices))
	for _, d := range devices {
		deviceByID[d.DeviceID] = d
	}

	var notFound []int64
	for _, id := range deviceIDs {
		if _, ok := deviceByID[id]; !ok {
			notFound = append(notFound, id)
		}
	}
	if len(notFound) > 0 {
		return fmt.Errorf("devices [%v] not found", notFound)
	}

	type sendToDevice struct {
		DeviceID    int64       `json:"deviceId"`
		MessageURL  string      `json:"messageUrl"`
		MessageType string      `json:"messageType"`
		MessageName string      `json:"messageName"`
		GroupName   interface{} `json:"groupName"`
		Priority    int64       `json:"priority"`
		FileType    string      `json:"fileType"`
		MetaDataID  int64       `json:"metaDataId"`
	}
	payload := make([]sendToDevice, 0, len(deviceIDs))
	for _, id := range deviceIDs {
		payload = append(payload, sendToDevice{
			DeviceID:    id,
			MessageURL:  fmt.Sprintf("workout-service/workout/FIT/%d", workoutID),
			MessageType: "workouts",
			MessageName: workout.WorkoutName,
			Priority:    1,
			FileType:    "FIT",
			MetaDataID:  workoutID,
		})
	}
	resp, err := c.ctx.MakeHTTPPost(ctx, urlDeviceMessage, payload, nil)
	drainResponse(resp)
	return err
}

// AddBloodPressure stores one blood-pressure reading. Returns false if the
// reading fails the local plausibility checks before the HTTP call is made.
func (c *Client) AddBloodPressure(ctx context.Context, data GarminBloodPressure) (bool, error) {
	if !data.IsValid() {
		return false, nil
	}
	rounded := time.Date(
		data.MeasurementDateTime.Year(),
		data.MeasurementDateTime.Month(),
		data.MeasurementDateTime.Day(),
		data.MeasurementDateTime.Hour(),
		data.MeasurementDateTime.Minute(),
		0, 0, data.MeasurementDateTime.Location())

	body := map[string]any{
		"version":                   nil,
		"systolic":                  data.Systolic,
		"diastolic":                 data.Diastolic,
		"pulse":                     data.Pulse,
		"multiMeasurement":          false,
		"notes":                     data.Notes,
		"sourceType":                "MANUAL",
		"measurementTimestampLocal": rounded.Format("2006-01-02T15:04:05.000"),
		"measurementTimestampGMT":   rounded.UTC().Format("2006-01-02T15:04:05.000"),
		"category":                  nil,
		"categoryName":              nil,
	}
	resp, err := c.ctx.MakeHTTPPost(ctx, urlBloodPressure, body, nil)
	drainResponse(resp)
	if err != nil {
		return false, err
	}
	return true, nil
}

// RemoveBloodPressure deletes a stored blood-pressure measurement by its
// (timestamp, version) identifier.
func (c *Client) RemoveBloodPressure(ctx context.Context, id GarminBloodPressureIdentifier) error {
	url := fmt.Sprintf("%s/%s/%d", urlBloodPressure, id.MeasurementTimestampGmt.Format(dateLayout), id.Version)
	resp, err := c.ctx.MakeHTTPDelete(ctx, url, nil)
	drainResponse(resp)
	return err
}

// AddWeight stores one weight measurement. Returns false if the reading
// fails the local plausibility checks (1–453 kg).
func (c *Client) AddWeight(ctx context.Context, data GarminWeight) (bool, error) {
	if !data.IsValid() {
		return false, nil
	}
	rounded := time.Date(
		data.MeasurementDateTime.Year(),
		data.MeasurementDateTime.Month(),
		data.MeasurementDateTime.Day(),
		data.MeasurementDateTime.Hour(),
		data.MeasurementDateTime.Minute(),
		0, 0, data.MeasurementDateTime.Location())

	body := map[string]any{
		"dateTimestamp": rounded.Format("2006-01-02T15:04:05.000"),
		"gmtTimestamp":  rounded.UTC().Format("2006-01-02T15:04:05.000"),
		"unitKey":       strings.ToLower(data.UnitKey.String()),
		"value":         data.Value,
	}
	resp, err := c.ctx.MakeHTTPPost(ctx, urlWeight, body, nil)
	drainResponse(resp)
	if err != nil {
		return false, err
	}
	return true, nil
}

// RemoveWeight deletes a stored weight measurement by its (date, samplePk)
// identifier.
func (c *Client) RemoveWeight(ctx context.Context, id GarminWeightIdentifier) error {
	url := fmt.Sprintf("%s/%s/byversion/%d", urlWeight, id.CalendarDate.Format(dateLayout), id.SamplePk)
	resp, err := c.ctx.MakeHTTPDelete(ctx, url, nil)
	drainResponse(resp)
	return err
}

// UploadFile streams the file at filepath to Garmin Connect. The extension
// determines the upload sub-URL — accepted: .fit/.tcx/.gpx/.csv/.xls/.xlsx.
func (c *Client) UploadFile(ctx context.Context, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return c.UploadFileFromReader(ctx, filepath.Base(filePath), f)
}

// UploadFileFromReader is the streaming variant of UploadFile.
func (c *Client) UploadFileFromReader(ctx context.Context, filename string, r io.Reader) error {
	ext := strings.ToLower(filepath.Ext(filename))
	if _, ok := validUploadExtensions[ext]; !ok {
		var allowed []string
		for k := range validUploadExtensions {
			allowed = append(allowed, k)
		}
		return fmt.Errorf("file extension %s is not supported. supported: %s", ext, strings.Join(allowed, ", "))
	}

	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	header := make(map[string][]string)
	header["Content-Disposition"] = []string{
		fmt.Sprintf(`form-data; name="userfile"; filename=%q`, filepath.Base(filename)),
	}
	header["Content-Type"] = []string{"application/octet-stream"}
	part, err := mw.CreatePart(header)
	if err != nil {
		return err
	}
	if _, err := io.Copy(part, r); err != nil {
		return err
	}
	if err := mw.Close(); err != nil {
		return err
	}

	url := urlUpload + "/" + ext
	resp, err := c.ctx.MakeHTTPRequest(ctx, http.MethodPost, url, nil, &buf, mw.FormDataContentType())
	drainResponse(resp)
	return err
}

// jsonRoundtrip is a tiny helper that re-encodes a value through the
// standard library — used in tests to spot wire-format diffs against
// Garmin's expected payload shape.
func jsonRoundtrip(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Force compile-time use of helper to avoid an unused-import warning when
// tests are stripped. The helper is internal and the encoder dependency
// would otherwise float unused.
var _ = jsonRoundtrip
