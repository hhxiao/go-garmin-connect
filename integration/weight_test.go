//go:build integration
// +build integration

package integration

import (
	"context"
	"fmt"
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
		t.Error(err)
		return
	}
	if len(r.DailyWeightSummaries) == 0 {
		t.Skip("no weight history")
	}
}

func TestAdd_And_Remove_Weight_Success(t *testing.T) {
	if !runDestructive() {
		t.Skip("set GARMIN_RUN_MUTATION_TESTS=1 to run state-mutating tests")
	}
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	now := time.Now()
	startDate := now.AddDate(0, 0, -1)
	endDate := now.AddDate(0, 0, 1)
	weight := garmin.GarminWeight{MeasurementDateTime: now, UnitKey: garmin.WeightUnitKg, Value: 42}
	expectedGrams := weight.Value * 1000

	t.Run("add returns true", func(t *testing.T) {
		added, err := c.AddWeight(ctx, weight)
		if err != nil {
			t.Error(err)
			return
		}
		if !added {
			t.Error("AddWeight returned false for valid input")
		}
	})

	t.Run("entry visible after add", func(t *testing.T) {
		r, err := c.GetWeightRange(ctx, startDate, endDate)
		if err != nil {
			t.Error(err)
			return
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
			t.Error("could not find the weight just added")
			return
		}
		if math.Abs(found.Weight-expectedGrams) > 0.1 {
			t.Errorf("weight mismatch: got %v, want %v", found.Weight, expectedGrams)
		}
	})

	t.Run("remove then not found", func(t *testing.T) {
		r, err := c.GetWeightRange(ctx, startDate, endDate)
		if err != nil {
			t.Error(err)
			return
		}
		var found *garmin.GarminWeightMeasurement
		for i := range r.DailyWeightSummaries {
			s := r.DailyWeightSummaries[i]
			if s.SummaryDate.Time.Year() != now.Year() || s.SummaryDate.Time.YearDay() != now.YearDay() {
				continue
			}
			for j := range s.AllWeightMetrics {
				if math.Abs(s.AllWeightMetrics[j].Weight-expectedGrams) < 0.1 {
					found = &s.AllWeightMetrics[j]
					break
				}
			}
		}
		if found == nil {
			t.Skip("no matching entry to remove")
		}
		if err := c.RemoveWeight(ctx, found.GarminWeightIdentifier); err != nil {
			t.Error(err)
			return
		}
		r2, err := c.GetWeightRange(ctx, startDate, endDate)
		if err != nil {
			t.Error(err)
			return
		}
		for _, s := range r2.DailyWeightSummaries {
			if s.SummaryDate.Time.Year() != now.Year() || s.SummaryDate.Time.YearDay() != now.YearDay() {
				continue
			}
			for _, m := range s.AllWeightMetrics {
				if m.SamplePk == found.SamplePk && math.Abs(m.Weight-found.Weight) < 0.1 {
					t.Error("weight still present after RemoveWeight")
				}
			}
		}
	})
}

func TestAddWeight_FailsForOutOfRange(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, v := range []float64{0.0, 454.0} {
		v := v
		t.Run(fmt.Sprintf("value=%.1f", v), func(t *testing.T) {
			w := garmin.GarminWeight{MeasurementDateTime: time.Now(), UnitKey: garmin.WeightUnitKg, Value: v}
			ok, err := c.AddWeight(ctx, w)
			if err != nil {
				t.Errorf("AddWeight(%v) error: %v", v, err)
				return
			}
			if ok {
				t.Errorf("AddWeight(%v) returned true; expected validation failure", v)
			}
		})
	}
}
