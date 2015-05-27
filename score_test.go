package main

import (
	"testing"
)

func TestScore(t *testing.T) {
	correct := "abcde"
	resps := []string{"abcde", "abc", ""}
	exps := []int{100, 60, 0}
	for i := 0; i < 3; i++ {
		act := score(correct, resps[i])
		if exps[i] != act {
			t.Error("Expected", exps[i], "got", act)
		}
	}

}
