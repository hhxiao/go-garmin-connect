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
	devicesOnce sync.Once
	cachedDevs  []garmin.GarminDevice
	devicesErr  error
)

func devices(t *testing.T, c *garmin.Client) []garmin.GarminDevice {
	t.Helper()
	devicesOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		cachedDevs, devicesErr = c.GetDevices(ctx)
	})
	if devicesErr != nil {
		t.Fatal(devicesErr)
	}
	if len(cachedDevs) == 0 {
		t.Skip("no devices registered")
	}
	return cachedDevs
}

func TestGetDevices_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	_ = devices(t, c)
}

func TestGetDeviceSettings_NotZero(t *testing.T) {
	c := lazyClient(t)
	devs := devices(t, c)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s, err := c.GetDeviceSettings(ctx, devs[0].DeviceID)
	if err != nil {
		t.Fatal(err)
	}
	if s.DeviceID == 0 {
		t.Fatal("expected non-zero DeviceID")
	}
}

func TestGetDeviceLastUsed_NotZero(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	last, err := c.GetDeviceLastUsed(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if last.UserDeviceID == 0 {
		t.Skip("no last-used device")
	}
}

func TestGetDeviceMessages_NotEmpty(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	msgs, err := c.GetDeviceMessages(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(msgs.Messages) == 0 {
		t.Skip("no device messages")
	}
}
