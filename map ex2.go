package main
func combinarmapas(mapa1, mapa2 map[string]interface{}) map[string]interface{} {
	novomapa := make(map[string]interface{})

	for chave, valor := range mapa1 {
		novomapa[chave] = valor
	}
}
for chave, valor :=range mapa2 {
	novomapa[chave] = valor
	return novomapa

}

func main() {

}
