package main

import (
	"fmt"
	"os"
)

func main() {
	dictionary := MakeDictionary()
	args := os.Args
	var letters string
	if len(args) < 2 {
		var err error
		letters, err = getLetters()
		if err != nil {
			fmt.Println("Error getting word:", err)
			return
		}
	} else {
		letters = args[1]
	}
	words := dictionary.FindWords(letters)
	for _, word := range words {
		fmt.Println(word)
	}
}
