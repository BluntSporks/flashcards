// Flashcards will quiz your from a deck of cards then give you a score.
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("Missing filename argument")
	}
	file := os.Args[1]
	cardMap := loadCardMap(file)
	cards := makeCards(cardMap)
	scores := make([]int, len(cards))
	Shuffle(cards)
	fmt.Printf("%d cards loaded\n", len(cards))
	for i, card := range cards {
		score := show(card)
		if score != 100 {
			for j, ans := range card.ansz {
				if j != 0 {
					fmt.Print(" or ")
				}
				fmt.Print(ans)
			}
			fmt.Println()
		}
		if score == -1 {
			break
		}
		fmt.Printf("Score: %d\n\n", score)
		scores[i] = score
	}
	fmt.Printf("Average score: %d\n", AvgScore(scores))
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
