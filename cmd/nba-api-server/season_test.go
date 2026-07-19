package main

import (
	"net/http/httptest"
	"testing"
	"time"
)

// withNowFunc pins nowFunc to a fixed instant for the duration of a test
// and restores the real clock afterward via t.Cleanup, so season-default
// tests don't depend on (or drift with) the actual wall-clock date.
func withNowFunc(t *testing.T, fixed time.Time) {
	t.Helper()
	original := nowFunc
	nowFunc = func() time.Time { return fixed }
	t.Cleanup(func() { nowFunc = original })
}

func TestCurrentSeasonDefault(t *testing.T) {
	tests := []struct {
		name string
		now  time.Time
		want string
	}{
		{name: "mid-season, January", now: time.Date(2026, time.January, 15, 0, 0, 0, 0, time.UTC), want: "2025-26"},
		{name: "off-season, July", now: time.Date(2026, time.July, 19, 0, 0, 0, 0, time.UTC), want: "2025-26"},
		{name: "September 30, still previous season", now: time.Date(2026, time.September, 30, 23, 59, 59, 0, time.UTC), want: "2025-26"},
		{name: "October 1, new season starts", now: time.Date(2026, time.October, 1, 0, 0, 0, 0, time.UTC), want: "2026-27"},
		{name: "December, new season underway", now: time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC), want: "2026-27"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withNowFunc(t, tt.now)
			if got := currentSeasonDefault(); got != tt.want {
				t.Errorf("currentSeasonDefault() at %s = %q, want %q", tt.now.Format("2006-01-02"), got, tt.want)
			}
		})
	}
}

func TestGetSeasonOrDefault(t *testing.T) {
	withNowFunc(t, time.Date(2026, time.July, 19, 0, 0, 0, 0, time.UTC))

	t.Run("query parameter present overrides the default", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/?Season=2019-20", nil)
		if got := getSeasonOrDefault(r); got != "2019-20" {
			t.Errorf("getSeasonOrDefault() = %q, want %q", got, "2019-20")
		}
	})

	t.Run("query parameter absent falls back to currentSeasonDefault", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		want := currentSeasonDefault()
		if got := getSeasonOrDefault(r); got != want {
			t.Errorf("getSeasonOrDefault() = %q, want %q", got, want)
		}
	})
}
