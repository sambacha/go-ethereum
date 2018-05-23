package aclock

import (
	"time"
)

var offset time.Duration

func init() {
	offset = 24 * time.Hour
}

func AdvanceTime(d time.Duration) time.Duration {
	offset += d
	return offset
}

func Reset() {
	offset = 0
}

func Now() time.Time {
	return time.Now().Add(offset)
}

func Since(t time.Time) time.Duration {
	return Now().Sub(t)
}
