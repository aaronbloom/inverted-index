package main

import (
	"flag"
	"fmt"
)

func main() {

	flag.Parse()
	root := flag.Arg(0)

	if len(root) < 1 {
		fmt.Println("Please provide a path to traverse")
		return
	}

	records, termIndex := indexPath(root)
	search(records, termIndex)

}
