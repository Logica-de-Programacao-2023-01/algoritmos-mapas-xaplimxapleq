package main

import "fmt"

func calcularMediaDespesas(despesas map[string]float64) float64 {
	total := 0.0
	for _, valor := range despesas {
		total += valor
	}

	media := total / float64(len(despesas))
	return media
}

func main() {
	despesas := map[string]float64{
		"xaplim": 200,
		"xapleq": 300,
		"xaploq": 400,
	}

	mediaDespesas := calcularMediaDespesas(despesas)
	fmt.Printf("A média das despesas é: %.2f reais\n", mediaDespesas)
}
