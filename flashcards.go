// Flashcards will quiz your from a deck of cards then give you a score.
package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// Parse arguments.
	limFlag := flag.Int("lim", 30, "Limit of maximum questions to include in the current session")
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Missing filename argument")
	}
	file := flag.Arg(0)

	// Load cards.
	cardMap := loadCardMap(file)
	cards := makeCards(cardMap)
	if len(cards) == 0 {
		log.Fatal("No cards loaded")
	}
	fmt.Printf("%d cards loaded\n", len(cards))

	// Set and display limit.
	var lim int
	if *limFlag >= len(cards) {
		lim = len(cards)
	} else {
		lim = *limFlag
		fmt.Printf("%d cards will be shown\n", lim)
	}
	fmt.Println()

	// Prepare score array.
	scores := make([]int, lim)

	// Shuffle cards.
	shuffleCards(cards)

	for i, card := range cards {
		// Break at limit.
		if i == lim {
			break
		}

		// Display a card.
		score := showCard(card)

		// Display answers if wrong.
		if score != 100 {
			for j, ans := range card.ansz {
				if j != 0 {
					fmt.Print(" or ")
				}
				fmt.Print(ans)
			}
			fmt.Println()
		}

		// Break if answer left empty so user can quit early.
		if score == -1 {
			break
		}

		// Display and save score.
		fmt.Printf("Score: %d\n\n", score)
		scores[i] = score
	}

	// Display average score.
	avg := avgScore(scores)
	fmt.Printf("Average score: %d\n", avg)
}
