package main

import (
	"math/rand"
	"time"
)

// shuffleCards puts the list of cards in random order.
func shuffleCards(array []card) {
	rand.Seed(time.Now().UnixNano())
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i)
		array[i], array[j] = array[j], array[i]
	}
}
