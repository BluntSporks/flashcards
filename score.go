package main

import "github.com/BluntSporks/string-functions"

// score returns the score of a user response on a scale from 0 to 100.
func score(correct string, resp string) int {
	dist := strfunc.LevDist(correct, resp)
	prop := 1.0 - float64(dist)/float64(len(correct))
	return int(prop * 100.0)
}
