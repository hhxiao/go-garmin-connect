//go:build integration
// +build integration

package integration

import (
	"context"
	"math"
	"testing"
	"time"
)

func TestGetSocialProfile_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	profile, err := c.GetSocialProfile(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if profile.DisplayName == "" {
		t.Fatal("expected non-empty DisplayName")
	}
}

func TestGetPersonalRecord_NotNil(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	profile, err := c.GetSocialProfile(ctx)
	if err != nil {
		t.Fatal(err)
	}
	prs, err := c.GetPersonalRecord(ctx, profile.DisplayName)
	if err != nil {
		t.Fatal(err)
	}
	if prs == nil {
		t.Fatal("expected non-nil PRs slice")
	}
}

func TestGetUserSettings_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	settings, err := c.GetUserSettings(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if settings.ID == 0 {
		t.Fatal("expected non-zero ID")
	}
}

// TestSetUserWeight bumps the configured weight by 1kg, reads it back, then
// restores the original value. Skipped unless GARMIN_RUN_DESTRUCTIVE=1 since
// it mutates account state.
//
//nolint:staticcheck // Calls the deprecated SetUserWeight on purpose to keep parity with C# tests.
func TestSetUserWeight(t *testing.T) {
	if !runDestructive() {
		t.Skip("set GARMIN_RUN_DESTRUCTIVE=1 to run state-mutating tests")
	}
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	original, err := c.GetUserSettings(ctx)
	if err != nil {
		t.Fatal(err)
	}
	shifted := original.UserData.Weight + 1000
	if err := c.SetUserWeight(ctx, shifted); err != nil {
		t.Fatal(err)
	}
	updated, err := c.GetUserSettings(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(updated.UserData.Weight-shifted) > 1 {
		t.Errorf("got %v, want %v", updated.UserData.Weight, shifted)
	}
	// Restore.
	if err := c.SetUserWeight(ctx, original.UserData.Weight); err != nil {
		t.Fatal(err)
	}
}

func TestSetUserSleepTimes(t *testing.T) {
	if !runDestructive() {
		t.Skip("set GARMIN_RUN_DESTRUCTIVE=1 to run state-mutating tests")
	}
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	original, err := c.GetUserSettings(ctx)
	if err != nil {
		t.Fatal(err)
	}

	want := struct{ sleep, wake int64 }{1, 2}
	if err := c.SetUserSleepTimes(ctx, &want.sleep, &want.wake); err != nil {
		t.Fatal(err)
	}
	updated, err := c.GetUserSettings(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if updated.UserSleep.DefaultSleepTime || updated.UserSleep.SleepTime != want.sleep {
		t.Errorf("sleep: got %+v, want %d", updated.UserSleep, want.sleep)
	}
	if updated.UserSleep.DefaultWakeTime || updated.UserSleep.WakeTime != want.wake {
		t.Errorf("wake: got %+v, want %d", updated.UserSleep, want.wake)
	}

	// Restore — pass nil to flag DefaultSleepTime / DefaultWakeTime when the
	// original record was at the device default.
	var s, w *int64
	if !original.UserSleep.DefaultSleepTime {
		v := original.UserSleep.SleepTime
		s = &v
	}
	if !original.UserSleep.DefaultWakeTime {
		v := original.UserSleep.WakeTime
		w = &v
	}
	if err := c.SetUserSleepTimes(ctx, s, w); err != nil {
		t.Fatal(err)
	}
}
