package main

// LevDistAux returns the levenshtein distance of two strings using the simple recursive version.
// See https://en.wikipedia.org/wiki/Levenshtein_distance.
func LevDistAux(s string, sLen int, t string, tLen int) int {
	// Base case is empty strings.
	if sLen == 0 {
		return tLen
	}
	if tLen == 0 {
		return sLen
	}

	// Test if last characters match.
	var cost int
	if s[sLen-1] == t[tLen-1] {
		cost = 0
	} else {
		cost = 1
	}

	// Return minimum of delete char from s, delete char from t, and delete char from both.
	dist1 := LevDistAux(s, sLen-1, t, tLen) + 1
	dist2 := LevDistAux(s, sLen, t, tLen-1) + 1
	dist3 := LevDistAux(s, sLen-1, t, tLen-1) + cost
	min := dist1
	if dist2 < min {
		min = dist2
	}
	if dist3 < min {
		min = dist3
	}
	return min
}
