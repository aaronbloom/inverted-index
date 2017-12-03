package main

import (
	"fmt"
	"github.com/aaronbloom/inverted-index/recordkeeper"
	"github.com/aaronbloom/inverted-index/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func indexPath(startPath string, ignoredDirectories []string, fileExtensionMatches []string) recordkeeper.RecordIndex {
	var totalCount int
	var matchCount int

	recordkeeper := recordkeeper.CreateRecordIndex()

	startTime := time.Now()

	filepath.Walk(startPath, func(path string, f os.FileInfo, err error) error {
		totalCount++

		for _, directoryName := range ignoredDirectories {
			if strings.Contains(path, directoryName) {
				fmt.Println("Skipping directory", path)
				return filepath.SkipDir
			}
		}

		for _, extension := range fileExtensionMatches {
			if strings.HasSuffix(path, extension) {
				matchCount++
				recordkeeper.StoreRecord(path, utils.FilePathTokenizer, utils.LowerCaseFilter)
				fmt.Printf("(%d) Visited: %s\n", matchCount, path)
			}
		}

		return nil
	})

	timeTaken := time.Since(startTime)

	fmt.Printf("Indexing time taken: %s\n", timeTaken.String())
	fmt.Printf("%fs per item\n", timeTaken.Seconds()/float64(totalCount))

	fmt.Printf("Total items found: %d\n", totalCount)
	fmt.Printf("Total exes found: %d\n", matchCount)

	return recordkeeper
}
