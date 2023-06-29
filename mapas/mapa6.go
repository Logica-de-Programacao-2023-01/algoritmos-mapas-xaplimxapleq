package main

import (
	"fmt"
	"strings"
)

func contarPalavras(s string) map[string]int {
	mapa := make(map[string]int)
	palavras := strings.Fields(s)

	for _, palavra := range palavras {
		mapa[palavra]++
	}

	return mapa
}

func unirMapas(mapas []map[string]int) map[string]int {
	bigMap := make(map[string]int)

	for _, m := range mapas {
		for chave, valor := range m {
			bigMap[chave] += valor
		}
	}

	return bigMap
}

func main() {
	a := "adoro festa"
	b := "ja comi pneu"
	c := "ta me macetando dog"

	mapas := []map[string]int{contarPalavras(a), contarPalavras(b), contarPalavras(c)}

	fmt.Println(unirMapas(mapas))
}
