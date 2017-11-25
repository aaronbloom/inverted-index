package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func storeRecord(path string, records *[]recordItem, termIndex *[]indexItem) {
	currentRecord := recordItem{len(*records), path}
	*records = append(*records, currentRecord)

	terms := strings.FieldsFunc(path, pathSplit)
	for i := 0; i < len(terms); i++ {
		addRecordToIndex(terms[i], currentRecord.id, termIndex)
	}
}

func pathSplit(r rune) bool {
	return r == '/' || r == '\\' || r == ' ' || r == '.'
}

func addRecordToIndex(term string, recordID int, termIndex *[]indexItem) {
	term = neutralString(term)
	for j := 0; j < len(*termIndex); j++ {
		if (*termIndex)[j].term == term {
			addRecordToTerm(termIndex, recordID, j)
			return
		}
	}

	insertNewTerm(term, recordID, termIndex)
}

func neutralString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}

func addRecordToTerm(termIndex *[]indexItem, recordID int, index int) {
	(*termIndex)[index].records = append((*termIndex)[index].records, recordID)
}

func insertNewTerm(term string, recordID int, termIndex *[]indexItem) {
	index := indexItem{
		id:      len(*termIndex),
		term:    term,
		records: []int{recordID},
	}
	*termIndex = append(*termIndex, index)
}

func indexPath(startPath string) ([]recordItem, []indexItem) {
	var totalCount int
	var matchCount int

	records := make([]recordItem, 0)
	termIndex := make([]indexItem, 0)

	startTime := time.Now()

	filepath.Walk(startPath, func(path string, f os.FileInfo, err error) error {
		totalCount++

		if strings.Contains(path, "node_modules") {
			fmt.Println("Skipping directory", path)
			return errors.New("Skipping directory")
		}
		if strings.HasSuffix(path, ".exe") {
			matchCount++
			storeRecord(path, &records, &termIndex)
			fmt.Printf("(%d) Visited: %s\n", matchCount, path)
		}

		return nil
	})

	timeTaken := time.Since(startTime)

	fmt.Printf("Indexing time taken: %s\n", timeTaken.String())
	fmt.Printf("%fs per item\n", timeTaken.Seconds()/float64(totalCount))

	fmt.Printf("Total items found: %d\n", totalCount)
	fmt.Printf("Total exes found: %d\n", matchCount)

	//fmt.Println("records", records)
	//fmt.Println("termIndex", termIndex)

	return records, termIndex
}
