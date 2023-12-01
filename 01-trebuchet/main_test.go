package main

import "testing"

func TestCalibrationSum(t *testing.T) {
	got := CalibrationSum()
	want := 142
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
