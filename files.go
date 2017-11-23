package main

import (
	"bufio"
	"time"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var totalCount int
var totalFileCount int
var totalExeCount int

var records []recordItem
var termIndex []indexItem

func storeRecord(path string) {
	currentRecord := recordItem{
		id: len(records),
		content: path,
	}
	records = append(records, currentRecord)

	terms := strings.FieldsFunc(path, pathSplit);
	for i := 0; i < len(terms); i++ {
		addRecordToIndex(terms[i], currentRecord.id)
	}
}

func addRecordToIndex(term string, recordID int) {
	term = neutralString(term)
	for j := 0; j < len(termIndex); j++ {
		if (termIndex[j].term == term) {
			addRecordToTerm(&termIndex[j], recordID)
			return
		}
	}

	insertNewTerm(term, recordID)
}

func neutralString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}

func pathSplit(r rune) bool {
	return r == '/' || r == '\\' || r == ' ' || r == '.'
}

func addRecordToTerm(term *indexItem, recordID int) {
	term.records = append(term.records, recordID)
}

func insertNewTerm(term string, recordID int) {
	index := indexItem{
		id: len(termIndex),
		term: term,
		records: []int{recordID},
	}
	termIndex = append(termIndex, index)
}

func visit(path string, f os.FileInfo, err error) error {
	totalCount++
	if strings.HasSuffix(path, ".exe") {
		totalExeCount++
		storeRecord(path)
		fmt.Printf("(%d) Visited: %s\n", totalExeCount, path)
	}

	return nil
}

func indexPath(startPath string) {
	startTime := time.Now()
	
	filepath.Walk(startPath, visit)

	timeTaken := time.Since(startTime)
	fmt.Printf("Indexing time taken: %s\n", timeTaken.String())
	fmt.Printf("%fs per item\n", timeTaken.Seconds() / float64(totalCount))

	fmt.Printf("Total items found: %d\n", totalCount)
	fmt.Printf("Total exes found: %d\n", totalExeCount)

	//fmt.Println("records", records)
	//fmt.Println("termIndex", termIndex)
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

func search() {
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

	search() // loop
}

func main() {
	flag.Parse()
	root := flag.Arg(0)

	if len(root) < 1 {
		fmt.Println("Please provide a path to traverse")
		return
	}

	records = make([]recordItem, 0)
	termIndex = make([]indexItem, 0)

	indexPath(root)
	search()

}

type recordItem struct {
	id int
	content string
}

type indexItem struct {
	id int
	term string
	records []int
}