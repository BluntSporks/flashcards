package main

import (
	"testing"
)

func TestMaxScore(t *testing.T) {
	choices := []string{"Hemingway", "Ernest Hemingway"}
	resp := "Hemingwhat"
	exp := 77
	act := maxScore(choices, resp)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}

}
