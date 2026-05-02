package garmin

import (
	"encoding/json"
	"testing"
	"time"
)

// TestLocalDateTimeParsesSupportedFormats mirrors C#'s DateTimeConverterTests.
// LocalDateTime must accept every shape Garmin emits in the wild.
func TestLocalDateTimeParsesSupportedFormats(t *testing.T) {
	cases := []struct {
		json string
		want time.Time
	}{
		{`"2026-04-26"`, time.Date(2026, 4, 26, 0, 0, 0, 0, time.UTC)},
		{`"2026-04-26 05:32:52"`, time.Date(2026, 4, 26, 5, 32, 52, 0, time.UTC)},
		{`"2026-04-26T05:32:52.4"`, time.Date(2026, 4, 26, 5, 32, 52, 400_000_000, time.UTC)},
		{`"2026-04-26T05:32:52.43"`, time.Date(2026, 4, 26, 5, 32, 52, 430_000_000, time.UTC)},
		{`"2026-04-26T05:32:52.431"`, time.Date(2026, 4, 26, 5, 32, 52, 431_000_000, time.UTC)},
	}
	for _, c := range cases {
		var got LocalDateTime
		if err := json.Unmarshal([]byte(c.json), &got); err != nil {
			t.Errorf("Unmarshal(%s) error: %v", c.json, err)
			continue
		}
		if !got.Time.Equal(c.want) {
			t.Errorf("Unmarshal(%s) = %v, want %v", c.json, got.Time, c.want)
		}
	}
}

// TestLocalDateTimeParsesTimezoneOffset is the C# Read_ParsesFormatWithTimezoneOffset
// case. The "+02:00" form must round-trip through the parser.
func TestLocalDateTimeParsesTimezoneOffset(t *testing.T) {
	const raw = `"2026-04-26T05:32:52.431+02:00"`
	var got LocalDateTime
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	want := time.Date(2026, 4, 26, 5, 32, 52, 431_000_000, time.FixedZone("+02:00", 2*60*60))
	if !got.Time.Equal(want) {
		t.Errorf("got %v, want %v", got.Time, want)
	}
}

// TestLocalDateTimeRejectsUnsupportedFormat is the C# Read_ThrowsForUnsupportedFormat case.
func TestLocalDateTimeRejectsUnsupportedFormat(t *testing.T) {
	const raw = `"26/04/2026"`
	var got LocalDateTime
	if err := json.Unmarshal([]byte(raw), &got); err == nil {
		t.Fatalf("expected error, got %v", got.Time)
	}
}

// TestDateRoundTrip checks the yyyy-MM-dd path used everywhere DateOnly was
// applied in the C# library.
func TestDateRoundTrip(t *testing.T) {
	d := NewDate(2026, time.April, 26)
	raw, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != `"2026-04-26"` {
		t.Fatalf("marshal got %s", raw)
	}
	var back Date
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatal(err)
	}
	if !back.Time.Equal(d.Time) {
		t.Errorf("round-trip mismatch: %v vs %v", back.Time, d.Time)
	}
}
