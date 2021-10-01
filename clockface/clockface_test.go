package maths

import (
	"testing"
	"time"
)

func TestSecondHandAtMidNight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}
