// Shuffle the list of cards in place.
package main

import (
	"math/rand"
	"time"
)

func Shuffle(array []Card) {
	rand.Seed(time.Now().UnixNano())
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i)
		array[i], array[j] = array[j], array[i]
	}
}
