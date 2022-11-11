package main

import "fmt"

func main() {
	names := []string{"Tiago", "Pedro", "Maria", "Joao"}

	// nao e a melhor forma pois precisa cuidar do índice sempre
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// nao precisa cuidar do índice (Tipo um forEach(value))
	for k, name := range names {
		fmt.Println(k, name)
	}
}
