package main

import "fmt"

func somaMapa(m map[string]int) int {
	soma := 0

	for _, valor := range m {
		soma += valor
	}

	return soma
}

func main() {
	m := map[string]int{
		"a": 13,
		"b": 22,
		"c": 17,
	}

	fmt.Println(somaMapa(m))
}
