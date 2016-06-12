package main

import (
	"flag"
	"fmt"
    "os"
	"wordchain/node"
)

func main() {
	indexFlagPtr := flag.String("i", "", "Path to a text file containing a list of properly formatted words")
	flag.Parse()
    args := os.Args[2:]

	fmt.Println(args)

	if *indexFlagPtr == "" {
		fmt.Println("Please specify a word list!")
	} else {
		root := node.BuildTree(*indexFlagPtr)
	}
}
