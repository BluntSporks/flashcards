package main

import (
	"testing"
)

func TestAvgScore(t *testing.T) {
	scores := []int{30, 60, 80, 100}
	exp := 67
	act := AvgScore(scores)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}

}
