package main

import "fmt"

func main() {
	x := "oi"
	y := "bom dia"

	z := fmt.Sprint(x, " ", y) //Não adiciona espaço entre os operadores automaticamente

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
