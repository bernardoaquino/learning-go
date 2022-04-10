package main

import "fmt"

func main() { //Main é uma Goroutine
	//Canais comunicam/transmitem dados entre Goroutines
	canal := make(chan int)
	go func() {
		canal <- 42
	}()
	fmt.Println(<-canal)

	//Buffer
	canal2 := make(chan int, 1)
	canal2 <- 42
	fmt.Println(<-canal2)

	//Canais direcionais(SendChannel(send chan<-) e ReceiveChannel(receive <-chan) Bidirecionais)
	canal3 := make(chan int)
	go send(canal3)
	receive(canal3)
}

func send(s chan<- int) {
	s <- 42
}

func receive(r <-chan int) {
	fmt.Println("O valor recebido do canal foi:", <-r)
}
