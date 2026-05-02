//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	garmin "github.com/sealbro/go-garmin-connect"
)

func TestGetBloodPressureRange_NotEmptyOrSkip(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	start := time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)
	end := time.Now().AddDate(1, 0, 0)
	r, err := c.GetBloodPressureRange(ctx, start, end)
	if err != nil {
		t.Fatal(err)
	}
	if len(r) == 0 {
		t.Skip("no blood-pressure history")
	}
}

func TestGetBloodPressureDaily_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	start := time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)
	end := time.Now().AddDate(1, 0, 0)
	r, err := c.GetBloodPressureRange(ctx, start, end)
	if err != nil {
		t.Fatal(err)
	}
	if len(r) == 0 {
		t.Skip("no blood-pressure history")
	}

	first := r[0]
	d, err := c.GetBloodPressureDaily(ctx, first.MeasurementTimestampLocal.Time)
	if err != nil {
		t.Fatal(err)
	}
	if len(d.BloodPressureMeasurements) == 0 {
		t.Fatal("expected at least one daily measurement")
	}
}

func TestAdd_And_Remove_BloodPressure_Success(t *testing.T) {
	if !runDestructive() {
		t.Skip("set GARMIN_RUN_DESTRUCTIVE=1 to run state-mutating tests")
	}
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	now := time.Now()
	startDate := now.AddDate(0, 0, -1)
	endDate := now.AddDate(0, 0, 1)
	bp := garmin.GarminBloodPressure{
		Diastolic: 80, Systolic: 120, Pulse: 60,
		MeasurementDateTime: now, Notes: "go-garmin-connect integration test",
	}

	added, err := c.AddBloodPressure(ctx, bp)
	if err != nil {
		t.Fatal(err)
	}
	if !added {
		t.Fatal("AddBloodPressure returned false for valid input")
	}

	r, err := c.GetBloodPressureRange(ctx, startDate, endDate)
	if err != nil {
		t.Fatal(err)
	}
	if len(r) == 0 {
		t.Fatal("blood-pressure entry not found after add")
	}

	first := r[0]
	if first.Diastolic != bp.Diastolic || first.Systolic != bp.Systolic || first.Pulse != bp.Pulse {
		t.Errorf("readings mismatch: got %+v, want %+v", first, bp)
	}

	if err := c.RemoveBloodPressure(ctx, first.GarminBloodPressureIdentifier); err != nil {
		t.Fatal(err)
	}

	r, err = c.GetBloodPressureRange(ctx, startDate, endDate)
	if err != nil {
		t.Fatal(err)
	}
	for _, m := range r {
		if m.Version == first.Version && m.MeasurementTimestampLocal.Time.Equal(first.MeasurementTimestampLocal.Time) {
			t.Fatal("entry still present after RemoveBloodPressure")
		}
	}
}

func TestAddBloodPressure_FailsForOutOfRange(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cases := []struct{ d, s, p int64 }{
		{29, 100, 100},
		{201, 100, 100},
		{100, 39, 100},
		{100, 301, 100},
		{100, 100, 0},
		{100, 100, 301},
	}
	for _, c2 := range cases {
		bp := garmin.GarminBloodPressure{
			Diastolic: c2.d, Systolic: c2.s, Pulse: c2.p,
			MeasurementDateTime: time.Now(), Notes: "",
		}
		ok, err := c.AddBloodPressure(ctx, bp)
		if err != nil {
			t.Fatalf("Add(%+v) error: %v", c2, err)
		}
		if ok {
			t.Errorf("Add(%+v) returned true; expected validation failure", c2)
		}
	}
}
