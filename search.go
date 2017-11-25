package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
)

func search(records []recordItem, termIndex []indexItem) {
	input := userInput()
	input = neutralString(input)

	if (input == "exit") {
		return
	}

	fmt.Println("Lookup:", input)

	startTime := time.Now()
	
	for i := 0; i < len(termIndex); i++ {
		termItem := termIndex[i]
		if(strings.Contains(termItem.term, input)) {
			fmt.Println("Found term", termItem.id, termItem.term)
			for j := 0; j < len(termItem.records); j++ {
				record := records[termItem.records[j]]
				fmt.Println("\tRecord associated:", record.id, record.content)
			}
		}
	}

	timeTaken := time.Since(startTime)
	fmt.Printf("Lookup time taken: %s\n", timeTaken.String())

	search(records, termIndex) // loop
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
