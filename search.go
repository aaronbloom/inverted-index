package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func search(recordsIndex InverseIndex) {
	input := userInput()
	input = neutralString(input)

	if input == "exit" {
		return
	}

	fmt.Println("Lookup:", input)

	startTime := time.Now()

	for i := 0; i < len(recordsIndex.indexItems); i++ {
		termItem := recordsIndex.indexItems[i]
		if strings.Contains(termItem.term, input) {
			fmt.Println("Found term", termItem.id, termItem.term)
			for j := 0; j < len(termItem.records); j++ {
				record := recordsIndex.Record(termItem.records[j])
				fmt.Println("\tRecord associated:", record.ID(), record.Content())
			}
		}
	}

	timeTaken := time.Since(startTime)
	fmt.Printf("Lookup time taken: %s\n", timeTaken.String())

	search(recordsIndex) // loop
}

func userInput() string {
	fmt.Print("\nSearch term: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "exit"
	}

	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)

	return input
}
