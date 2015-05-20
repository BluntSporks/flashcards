package card

import (
	"testing"
)

func TestLevDist(t *testing.T) {
	origs := []string{" Ab C ", "abc", "abc", "abc"}
	mods := []string{"abc", "abcd", "ab", "abd"}
	exps := []int{0, 1, 1, 1}
	for i := 0; i < 4; i++ {
		act := LevDist(origs[i], mods[i])
		if exps[i] != act {
			t.Error("Expected", exps[i], "got", act)
		}
	}

}
