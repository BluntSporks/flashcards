package main

import (
	"bufio"
	"fmt"
	"os"
)

// showCard displays the question and grades the answer.
func showCard(card card) int {
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
	return maxScore(card.ansz, resp)
}
