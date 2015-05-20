package card

// MaxScore returns the best score for a series of correct choices.
func MaxScore(choices []string, resp string) int {
	max := 0
	for _, choice := range choices {
		score := Score(choice, resp)
		if score > max {
			max = score
		}
	}
	return max
}
