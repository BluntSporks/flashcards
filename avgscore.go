package main

// avgScore returns the averge of a series of scores.
func avgScore(scores []int) int {
	if len(scores) == 0 {
		return 0
	}
	sum := 0
	for _, score := range scores {
		sum += score
	}
	return int(float64(sum) / float64(len(scores)))
}
