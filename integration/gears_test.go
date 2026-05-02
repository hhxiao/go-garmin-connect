//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"
	"time"
)

func TestGetGearTypes_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	types, err := c.GetGearTypes(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	if len(types) == 0 {
		t.Error("expected at least one gear type")
	}
}

func TestGetUserGears_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	acts, err := c.GetActivities(ctx, 1, 1)
	if err != nil {
		t.Error(err)
		return
	}
	if len(acts) == 0 {
		t.Skip("no activities")
	}

	gears, err := c.GetUserGears(ctx, acts[0].OwnerID)
	if err != nil {
		t.Error(err)
		return
	}
	if len(gears) == 0 {
		t.Skip("no gear registered")
	}
}

func TestGetActivityGears_NotNil(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	acts, err := c.GetActivities(ctx, 1, 1)
	if err != nil {
		t.Error(err)
		return
	}
	if len(acts) == 0 {
		t.Skip("no activities")
	}

	gears, err := c.GetActivityGears(ctx, acts[0].ActivityID)
	if err != nil {
		t.Error(err)
		return
	}
	if gears == nil {
		t.Error("expected non-nil slice")
	}
}
