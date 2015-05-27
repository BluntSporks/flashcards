package main

import "strings"

// makeCards makes the cards into card structs.
func makeCards(cardMap map[string]string) []card {
	cardCnt := len(cardMap)
	cards := make([]card, cardCnt)
	i := 0
	for ques, ansList := range cardMap {
		ansz := strings.Split(ansList, "|")
		card := card{ques, ansz}
		cards[i] = card
		i++
	}
	return cards
}
