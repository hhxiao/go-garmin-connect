//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"
	"time"
)

func TestGetReportHrvStatus_NotEmptyOrSkip(t *testing.T) {
	c := lazyClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	end := time.Now()
	start := end.AddDate(0, 0, -3)
	report, err := c.GetReportHrvStatus(ctx, start, end)
	if err != nil {
		t.Fatal(err)
	}
	// Mirrors C# behaviour: HRV may genuinely be absent for some accounts;
	// an empty list is allowed but the call must succeed.
	if len(report.HrvSummaries) == 0 {
		t.Skip("no HRV summaries in window")
	}
}
