package main

import "fmt"

func main() {
	//Aqui começa
	//... funções variáticas, número ilimitado de parâmetros(Todos os tipo em GO implementam uma interface vazia)
	numeroBytes, erros := fmt.Println("Hello World", "Oi galera!", 100) //Utilizando _ você pode descartar variáveis que você não irá usar
	fmt.Println(numeroBytes, erros)
	//Aqui termina
}
