package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	cliente1 := cliente{
		nome:      "Joao",
		sobrenome: "da Silva",
		fumante:   false,
	}

	cliente2 := cliente{"Joana", "Pereira", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)

	//Struct embutidos
	pessoa1 := pessoa{"Alfredo", 30}
	pessoa2 := profissional{
		pessoa{"Jones", 31},
		"Pizzaiolo",
		10000,
	}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.nome)

	//Struct anonimo
	anonimo := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 50,
	}
	fmt.Println(anonimo)
}
