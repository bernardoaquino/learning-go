package main

import "fmt"

func main() {
	// func (receiver) nome(parameters) (returns) { code }

	// Tudo em Go é PASS BY VALUE
	s := "abc"
	fmt.Println(s) //Pega o valor "abc" e não a referência s

	basica()
	argumento("Salve!")

	valor := soma1(10, 10)
	fmt.Println(valor)

	total, quantos, oi := soma2("Batata", 10, 10, 1, 2, 3, 5)
	fmt.Println(total, quantos, oi)

	//Defer - Deixa pra última hora(Quando o bloco de código chegar ao fim)
	defer fmt.Println("Com defer, veio primeiro")
	fmt.Println("Sem defer, veio depois")
	/* Exemplo:
	abreArquivo()
	defer fecharquivo()
	usaraquivo()
	*/
}

func basica() {
	fmt.Println("Olá, bom dia!")
}

func argumento(s string) {
	fmt.Println(s)
}

//Retorno definido
func soma1(x, y int) int {
	return x + y
}

//Parâmetro variádico(tem que ser o último) e multiplo retorno
func soma2(s string, x ...int) (int, int, string) { //Recebe um slice de int - Valores variáveis - Pode receber nenhum parâmetro variádico
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), s
}
