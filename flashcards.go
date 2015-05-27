// Flashcards will quiz your from a deck of cards then give you a score.
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
	Shuffle(cards)

	for i, card := range cards {
		// Break at limit.
		if i == lim {
			break
		}

		// Display a card.
		score := show(card)

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
	avg := AvgScore(scores)
	fmt.Printf("Average score: %d\n", avg)
}

// loadCards loads the card mappings from a file.
func loadCardMap(file string) map[string]string {
	hdl, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer hdl.Close()
	reader := csv.NewReader(hdl)
	reader.FieldsPerRecord = 2
	cardMap := make(map[string]string)
	lineNum := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		lineNum++
		ques := record[0]
		if ques == "" {
			panic("Empty question on line " + string(lineNum))
		}
		ansList := record[1]
		if ansList == "" {
			panic("Empty answer list on line " + string(lineNum))
		}
		cardMap[ques] = ansList
	}
	return cardMap
}

// makeCards makes the cards into card structs.
func makeCards(cardMap map[string]string) []Card {
	cardCnt := len(cardMap)
	cards := make([]Card, cardCnt)
	i := 0
	for ques, ansList := range cardMap {
		ansz := strings.Split(ansList, "|")
		card := Card{ques, ansz}
		cards[i] = card
		i++
	}
	return cards
}

// show displays the question and grades the answer.
func show(card Card) int {
	fmt.Println(card.ques)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	resp := scanner.Text()
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	if resp == "" {
		return -1
	}
	return MaxScore(card.ansz, resp)
}
