package main

import (
	"fmt"
	"strings"
)

func countLettersAndWords(s string) (map[string]int, int) {
	letterCount := make(map[string]int)
	wordCount := 0

	words := strings.Fields(s)

	for _, word := range words {
		wordCount++

		for _, letter := range word {
			letterCount[string(letter)]++
		}
	}

	return letterCount, wordCount
}

func main() {
	s := "quero abobra com leite"
	letterCount, wordCount := countLettersAndWords(s)

	fmt.Println("Contagem de letras:")
	for letter, count := range letterCount {
		fmt.Printf("%s: %d\n", letter, count)
	}

	fmt.Println("Contagem de palavras:", wordCount)
}
