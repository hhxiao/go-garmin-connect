//go:build integration
// +build integration

package integration

import (
	"context"
	"sync"
	"testing"
	"time"

	garmin "github.com/sealbro/go-garmin-connect"
)

var (
	activitiesOnce sync.Once
	cachedActs     []garmin.GarminActivity
	activitiesErr  error
)

// firstActivity returns activity #2 (zero-indexed) — same offset the C# tests
// use to avoid hitting whichever activity is currently being uploaded.
func firstActivity(t *testing.T, c *garmin.Client) garmin.GarminActivity {
	t.Helper()
	activitiesOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		cachedActs, activitiesErr = c.GetActivities(ctx, 2, 1)
	})
	if activitiesErr != nil {
		t.Fatalf("GetActivities: %v", activitiesErr)
	}
	if len(cachedActs) == 0 {
		t.Skip("no activities available for this account")
	}
	return cachedActs[0]
}

func TestGetActivities_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	_ = firstActivity(t, c)
}

func TestGetActivitiesByDate_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	now := time.Now()
	acts, err := c.GetActivitiesByDate(ctx, now.AddDate(0, 0, -30), now.AddDate(0, 0, -2), "walking")
	if err != nil {
		t.Fatal(err)
	}
	if len(acts) == 0 {
		t.Skip("no walking activities in the last 30 days")
	}
}

func TestDownloadActivity_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	act := firstActivity(t, c)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	data, err := c.DownloadActivity(ctx, act.ActivityID, garmin.FormatTCX)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty download")
	}
}

func TestGetActivityExerciseSets_HasID(t *testing.T) {
	c := lazyClient(t)
	act := firstActivity(t, c)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	sets, err := c.GetActivityExerciseSets(ctx, act.ActivityID)
	if err != nil {
		t.Fatal(err)
	}
	if sets.ActivityID == 0 {
		t.Fatal("expected non-zero ActivityID")
	}
}

func TestGetActivityHrInTimezones_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	act := firstActivity(t, c)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	zones, err := c.GetActivityHrInTimezones(ctx, act.ActivityID)
	if err != nil {
		t.Fatal(err)
	}
	if len(zones) == 0 {
		t.Skip("no HR zone data for this activity")
	}
}

func TestGetActivitySplits_HasID(t *testing.T) {
	c := lazyClient(t)
	act := firstActivity(t, c)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	splits, err := c.GetActivitySplits(ctx, act.ActivityID)
	if err != nil {
		t.Fatal(err)
	}
	if splits.ActivityID == 0 {
		t.Fatal("expected non-zero ActivityID")
	}
}

func TestGetActivityWeather_HasIssueDate(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	now := time.Now()
	acts, err := c.GetActivitiesByDate(ctx, now.AddDate(0, 0, -30), now.AddDate(0, 0, -2), "walking")
	if err != nil {
		t.Fatal(err)
	}
	if len(acts) == 0 {
		t.Skip("no walking activities in window")
	}
	w, err := c.GetActivityWeather(ctx, acts[0].ActivityID)
	if err != nil {
		t.Fatal(err)
	}
	if w.IssueDate.Time.IsZero() {
		t.Fatal("expected IssueDate to be populated")
	}
}

func TestGetActivityDetails_HasMetrics(t *testing.T) {
	c := lazyClient(t)
	act := firstActivity(t, c)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	details, err := c.GetActivityDetails(ctx, act.ActivityID, 50, 50)
	if err != nil {
		t.Fatal(err)
	}
	if len(details.ActivityDetailMetrics) == 0 {
		t.Fatal("expected non-empty metrics")
	}
}

func TestGetActivitySplitSummaries_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	act := firstActivity(t, c)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s, err := c.GetActivitySplitSummaries(ctx, act.ActivityID)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.SplitSummaries) == 0 {
		t.Skip("no split summaries on this activity")
	}
}
