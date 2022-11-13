package main

import "fmt"

type Document interface {
	Doc() string
}

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

func (pp PrivatePerson) Doc() string {
	return fmt.Sprintf("Ola meu cpf é %s ", pp.CPF)
}

func (jp JuridicalPerson) Doc() string {
	return fmt.Sprintf("Ola meu cnpj é %s ", jp.CNPJ)
}

type JuridicalPerson struct {
	People
	CompanyName string
	CNPJ        string
}

func (p People) String() string {
	return fmt.Sprintf("Ola meu nome é %s e eu tenho %d anos.", p.Name, p.Age)
}

func show(d Document) {
	// if d, ok := d.(PrivatePerson); ok {
	// 	fmt.Println(d.LastName) //assercao de tipo, da para utilizar os atributos da quele tipo
	// }

	switch d.(type) { //forma de fazer assercao para varios tipos de dados
	case PrivatePerson:
		fmt.Println(d.(PrivatePerson).LastName)
	case JuridicalPerson:
		fmt.Println(d.(JuridicalPerson).CompanyName)
	default:
		fmt.Println("Tipo desconhecido")

	}
	fmt.Println(d)
	fmt.Println(d.Doc())
}

func main() {
	pp := PrivatePerson{
		People{
			Name:   "Pedro",
			Age:    21,
			Status: true,
		},
		"Dalpa",
		"000.000.000-00",
	}

	show(pp)
}
