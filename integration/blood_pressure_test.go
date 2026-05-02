//go:build integration
// +build integration

package integration

import (
	"context"
	"fmt"
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
		t.Error(err)
		return
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
		t.Error(err)
		return
	}
	if len(r) == 0 {
		t.Skip("no blood-pressure history")
	}

	first := r[0]
	d, err := c.GetBloodPressureDaily(ctx, first.MeasurementTimestampLocal.Time)
	if err != nil {
		t.Error(err)
		return
	}
	if len(d.BloodPressureMeasurements) == 0 {
		t.Error("expected at least one daily measurement")
	}
}

func TestAdd_And_Remove_BloodPressure_Success(t *testing.T) {
	if !runDestructive() {
		t.Skip("set GARMIN_RUN_MUTATION_TESTS=1 to run state-mutating tests")
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

	t.Run("add returns true", func(t *testing.T) {
		added, err := c.AddBloodPressure(ctx, bp)
		if err != nil {
			t.Error(err)
			return
		}
		if !added {
			t.Error("AddBloodPressure returned false for valid input")
		}
	})

	t.Run("entry visible after add", func(t *testing.T) {
		r, err := c.GetBloodPressureRange(ctx, startDate, endDate)
		if err != nil {
			t.Error(err)
			return
		}
		if len(r) == 0 {
			t.Error("blood-pressure entry not found after add")
			return
		}
		first := r[0]
		if first.Diastolic != bp.Diastolic || first.Systolic != bp.Systolic || first.Pulse != bp.Pulse {
			t.Errorf("readings mismatch: got %+v, want %+v", first, bp)
		}
	})

	t.Run("remove then not found", func(t *testing.T) {
		r, err := c.GetBloodPressureRange(ctx, startDate, endDate)
		if err != nil {
			t.Error(err)
			return
		}
		if len(r) == 0 {
			t.Skip("no entries to remove")
		}
		first := r[0]
		if err := c.RemoveBloodPressure(ctx, first.GarminBloodPressureIdentifier); err != nil {
			t.Error(err)
			return
		}
		r2, err := c.GetBloodPressureRange(ctx, startDate, endDate)
		if err != nil {
			t.Error(err)
			return
		}
		for _, m := range r2 {
			if m.Version == first.Version && m.MeasurementTimestampLocal.Time.Equal(first.MeasurementTimestampLocal.Time) {
				t.Error("entry still present after RemoveBloodPressure")
			}
		}
	})
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
	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("d=%d_s=%d_p=%d", tc.d, tc.s, tc.p), func(t *testing.T) {
			bp := garmin.GarminBloodPressure{
				Diastolic: tc.d, Systolic: tc.s, Pulse: tc.p,
				MeasurementDateTime: time.Now(),
			}
			ok, err := c.AddBloodPressure(ctx, bp)
			if err != nil {
				t.Errorf("Add(%+v) error: %v", tc, err)
				return
			}
			if ok {
				t.Errorf("Add(%+v) returned true; expected validation failure", tc)
			}
		})
	}
}
