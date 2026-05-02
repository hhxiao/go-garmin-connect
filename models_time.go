package garmin

import (
	"fmt"
	"strings"
	"time"
)

// Date is a calendar date without a time-of-day. Garmin sends these as
// "yyyy-MM-dd" strings; we translate to/from time.Time normalized to UTC
// midnight so callers can use it like any other time value.
type Date struct{ time.Time }

const dateLayout = "2006-01-02"

// UnmarshalJSON parses an `"yyyy-MM-dd"` string. JSON null and empty strings
// produce a zero Date.
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		d.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return fmt.Errorf("parse date %q: %w", s, err)
	}
	d.Time = t
	return nil
}

// MarshalJSON emits `"yyyy-MM-dd"`.
func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + d.Time.Format(dateLayout) + `"`), nil
}

// String returns the canonical "yyyy-MM-dd" representation.
func (d Date) String() string { return d.Time.Format(dateLayout) }

// NewDate builds a Date from year/month/day at UTC midnight.
func NewDate(year int, month time.Month, day int) Date {
	return Date{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

// LocalDateTime is a permissive datetime that accepts the half-dozen formats
// Garmin sprinkles across its API surface (with/without subseconds, with/
// without timezone offset, and bare dates). It mirrors
// `Garmin.Connect.Converters.DateTimeConverter`.
type LocalDateTime struct{ time.Time }

var datetimeFormats = []string{
	"2006-01-02T15:04:05.000-07:00", // Format5: with offset
	"2006-01-02T15:04:05.000",       // Format4: ms
	"2006-01-02T15:04:05.00",        // Format3: cs
	"2006-01-02T15:04:05.0",         // Format2: ds
	"2006-01-02 15:04:05",           // Format1: space-separated
	"2006-01-02",                    // Format0: bare date
}

// UnmarshalJSON parses any of the C# DateTimeConverter formats.
func (d *LocalDateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		d.Time = time.Time{}
		return nil
	}
	for _, layout := range datetimeFormats {
		if t, err := time.Parse(layout, s); err == nil {
			d.Time = t
			return nil
		}
	}
	// Fall back to RFC3339 — some endpoints use it.
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		d.Time = t
		return nil
	}
	return fmt.Errorf("unrecognized datetime %q", s)
}

// MarshalJSON serializes as UTC `yyyy-MM-ddTHH:mm:ss.fff`, matching the C#
// reference implementation.
func (d LocalDateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + d.Time.UTC().Format("2006-01-02T15:04:05.000") + `"`), nil
}
