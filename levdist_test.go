package main

import "testing"

func TestLevDist(t *testing.T) {
	orig := "abc"
	mods := []string{"Abc", "abcd", "ab", "abd"}
	exps := []int{1, 1, 1, 1}
	for i := 0; i < 4; i++ {
		act := LevDist(orig, mods[i])
		if exps[i] != act {
			t.Error("Expected", exps[i], "got", act)
		}
	}

}
