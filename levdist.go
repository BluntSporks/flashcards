package card

// LevDist returns the levenshtein distance.
func LevDist(s string, t string) int {
	return LevDistAux(s, len(s), t, len(t))
}
