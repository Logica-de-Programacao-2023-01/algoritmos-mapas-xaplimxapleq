package main

import (
	"fmt"
	"sort"
	"strings"
)

func encontrarAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {
		sorted := sortString(palavra)
		anagramas[sorted] = append(anagramas[sorted], palavra)
	}

	return anagramas
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func main() {
	palavras := []string{"tome", "gap", "barril", "tromba", "poggers", "eba", "dog"}

	anagramas := encontrarAnagramas(palavras)

	for _, grupo := range anagramas {
		fmt.Println(grupo)
	}
}
