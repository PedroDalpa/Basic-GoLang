package main

import "fmt"

type Category struct {
	Name   string
	Father *Category
}

func (c Category) HasParent() bool { // atrelado a struct
	return c.Father != nil
}

func (c *Category) SetParent(father *Category) { // com o ponteiro ele vai como referencia entao altero os valores, sem ponteiro apenas recupero o valor
	c.Father = father //ao chamar essa funcao cria o atributo father vazio
}

type People struct {
	Name   string
	Age    uint8
	Status bool
}

func (p People) String() string {
	return fmt.Sprintf("Ola meu nome é %s e eu tenho %d anos.", p.Name, p.Age)
}

func main() {
	p := People{
		Name:   "Tiago",
		Age:    38,
		Status: true,
	}

	cat := Category{Name: "Cat 1"}
	cat.SetParent(&Category{Name: "Cat father"})
	fmt.Println(p)
	fmt.Println(cat.Father)
	if !cat.HasParent() {
		fmt.Printf("Nao tem pai")
	}

}
