package main

import (
	"fmt"
	"github.com/aaronbloom/inverted-index/recordkeeper"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func pathSplit(r rune) bool {
	return r == '/' || r == '\\' || r == ' ' || r == '.'
}

func neutralString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}

func indexPath(startPath string) InverseIndex {
	var totalCount int
	var matchCount int

	records := recordkeeper.Create()
	termIndex := make([]indexItem, 0)

	inverseIndex := InverseIndex{
		indexItems:  termIndex,
		recordItems: records,
	}

	startTime := time.Now()

	filepath.Walk(startPath, func(path string, f os.FileInfo, err error) error {
		totalCount++

		if strings.Contains(path, "node_modules") {
			fmt.Println("Skipping directory", path)
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".exe") {
			matchCount++
			inverseIndex.StoreRecord(path)
			fmt.Printf("(%d) Visited: %s\n", matchCount, path)
		}

		return nil
	})

	timeTaken := time.Since(startTime)

	fmt.Printf("Indexing time taken: %s\n", timeTaken.String())
	fmt.Printf("%fs per item\n", timeTaken.Seconds()/float64(totalCount))

	fmt.Printf("Total items found: %d\n", totalCount)
	fmt.Printf("Total exes found: %d\n", matchCount)

	return inverseIndex
}
