package main

import (
	"flag"
	"fmt"

	"github.com/aaronbloom/inverted-index/recordkeeper"
)

func main() {

	mode := flag.String("m", "", "mode - either 'create' (c) or 'search' (s)")
	path := flag.String("p", "./", "path to traverse and create index files (create mode only)")
	indexDir := flag.String("i", "./", "index file directory")
	flag.Parse()

	if len(*indexDir) < 1 {
		fmt.Println("Please provide an index file directory (-i). See --help")
		return
	}

	recordsIndexFile := *indexDir + "_records"
	termIndexFile := *indexDir + "_terms"

	if *mode == "c" || *mode == "create" {
		if err := create(*path, recordsIndexFile, termIndexFile); err != nil {
			fmt.Printf("error in create mode: %v", err)
		}
		return
	}

	if *mode == "s" || *mode == "search" {
		if err := loadIndexAndSearch(recordsIndexFile, termIndexFile); err != nil {
			fmt.Printf("error in search mode: %v", err)
		}
		return
	}

	fmt.Println("Please provide a valid mode (-m). See --help")
	return
}

func loadIndexAndSearch(recordsIndexFile string, termIndexFile string) error {
	fmt.Printf("Loading record index from %s\n", recordsIndexFile)
	fmt.Printf("Loading term index from %s\n", termIndexFile)
	index, err := recordkeeper.ReadFromFile(recordsIndexFile, termIndexFile)
	if err != nil {
		return fmt.Errorf("could not read index file: %v", err)
	}

	for {
		stop, err := search(index)
		if err != nil {
			return err
		}
		if stop {
			return nil
		}
	}
}

func create(path string, recordsIndexFile string, termIndexFile string) error {
	if len(path) < 1 {
		return fmt.Errorf("please provide a path to traverse")
	}

	ignoredDirectories := []string{"node_modules"}
	fileExtensionMatches := []string{".exe", ".sh", ".bat", ".cmd"}

	fmt.Printf("Indexing path provided: %s\n", path)

	index := indexPath(path, ignoredDirectories, fileExtensionMatches)

	fmt.Printf("Saving records index: %s\n", recordsIndexFile)
	fmt.Printf("Saving terms index: %s\n", termIndexFile)

	err := index.SaveToFile(recordsIndexFile, termIndexFile)
	if err != nil {
		return fmt.Errorf("could not save index to file: %v", err)
	}

	return nil
}
