package card

import (
	"testing"
)

func TestMaxScore(t *testing.T) {
	choices := []string{"Hemingway", "Ernest Hemingway"}
	resp := "Hemingwhat"
	exp := 77
	act := MaxScore(choices, resp)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}

}
