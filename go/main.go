package main

import (
	"fmt"
	"os"
)

func main() {
	dictionary := MakeDictionary()
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Letters not provided.")
		return
	}
	words := dictionary.FindWords(args[1])
	for _, word := range words {
		fmt.Println(word)
	}
}
