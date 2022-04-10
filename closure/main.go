package main

import "fmt"

func main() {
	//Closure - Captura um contexto e utiliza em outro
	a := i()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := i() //Gera outro valor para x

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func i() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}
