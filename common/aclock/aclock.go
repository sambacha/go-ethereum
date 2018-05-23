package aclock

import (
	"fmt"
	"time"
)

var offset time.Duration

func AdvanceTime(d time.Duration) (time.Duration, error) {
	if offset+d < offset {
		return 0, fmt.Errorf("aclock: offset overflow")
	}
	offset += d
	return offset, nil
}

func NowWithOffset() (time.Time, time.Duration) {
	return time.Now().Add(offset), offset
}

func Reset() {
	offset = 0
}

func Now() time.Time {
	return time.Now().Add(offset)
}
