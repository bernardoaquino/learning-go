package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//Concorrência: Padrão de design, onde várias coisas executam de maneira independente e ao mesmo tempo
	//Paralelismo: Código concorrente sendo executado em um sistema que possui vários processadores

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	//Goroutines = "Threads" = Linhas/Encadeamentos de funções
	wg.Add(2)
	go func1()
	go func2()

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Wait()

	//Condição de corrida
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totalGoRoutines := 10

	var wg2 sync.WaitGroup
	wg2.Add(totalGoRoutines)

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			v := contador
			runtime.Gosched() // yield - Deixa a thread aqui, vai executar outra e volta pra cá depois
			v++
			contador = v
			wg2.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg2.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)

	//Mutex(Tranca um variável/Trecho de código e somente 1 Goroutine pode utilizar o valor dela naquele momento)
	var mu sync.Mutex
	contador2 := 0
	totalGoRoutines2 := 10

	var wg3 sync.WaitGroup
	wg3.Add(totalGoRoutines2)

	for i := 0; i < totalGoRoutines2; i++ {
		go func() {
			mu.Lock()
			v := contador2
			runtime.Gosched() // yield
			v++
			contador2 = v
			mu.Unlock()
			wg3.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg3.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador2)
}

func func1() {
	for i := 0; i < 5; i++ {
		fmt.Println("Func1:", i)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 5; i++ {
		fmt.Println("Func2:", i)
	}
	wg.Done()
}
