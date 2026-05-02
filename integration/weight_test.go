//go:build integration
// +build integration

package integration

import (
	"context"
	"math"
	"testing"
	"time"

	garmin "github.com/sealbro/go-garmin-connect"
)

func TestGetWeightRange_NotEmptyOrSkip(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	start := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Now().AddDate(1, 0, 0)
	r, err := c.GetWeightRange(ctx, start, end)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.DailyWeightSummaries) == 0 {
		t.Skip("no weight history")
	}
}

func TestAdd_And_Remove_Weight_Success(t *testing.T) {
	if !runDestructive() {
		t.Skip("set GARMIN_RUN_DESTRUCTIVE=1 to run state-mutating tests")
	}
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	now := time.Now()
	startDate := now.AddDate(0, 0, -1)
	endDate := now.AddDate(0, 0, 1)
	weight := garmin.GarminWeight{MeasurementDateTime: now, UnitKey: garmin.WeightUnitKg, Value: 42}
	expectedGrams := weight.Value * 1000

	added, err := c.AddWeight(ctx, weight)
	if err != nil {
		t.Fatal(err)
	}
	if !added {
		t.Fatal("AddWeight returned false for valid input")
	}

	r, err := c.GetWeightRange(ctx, startDate, endDate)
	if err != nil {
		t.Fatal(err)
	}
	var found *garmin.GarminWeightMeasurement
	for i := range r.DailyWeightSummaries {
		s := r.DailyWeightSummaries[i]
		if s.SummaryDate.Time.Year() != now.Year() || s.SummaryDate.Time.YearDay() != now.YearDay() {
			continue
		}
		for j := range s.AllWeightMetrics {
			m := s.AllWeightMetrics[j]
			if math.Abs(m.Weight-expectedGrams) < 0.1 {
				found = &s.AllWeightMetrics[j]
				break
			}
		}
	}
	if found == nil {
		t.Fatal("could not find the weight just added")
	}
	if math.Abs(found.Weight-expectedGrams) > 0.1 {
		t.Errorf("weight mismatch: got %v, want %v", found.Weight, expectedGrams)
	}

	if err := c.RemoveWeight(ctx, found.GarminWeightIdentifier); err != nil {
		t.Fatal(err)
	}

	r, err = c.GetWeightRange(ctx, startDate, endDate)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range r.DailyWeightSummaries {
		if s.SummaryDate.Time.Year() != now.Year() || s.SummaryDate.Time.YearDay() != now.YearDay() {
			continue
		}
		for _, m := range s.AllWeightMetrics {
			if m.SamplePk == found.SamplePk && math.Abs(m.Weight-found.Weight) < 0.1 {
				t.Fatal("weight still present after RemoveWeight")
			}
		}
	}
}

func TestAddWeight_FailsForOutOfRange(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, v := range []float64{0.0, 454.0} {
		w := garmin.GarminWeight{MeasurementDateTime: time.Now(), UnitKey: garmin.WeightUnitKg, Value: v}
		ok, err := c.AddWeight(ctx, w)
		if err != nil {
			t.Fatalf("AddWeight(%v) error: %v", v, err)
		}
		if ok {
			t.Errorf("AddWeight(%v) returned true; expected validation failure", v)
		}
	}
}
