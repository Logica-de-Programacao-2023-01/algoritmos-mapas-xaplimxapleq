package main

import (
	"fmt"
)

func frequenciaLetras(s string) map[string]int {
	frequencias := make(map[string]int)

	for _, char := range s {
		// Ignora caracteres de espaço em branco
		if char == ' ' {
			continue
		}

		letra := string(char)
		frequencias[letra]++
	}

	return frequencias
}

func main() {
	fmt.Println(frequenciaLetras("xãplimxapleq ou xapleq"))
}
