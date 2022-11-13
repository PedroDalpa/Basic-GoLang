package main

import "fmt"

type People struct {
	Name   string
	Age    uint8
	Status bool
}

type PrivatePerson struct {
	People
	LastName string
	CPF      string
}

type JuridicalPerson struct {
	People
	CompanyName string
	CNPJ        string
}

func (p People) String() string {
	return fmt.Sprintf("Ola meu nome é %s e eu tenho %d anos.", p.Name, p.Age)
	// pode continuar acessando direto, ah menos que a estrutura filha possua o mesmo atributo da estrutura pai
}

func main() {
	p := PrivatePerson{
		People{
			Name:   "Pedro",
			Age:    21,
			Status: true,
		},
		"Dalpa",
		"000.000.000-00",
	}

	fmt.Println(p)

}
