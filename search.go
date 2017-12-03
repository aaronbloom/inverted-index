package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aaronbloom/inverted-index/recordkeeper"
	"github.com/aaronbloom/inverted-index/utils"
)

func search(recordsIndex *recordkeeper.RecordIndex) (stop bool, err error) {
	input, err := userInput()
	if err != nil {
		return true, fmt.Errorf("error prompting for user input: %v", err)
	}

	if input == "exit" {
		return true, nil
	}

	startTime := time.Now()

	var results = (*recordsIndex).Search(input, utils.LowerCaseFilter)

	timeTaken := time.Since(startTime)

	for _, record := range results {
		fmt.Printf("\t%d %s\n", record.ID(), record.Contents())
	}

	resultsCount := len(results)
	if resultsCount > 0 {
		fmt.Printf("\nCount: %d\n", resultsCount)
	}

	fmt.Printf("Lookup time taken: %s\n", timeTaken.String())

	return false, nil
}

func userInput() (string, error) {
	fmt.Print("\nSearch term: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading user input string: %v", err)
	}

	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)

	return input, nil
}
