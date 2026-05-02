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
		t.Error(err)
		return
	}
	if profile.DisplayName == "" {
		t.Error("expected non-empty DisplayName")
	}
}

func TestGetPersonalRecord_NotNil(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	profile, err := c.GetSocialProfile(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	prs, err := c.GetPersonalRecord(ctx, profile.DisplayName)
	if err != nil {
		t.Error(err)
		return
	}
	if prs == nil {
		t.Error("expected non-nil PRs slice")
	}
}

func TestGetUserSettings_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	settings, err := c.GetUserSettings(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	if settings.ID == 0 {
		t.Error("expected non-zero ID")
	}
}

//nolint:staticcheck
func TestSetUserWeight(t *testing.T) {
	if !runDestructive() {
		t.Skip("set GARMIN_RUN_MUTATION_TESTS=1 to run state-mutating tests")
	}
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	original, err := c.GetUserSettings(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("set shifted weight", func(t *testing.T) {
		shifted := original.UserData.Weight + 1000
		if err := c.SetUserWeight(ctx, shifted); err != nil {
			t.Error(err)
			return
		}
		updated, err := c.GetUserSettings(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		if math.Abs(updated.UserData.Weight-shifted) > 1 {
			t.Errorf("got %v, want %v", updated.UserData.Weight, shifted)
		}
	})

	t.Run("restore original weight", func(t *testing.T) {
		if err := c.SetUserWeight(ctx, original.UserData.Weight); err != nil {
			t.Error(err)
		}
	})
}

func TestSetUserSleepTimes(t *testing.T) {
	if !runDestructive() {
		t.Skip("set GARMIN_RUN_MUTATION_TESTS=1 to run state-mutating tests")
	}
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	original, err := c.GetUserSettings(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	want := struct{ sleep, wake int64 }{1, 2}

	t.Run("set new sleep times", func(t *testing.T) {
		if err := c.SetUserSleepTimes(ctx, &want.sleep, &want.wake); err != nil {
			t.Error(err)
			return
		}
		updated, err := c.GetUserSettings(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		if updated.UserSleep.DefaultSleepTime || updated.UserSleep.SleepTime != want.sleep {
			t.Errorf("sleep: got %+v, want %d", updated.UserSleep, want.sleep)
		}
		if updated.UserSleep.DefaultWakeTime || updated.UserSleep.WakeTime != want.wake {
			t.Errorf("wake: got %+v, want %d", updated.UserSleep, want.wake)
		}
	})

	t.Run("restore original sleep times", func(t *testing.T) {
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
			t.Error(err)
		}
	})
}
