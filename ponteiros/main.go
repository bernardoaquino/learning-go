package main

import "fmt"

func main() {
	//Ponteiro é uma variável que armazena o valor de um endereço de memória
	x := 10
	y := &x
	fmt.Println(y)
	fmt.Println(*y) //dreferencing - Aponta para o valor que está no endereço
	fmt.Println(*&x)
	fmt.Printf("%T, %T\n", x, y)

	w := 0
	z := &w
	*z++
	fmt.Println(w, z)

	//Utiliza o ponteiro passando como parametro para funcao, para nao precisar realizar uma copia(Custo computacional)
	k := 0
	estaRecebeValor(k)
	estaRecebePonteiro(&k)
	fmt.Println(k)
}

func estaRecebeValor(k int) {
	k++
	fmt.Println("Na funcao: ", k)
}

func estaRecebePonteiro(k *int) {
	*k++
}
