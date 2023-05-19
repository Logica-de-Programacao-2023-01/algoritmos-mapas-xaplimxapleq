package main

import "fmt"

func contarcaracteres(str string) map[rune]int {
	mapa := make(map[rune]int)
	for _, char := range str {
		mapa[char]++
	}
	return mapa
}

func main() {
	texto := "sobradinho"
	resultado := contarcaracteres(texto)
	fmt.Println(resultado)
}
