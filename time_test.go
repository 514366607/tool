package tools

import (
	"testing"
	"time"
)

func TestWeekCount(t *testing.T) {
	var w = WeekCount(time.Date(2019, 7, 3, 0, 0, 0, 0, new(time.Location)), 0)
	if w != 2583 {
		t.Fatalf("这天是周三，应该是2583,非%d", w)
	}

	w = WeekCount(time.Date(2019, 6, 30, 0, 0, 0, 0, new(time.Location)), 0)
	if w != 2582 {
		t.Fatalf("这天是周日，应该是2582,非%d", w)
	}
}
