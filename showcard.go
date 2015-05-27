package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// showCard displays the question and grades the answer.
func showCard(card card) int {
	fmt.Println(card.ques)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	resp := scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if resp == "" {
		return -1
	}
	return maxScore(card.ansz, resp)
}
