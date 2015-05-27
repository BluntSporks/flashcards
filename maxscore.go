package main

// maxScore returns the best score for a series of correct choices.
func maxScore(choices []string, resp string) int {
	max := 0
	for _, choice := range choices {
		score := score(choice, resp)
		if score > max {
			max = score
		}
	}
	return max
}
