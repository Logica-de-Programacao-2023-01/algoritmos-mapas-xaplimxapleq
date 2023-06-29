package main

import "fmt"

func pares(s []int) map[int]int {
	frequenciaS := make(map[int]int)

	for _, num := range s {
		if num%2 == 0 {
			frequenciaS[num]++
		}
	}

	for numero, frequencia := range frequenciaS {
		frequenciaS[numero] = frequencia / 2
	}

	return frequenciaS
}

func main() {
	s := []int{0, 2, 3, 2, 2, 1, 2, 1, 5, 3, 2, 4, 2, 4}
	fmt.Println(pares(s))
}
