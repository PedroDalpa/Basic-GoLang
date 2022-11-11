package main

import "fmt"

type People struct {
	Name   string
	Age    uint8
	Status bool
	Mother *People
}

func main() {
	names := []string{"Tiago", "Pedro", "Maria", "Joao"}
	fmt.Println(names, len(names), cap(names)) // capacidade igual a inicializada

	names = append(names, "Yasmin")
	fmt.Println(names, len(names), cap(names)) // aumenta mais a capacidade do que a necessária, para nao ter que aumentar toda hora

	names = append(names, "Diego")
	fmt.Println(names, len(names), cap(names))

	numbers := make([]int, 10, 20) // vai gerar um array com as 10 primeiras posições ja preenchidas com o valor 0 do tipo, int = 0, string =""....
	numbers = append(numbers, 32)  // ira adicionar apos as 10 casas ja preenchidas
	fmt.Println(numbers, len(numbers), cap(numbers))

	ages := make(map[string]uint8) // map
	ages["Tiago"] = 38
	ages["Pedro"] = 21
	ages["Ana"] = 32

	value, ok := ages["Lucas"] //se nao encontrar o ok vem como false
	fmt.Println(value, ok)
	fmt.Println(ages)

	var p People

	p.Name = "Pedro"
	p.Age = 21
	fmt.Println(p)

	p = People{
		Name:   "Tiago",
		Age:    38,
		Status: true,
		Mother: &People{
			Name: "teste",
			Age:  62,
		},
	}
	fmt.Println(p, p.Mother)
}
