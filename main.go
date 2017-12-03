package main

import (
	"flag"
	"fmt"
)

func main() {

	path := flag.String("p", "./", "path to traverse and index files")
	flag.Parse()

	if len(*path) < 1 {
		fmt.Println("Please provide a path to traverse")
		return
	}

	ignoredDirectories := []string{"node_modules"}
	fileExtensionMatches := []string{".exe", ".sh", ".bat", ".cmd"}

	fmt.Printf("Indexing path provided: %s\n", *path)

	index := indexPath(*path, ignoredDirectories, fileExtensionMatches)
	search(index)

}
