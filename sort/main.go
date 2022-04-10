package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorPotencia []carro

func (x ordenarPorPotencia) Len() int {
	return len(x)
}

func (x ordenarPorPotencia) Less(i, j int) bool {
	return x[i].potencia < x[j].potencia
}

func (a ordenarPorPotencia) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ordenarPorConsumo []carro

func (x ordenarPorConsumo) Len() int {
	return len(x)
}

func (x ordenarPorConsumo) Less(i, j int) bool {
	return x[i].consumo > x[j].consumo
}

func (a ordenarPorConsumo) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ordenarPorLucroProDonoDoPosto []carro

func (x ordenarPorLucroProDonoDoPosto) Len() int {
	return len(x)
}

func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool {
	return x[i].consumo < x[j].consumo
}

func (a ordenarPorLucroProDonoDoPosto) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	ss := []string{"abobora", "maca", "banana", "uva", "manga"}
	fmt.Println(ss)
	sort.Strings(ss)
	fmt.Println(ss)

	si := []int{123, 456, 321, 789, 987, 654}
	fmt.Println(si)
	sort.Ints(si)
	fmt.Println(si)

	//Sort customizavel
	carros := []carro{{"Chevete", 900, 5}, {"Porsche", 300, 5}, {"Fusca", 20, 30}}
	fmt.Println(carros)
	sort.Sort(ordenarPorPotencia(carros)) //Chama os metodos acima porque esse objeto virou um sort
	fmt.Println(carros)
	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println(carros)
	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))
	fmt.Println(carros)
}
