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
	workoutsOnce sync.Once
	cachedWOs    []garmin.GarminWorkout
	workoutsErr  error
)

func smallestWorkoutPage(t *testing.T, c *garmin.Client) ([]garmin.GarminWorkout, bool) {
	t.Helper()
	workoutsOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		params := garmin.DefaultWorkoutsParameters()
		params.OrderSeq = garmin.OrderSeqAsc
		params.OrderBy = garmin.WorkoutsOrderByCreatedDate
		params.Limit = 5
		cachedWOs, workoutsErr = c.GetWorkouts(ctx, params)
	})
	if workoutsErr != nil {
		t.Errorf("GetWorkouts: %v", workoutsErr)
		return nil, false
	}
	if len(cachedWOs) == 0 {
		t.Skip("no workouts on account")
		return nil, false
	}
	return cachedWOs, true
}

func TestGetWorkouts_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	_, ok := smallestWorkoutPage(t, c)
	if !ok {
		t.Error("expected at least one workout")
	}
}

func TestGetWorkoutTypes_NotZero(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	types, err := c.GetWorkoutTypes(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	if len(types.WorkoutSportTypes) == 0 {
		t.Error("expected non-empty WorkoutSportTypes")
	}
}

func TestGetWorkout_NotZero(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	params := garmin.DefaultWorkoutsParameters()
	params.OrderBy = garmin.WorkoutsOrderByUpdateDate
	page, err := c.GetWorkouts(ctx, params)
	if err != nil {
		t.Error(err)
		return
	}
	if len(page) == 0 {
		t.Skip("no workouts on account")
	}
	w, err := c.GetWorkout(ctx, page[0].WorkoutID)
	if err != nil {
		t.Error(err)
		return
	}
	if w.WorkoutID == 0 {
		t.Error("expected non-zero WorkoutID")
	}
}

func TestUpdateWorkout(t *testing.T) {
	if !runDestructive() {
		t.Skip("set GARMIN_RUN_MUTATION_TESTS=1 to run state-mutating tests")
	}
	c := lazyClient(t)
	page, ok := smallestWorkoutPage(t, c)
	if !ok {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	const want = 2000.0
	workout, err := c.GetWorkout(ctx, page[0].WorkoutID)
	if err != nil {
		t.Error(err)
		return
	}

	originalStep := workout.WorkoutSegments[0].WorkoutSteps[0]
	originalValue := originalStep.EndConditionValue

	t.Run("set new EndConditionValue", func(t *testing.T) {
		v := want
		workout.WorkoutSegments[0].WorkoutSteps[0].EndConditionValue = &v
		if err := c.UpdateWorkout(ctx, workout); err != nil {
			t.Error(err)
			return
		}
		got, err := c.GetWorkout(ctx, workout.WorkoutID)
		if err != nil {
			t.Error(err)
			return
		}
		if v := got.WorkoutSegments[0].WorkoutSteps[0].EndConditionValue; v == nil || *v != want {
			t.Errorf("EndConditionValue: got %v, want %v", v, want)
		}
	})

	t.Run("restore original EndConditionValue", func(t *testing.T) {
		got, err := c.GetWorkout(ctx, workout.WorkoutID)
		if err != nil {
			t.Error(err)
			return
		}
		got.WorkoutSegments[0].WorkoutSteps[0].EndConditionValue = originalValue
		if err := c.UpdateWorkout(ctx, got); err != nil {
			t.Error(err)
		}
	})
}

func TestScheduleWorkout(t *testing.T) {
	if !runDestructive() {
		t.Skip("set GARMIN_RUN_MUTATION_TESTS=1 to run state-mutating tests")
	}
	c := lazyClient(t)
	page, ok := smallestWorkoutPage(t, c)
	if !ok {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	w := page[0]
	scheduleDate := w.CreatedDate.Time

	week, err := c.GetCalendarByWeek(ctx, scheduleDate)
	if err != nil {
		t.Error(err)
		return
	}
	for _, item := range week.CalendarItems {
		if item.WorkoutID != nil && *item.WorkoutID == w.WorkoutID {
			t.Skipf("workout %d already scheduled in week", w.WorkoutID)
		}
	}

	t.Run("schedule workout", func(t *testing.T) {
		if err := c.ScheduleWorkout(ctx, w.WorkoutID, scheduleDate); err != nil {
			t.Error(err)
			return
		}
		week, err := c.GetCalendarByWeek(ctx, scheduleDate)
		if err != nil {
			t.Error(err)
			return
		}
		var calendarID int64
		for _, item := range week.CalendarItems {
			if item.WorkoutID != nil && *item.WorkoutID == w.WorkoutID {
				calendarID = item.ID
				break
			}
		}
		if calendarID == 0 {
			t.Error("scheduled workout did not appear in calendar")
		}
	})

	t.Run("remove scheduled workout", func(t *testing.T) {
		week, err := c.GetCalendarByWeek(ctx, scheduleDate)
		if err != nil {
			t.Error(err)
			return
		}
		var calendarID int64
		for _, item := range week.CalendarItems {
			if item.WorkoutID != nil && *item.WorkoutID == w.WorkoutID {
				calendarID = item.ID
				break
			}
		}
		if calendarID == 0 {
			t.Skip("no scheduled workout found to remove")
		}
		if err := c.RemoveScheduledWorkout(ctx, calendarID); err != nil {
			t.Error(err)
			return
		}
		week, err = c.GetCalendarByWeek(ctx, scheduleDate)
		if err != nil {
			t.Error(err)
			return
		}
		for _, item := range week.CalendarItems {
			if item.WorkoutID != nil && *item.WorkoutID == w.WorkoutID {
				t.Error("workout still present after RemoveScheduledWorkout")
			}
		}
	})
}
