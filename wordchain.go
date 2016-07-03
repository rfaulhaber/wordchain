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
	startWord := args[0]
	goalWord := args[1]

	fmt.Println(args)
	fmt.Println("start word: ", startWord, "goal word:", goalWord)

	if *indexFlagPtr == "" {
		fmt.Println("Please specify a word list!")
	} else if len(startWord) != len(goalWord) {
		fmt.Println("Start word and goal word must be of equal length!")
	} else {
		root := node.BuildTree(*indexFlagPtr)

		currentWord := startWord

		fmt.Print(startWord)
		for currentWord != startWord {
			currentWord = suggestNextWord(currentWord, goalWord, root)
			fmt.Print(currentWord)
		}
	}
}

func suggestNextWord(currentWord string, goalWord string, root *node.Node) string {
	thisWord := currentWord

	for !isWordInTree(thisWord, root) {
		for i, c := range goalWord {
			if string(thisWord[i]) == string(c) {
				continue
			} else {
				thisWord = stringReplace(thisWord, i, string(c))
				break
			}
		}
	}

	return thisWord
}

func isWordInTree(word string, root *node.Node) bool {
	return root.HasWord(word)
}

func stringReplace(word string, index int, newChar string) string {
	return word[:index] + string(newChar) + word[index + 1:]
}
