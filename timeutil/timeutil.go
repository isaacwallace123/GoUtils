package timeutil

import (
	"fmt"
	"time"
)

// NowUTC returns the current UTC time.
func NowUTC() time.Time {
	return time.Now().UTC()
}

// NowLocal returns the current local time.
func NowLocal() time.Time {
	return time.Now()
}

// FormatRFC3339 returns the time formatted using RFC3339.
func FormatRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

// FormatCustom formats the time using a custom layout.
func FormatCustom(t time.Time, layout string) string {
	return t.Format(layout)
}

// Since returns the duration since the given time.
func Since(t time.Time) time.Duration {
	return time.Since(t)
}

// Until returns the duration until the given time.
func Until(t time.Time) time.Duration {
	return time.Until(t)
}

// IsExpired returns true if the time has already passed.
func IsExpired(t time.Time) bool {
	return time.Now().After(t)
}

// Parse parses a time string using the given layout.
func Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

// Tick returns the current time since Unix epoch as float64 seconds.
func Tick() float64 {
	return float64(time.Now().UnixNano()) / 1e9
}

// ClockStart returns a timestamp to measure elapsed time from.
func ClockStart() time.Time {
	return time.Now()
}

// ClockSince returns elapsed time in seconds since a timestamp.
func ClockSince(start time.Time) float64 {
	return time.Since(start).Seconds()
}

// Elapsed is an alias for ClockSince.
func Elapsed(start time.Time) float64 {
	return ClockSince(start)
}

// SleepSeconds pauses execution for a given number of seconds.
func SleepSeconds(seconds float64) {
	time.Sleep(time.Duration(seconds * float64(time.Second)))
}

// Time conversion helpers

func SecondsToMinutes(seconds float64) float64 {
	return seconds / 60
}

func SecondsToHours(seconds float64) float64 {
	return seconds / 3600
}

func SecondsToDays(seconds float64) float64 {
	return seconds / 86400
}

func SecondsToWeeks(seconds float64) float64 {
	return seconds / (86400 * 7)
}

// FormatDuration breaks a duration in seconds into a human-readable string.
func FormatDuration(seconds float64) string {
	weeks := int(seconds) / 604800
	days := int(seconds) % 604800 / 86400
	hours := int(seconds) % 86400 / 3600
	minutes := int(seconds) % 3600 / 60
	secs := int(seconds) % 60

	result := ""
	if weeks > 0 {
		result += fmt.Sprintf("%dw ", weeks)
	}
	if days > 0 {
		result += fmt.Sprintf("%dd ", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%dh ", hours)
	}
	if minutes > 0 {
		result += fmt.Sprintf("%dm ", minutes)
	}
	if secs > 0 || result == "" {
		result += fmt.Sprintf("%ds", secs)
	}
	return result
}

// Roblox-style DateTime methods

func DateTimeNow() time.Time {
	return time.Now()
}

func DateTimeFromUnix(seconds int64) time.Time {
	return time.Unix(seconds, 0)
}

func ToUnixSeconds(t time.Time) int64 {
	return t.Unix()
}

func ToUnixMilliseconds(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func ToUnixMicroseconds(t time.Time) int64 {
	return t.UnixNano() / 1e3
}

func ToUnixNanoseconds(t time.Time) int64 {
	return t.UnixNano()
}

func DateTimeComponents(t time.Time) (year, month, day, hour, minute, second int) {
	y, m, d := t.Date()
	h, min, sec := t.Clock()
	return y, int(m), d, h, min, sec
}

func FormatClock(t time.Time) string {
	return t.Format("15:04:05")
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func DateTimeDiffSeconds(a, b time.Time) float64 {
	return a.Sub(b).Abs().Seconds()
}

// Clock tracks elapsed time since Start.
type Clock struct {
	start time.Time
}

// Start resets the clock to now.
func (c *Clock) Start() {
	c.start = time.Now()
}

// Elapsed returns seconds since the clock was started.
func (c *Clock) Elapsed() float64 {
	return time.Since(c.start).Seconds()
}

// Reset sets start to now and returns time elapsed since last start.
func (c *Clock) Reset() float64 {
	elapsed := time.Since(c.start).Seconds()
	c.start = time.Now()
	return elapsed
}

// CountdownTimer represents a timer that counts down from a set duration.
type CountdownTimer struct {
	duration time.Duration
	start    time.Time
}

// NewCountdownTimer creates a new countdown timer from seconds.
func NewCountdownTimer(seconds float64) *CountdownTimer {
	return &CountdownTimer{
		duration: time.Duration(seconds * float64(time.Second)),
		start:    time.Now(),
	}
}

// Remaining returns the remaining time in seconds.
func (ct *CountdownTimer) Remaining() float64 {
	remain := ct.duration - time.Since(ct.start)
	if remain < 0 {
		return 0
	}
	return remain.Seconds()
}

// Expired returns true if the countdown has finished.
func (ct *CountdownTimer) Expired() bool {
	return time.Since(ct.start) >= ct.duration
}

// Reset restarts the countdown from now.
func (ct *CountdownTimer) Reset() {
	ct.start = time.Now()
}
