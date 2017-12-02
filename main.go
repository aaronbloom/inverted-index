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

	fmt.Printf("Indexing path provided: %s\n", *path)

	index := indexPath(*path)
	search(index)

}
