package main

import (
	"flag"
	"fmt"

	"github.com/aaronbloom/inverted-index/recordkeeper"
)

func main() {

	mode := flag.String("m", "c", "mode - either 'create' (c) or 'search' (s)")
	path := flag.String("p", "./", "path to traverse and index files")
	indexFile := flag.String("i", "./", "index file")
	flag.Parse()

	if *mode == "c" || *mode == "create" {
		if err := create(*path); err != nil {
			fmt.Printf("error in create mode: %v", err)
		}
	}

	if *mode == "s" || *mode == "search" {
		if err := loadIndexAndSearch(*indexFile); err != nil {
			fmt.Printf("error in search mode: %v", err)
		}
	}
}

func loadIndexAndSearch(indexPath string) error {
	if len(indexPath) < 1 {
		return fmt.Errorf("please provide a valid path to traverse")
	}

	index, err := recordkeeper.ReadFromFile(indexPath)
	if err != nil {
		return fmt.Errorf("could not read index file: %v", err)
	}

	search(&index)

	return nil
}

func create(path string) error {
	if len(path) < 1 {
		return fmt.Errorf("please provide a path to traverse")
	}

	ignoredDirectories := []string{"node_modules"}
	fileExtensionMatches := []string{".exe", ".sh", ".bat", ".cmd"}

	fmt.Printf("Indexing path provided: %s\n", path)

	index := indexPath(path, ignoredDirectories, fileExtensionMatches)

	indexFilePath := "./recordsindex"
	fmt.Printf("Saving index: %s\n", indexFilePath)

	err := index.SaveToFile(indexFilePath)
	if err != nil {
		return fmt.Errorf("could not save index to file: %v", err)
	}

	return nil
}
