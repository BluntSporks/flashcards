package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// loadCardMap loads the card mappings from a file.
func loadCardMap(file string) map[string]string {
	hdl, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}
		lineNum++
		ques := record[0]
		if ques == "" {
			log.Fatal("Empty question on line " + string(lineNum))
		}
		ansList := record[1]
		if ansList == "" {
			log.Fatal("Empty answer list on line " + string(lineNum))
		}
		cardMap[ques] = ansList
	}
	return cardMap
}
