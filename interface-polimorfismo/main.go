package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesArrancados int
	salário          float64
}

type arquiteto struct {
	pessoa
	tipoDeConstrução string
	tamanhoDaLoucura string
}

func (x dentista) oiBomDia() {
	fmt.Println("Meu nome é", x.nome, "eu já arranquei", x.dentesArrancados, "dentes e bom dia!")
}

func (x arquiteto) oiBomDia() {
	fmt.Println("Meu nome é", x.nome, "bom dia!")
}

//Interface determina o conjunto de métodos necessários para implementar ela
type gente interface {
	oiBomDia()
}

func serHumano(g gente) {
	g.oiBomDia()

	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho:", g.(dentista).salário)
	case arquiteto:
		fmt.Println("Eu construo:", g.(arquiteto).tipoDeConstrução)
	}
}

func main() {
	//Polimorfismo - Mesma função em tipos diferentes
	mrdentista := dentista{pessoa{"Mister Dente", "da Silva", 50}, 893, 1000}
	mrpredio := arquiteto{pessoa{"Mister prédio", "sobrenome", 51}, "Hospital", "Demais"}

	serHumano(mrdentista)
	serHumano(mrpredio)

	mrdentista.oiBomDia()
	mrpredio.oiBomDia()
}
