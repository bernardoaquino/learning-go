package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia!")
}

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

	//Métodos(Função anexada a um tipo) - Tipos são tudo - Adiciona funcionalidade a um tipo especificamente
	mauricio := pessoa{"Maurício", 30}
	mauricio.oibomdia()

	//Função anônima - Declara e chama ao mesmo tempo
	x := 10
	func(x int) {
		fmt.Println(x * 10)
	}(x)

	//Funções como expressão
	y := func(x int) int {
		return x * 10
	}
	fmt.Println(y(x))

	//Função que retorna uma função
	w := retornaUmaFuncao()
	z := w(3)
	fmt.Println(z)
	fmt.Println(retornaUmaFuncao()(3))

	//Callback passa uma função como argumento
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

func retornaUmaFuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
