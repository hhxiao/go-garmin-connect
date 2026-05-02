//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"
	"time"
)

func wellnessWindow() (start, end time.Time) {
	end = time.Now()
	start = end.AddDate(0, 0, -1)
	return
}

func TestGetUserSummary_NotZero(t *testing.T) {
	c := lazyClient(t)
	start, _ := wellnessWindow()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	s, err := c.GetUserSummary(ctx, start)
	if err != nil {
		t.Error(err)
		return
	}
	if s.UserProfileID == 0 {
		t.Error("expected non-zero UserProfileID")
	}
}

func TestGetWellnessStepsData_NotNil(t *testing.T) {
	c := lazyClient(t)
	start, _ := wellnessWindow()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	steps, err := c.GetWellnessStepsData(ctx, start)
	if err != nil {
		t.Error(err)
		return
	}
	if steps == nil {
		t.Error("expected non-nil slice")
	}
}

func TestGetWellnessSleepData_HasDate(t *testing.T) {
	c := lazyClient(t)
	start, _ := wellnessWindow()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	sleep, err := c.GetWellnessSleepData(ctx, start)
	if err != nil {
		t.Error(err)
		return
	}
	if sleep.DailySleepDto.UserProfilePk == 0 {
		t.Skip("no sleep data for yesterday")
	}
}

func TestGetWellnessHeartRates_HasDate(t *testing.T) {
	c := lazyClient(t)
	start, _ := wellnessWindow()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	hr, err := c.GetWellnessHeartRates(ctx, start)
	if err != nil {
		t.Error(err)
		return
	}
	if hr.UserProfilePk == 0 {
		t.Skip("no HR data for yesterday")
	}
}

func TestGetWellnessBodyBattery_NotNil(t *testing.T) {
	c := lazyClient(t)
	start, end := wellnessWindow()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	bb, err := c.GetWellnessBodyBatteryData(ctx, start, end)
	if err != nil {
		t.Error(err)
		return
	}
	if bb == nil {
		t.Error("expected non-nil slice")
	}
}

func TestGetHydrationData_NotZero(t *testing.T) {
	c := lazyClient(t)
	start, _ := wellnessWindow()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	h, err := c.GetHydrationData(ctx, start)
	if err != nil {
		t.Error(err)
		return
	}
	if h.UserID == 0 {
		t.Skip("no hydration data")
	}
}

func TestGetBodyComposition_HasDates(t *testing.T) {
	c := lazyClient(t)
	start, end := wellnessWindow()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	bc, err := c.GetBodyComposition(ctx, start, end)
	if err != nil {
		t.Error(err)
		return
	}
	if bc.StartDate.IsZero() && bc.EndDate.IsZero() {
		t.Skip("no body composition data")
	}
}
