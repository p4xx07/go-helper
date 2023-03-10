package time_helper

import (
	"testing"
	"time"
)

func TestFormatHHMMSSmm(t *testing.T) {
	expected := "01:01:01.01"
	duration, _ := time.ParseDuration("01h01m01s01ms")
	result := FormatHHMMSSmm(duration)
	if result != expected {
		panic("not equal")
	}
}
