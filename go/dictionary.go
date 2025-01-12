package main

import (
	"bufio"
	"log"
	"os"
)

func LoadWords() *[]string {
	file, err := os.Open("../wordlist.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var wordList []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	return &wordList
}

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEndOfWord = true
}

func (t *Trie) searchWords(node *TrieNode, letters map[rune]bool, path []rune, result *[]string, centerLetter rune, containsCenterLetter bool) {
	if node.isEndOfWord && containsCenterLetter {
		*result = append(*result, string(path))
	}

	for ch, child := range node.children {
		if letters[ch] {
			if ch == centerLetter {
				t.searchWords(child, letters, append(path, ch), result, centerLetter, true)
			} else {
				t.searchWords(child, letters, append(path, ch), result, centerLetter, containsCenterLetter)
			}
		}
	}
}

func (t *Trie) FindWords(letters string) []string {
	lettersSet := make(map[rune]bool)
	for _, ch := range letters {
		lettersSet[ch] = true
	}
	centerLetter := rune(letters[0])
	result := []string{}
	t.searchWords(t.root, lettersSet, []rune{}, &result, centerLetter, false)
	return result
}

func MakeDictionary() *Trie {
	dictionary := NewTrie()
	wordList := LoadWords()
	for _, word := range *wordList {
		if len(word) >= 4 {
			dictionary.Insert(word)
		}
	}
	return dictionary
}
