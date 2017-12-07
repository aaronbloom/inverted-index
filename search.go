package main

import (
	"bufio"
	"fmt"
	"github.com/aaronbloom/inverted-index/recordkeeper"
	"github.com/aaronbloom/inverted-index/utils"
	"os"
	"strings"
	"time"
)

func search(recordsIndex recordkeeper.InvertedIndex) {
	input := userInput()

	if input == "exit" {
		return
	}

	startTime := time.Now()

	var results = recordsIndex.Search(input, utils.LowerCaseFilter)

	timeTaken := time.Since(startTime)

	for _, record := range results {
		fmt.Printf("\t%d %s\n", record.ID(), record.Contents())
	}

	resultsCount := len(results)
	if resultsCount > 0 {
		fmt.Printf("\nCount: %d\n", resultsCount)
	}

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
