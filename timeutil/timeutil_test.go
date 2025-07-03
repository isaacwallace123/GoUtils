package timeutil

import (
	"testing"
	"time"
)

// TestNowUTCAndLocal ensures UTC and Local return time.Time values with different locations.
func TestNowUTCAndLocal(t *testing.T) {
	utc := NowUTC()
	local := NowLocal()
	if utc.Location() == local.Location() {
		t.Error("UTC and local should not have the same Location")
	}
}

// TestFormatAndParse checks correct round-trip format and parsing of RFC3339.
func TestFormatAndParse(t *testing.T) {
	now := time.Now()
	rfc := FormatRFC3339(now)
	parsed, err := Parse(time.RFC3339, rfc)
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}
	// Allow up to 1 second difference due to truncation, otherwise check equality
	if !parsed.Equal(now.UTC().Truncate(time.Second)) {
		if parsed.Sub(now.UTC().Truncate(time.Second)) > time.Second {
			t.Errorf("Format/Parse mismatch: %v vs %v", parsed, now)
		}
	}
}

// TestClock checks Clock struct's elapsed/Reset functionality.
func TestClock(t *testing.T) {
	c := &Clock{}
	c.Start()
	time.Sleep(20 * time.Millisecond)
	elapsed := c.Elapsed()
	if elapsed <= 0 {
		t.Error("Clock.Elapsed should be > 0")
	}
	resetElapsed := c.Reset()
	if resetElapsed < elapsed {
		t.Error("Clock.Reset should be >= Elapsed")
	}
}

// TestCountdownTimer ensures timer expires after the duration and resets correctly.
func TestCountdownTimer(t *testing.T) {
	ct := NewCountdownTimer(0.05)
	if ct.Expired() {
		t.Error("Timer should not be expired immediately")
	}
	time.Sleep(60 * time.Millisecond)
	if !ct.Expired() {
		t.Error("Timer should be expired after wait")
	}
	ct.Reset()
	if ct.Remaining() <= 0 {
		t.Error("Remaining should be > 0 after reset")
	}
}

// TestTimeHelpers checks float conversions of seconds to minutes/hours/days.
func TestTimeHelpers(t *testing.T) {
	secs := 3661.0 // 1h 1m 1s

	// Test float minutes conversion
	gotMin := SecondsToMinutes(secs)
	wantMin := 61.016666666666666
	if (gotMin < wantMin-1e-6) || (gotMin > wantMin+1e-6) {
		t.Errorf("SecondsToMinutes failed: got %v, want %v", gotMin, wantMin)
	}

	// Test float hours conversion
	gotHr := SecondsToHours(secs)
	wantHr := 1.0169444444444444
	if (gotHr < wantHr-1e-6) || (gotHr > wantHr+1e-6) {
		t.Errorf("SecondsToHours failed: got %v, want %v", gotHr, wantHr)
	}

	// Test float days conversion
	gotDay := SecondsToDays(secs)
	wantDay := 0.042372685185185186
	if (gotDay < wantDay-1e-6) || (gotDay > wantDay+1e-6) {
		t.Errorf("SecondsToDays failed: got %v, want %v", gotDay, wantDay)
	}
}

// TestFormatDuration checks the human-readable formatting of durations.
func TestFormatDuration(t *testing.T) {
	d := FormatDuration(90061) // Should be "1d 1h 1m 1s"
	if d != "1d 1h 1m 1s" {
		t.Errorf("FormatDuration unexpected: %q", d)
	}
}
