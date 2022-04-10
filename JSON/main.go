package main

import (
	"encoding/json"
	"fmt"
)

//json tags que define informacoes que o propriedade se refere no JSON
type Pessoa struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

func main() {
	//Marshal converte em JSON e vai para variavel - Enconder, encode manda direto
	jamesBond := Pessoa{"James", "Bond", 40, "Agente Secreto", 10000.50}
	darthVader := Pessoa{"Anakin", "Skywalker", 50, "Vilao", 50000.83}

	j, err := json.Marshal(jamesBond)
	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(darthVader)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j)) //Para exportar para JSON os nomes devem ser maiusculos
	fmt.Println(string(d))

	//Unmarshal converte JSON em struct
	sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","Contabancaria":10000.5}`)
	var jamesBond2 Pessoa
	err2 := json.Unmarshal(sb, &jamesBond2)
	if err2 != nil {
		fmt.Println("error: ", err2)
	}
	fmt.Println(jamesBond2)
	fmt.Println(jamesBond2.Profissao)
}
