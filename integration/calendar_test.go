//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	garmin "github.com/sealbro/go-garmin-connect"
)

func TestGetCalendarYear_HasSummaries(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	year, err := c.GetCalendarByYear(ctx, time.Now().Year())
	if err != nil {
		t.Fatal(err)
	}
	if len(year.YearSummaries) == 0 {
		t.Fatal("expected at least one yearSummary")
	}
}

func TestGetCalendarMonth_MatchesQuery(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	prev := time.Now().AddDate(0, -1, 0)
	want := garmin.GarminMonth(int(prev.Month()) - 1)
	month, err := c.GetCalendarByMonth(ctx, prev.Year(), want)
	if err != nil {
		t.Fatal(err)
	}
	if month.Year != prev.Year() {
		t.Errorf("year: got %d, want %d", month.Year, prev.Year())
	}
	if month.Month != want {
		t.Errorf("month: got %d, want %d", month.Month, want)
	}
	if len(month.CalendarItems) == 0 {
		t.Skip("no calendar items in previous month")
	}
}

func TestGetCalendarWeek_BoundsContainQueryDay(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	day := time.Now().AddDate(0, 0, -7)
	week, err := c.GetCalendarByWeek(ctx, day)
	if err != nil {
		t.Fatal(err)
	}
	if day.Before(week.StartDate.Time) || day.After(week.EndDate.Time.AddDate(0, 0, 1)) {
		t.Fatalf("day %v not within %v–%v", day, week.StartDate.Time, week.EndDate.Time)
	}
}
