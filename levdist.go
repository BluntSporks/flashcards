package card

import "strings"

// LevDist returns the levenshtein distance, customized to ignore case and spacing.
func LevDist(s string, t string) int {
	sBase := strings.Replace(strings.ToLower(s), " ", "", -1)
	tBase := strings.Replace(strings.ToLower(t), " ", "", -1)
	return LevDistAux(sBase, len(sBase), tBase, len(tBase))
}
