package main

import "fmt"

var w = "Opa" //Variável de escopo do package - Não consegue usar operador curto de atribuição fora de um bloco de código
var b int = 20

func main() {
	x := 10 //Operador curto atribui um tipo com base no valor da variavel(Operador de declaração - Marmota)
	y := "Olá bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x = 20 //Operador de atribuição
	fmt.Printf("x: %v, %T\n", x, x)
	x, z := 19, 30 //Operador de declaração deve declarar pelo menos 1 variável
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)

	//Tipos são estáticos, se uma variável x é int ela vai ser int até o fim do programa(Tipagem stática, não muda o tipo da variável)
	fmt.Printf("b: %v, %T\n", b, b)
}
