package main

import "fmt"

var x [5]int
var y [6]int

func main() {
	//Array
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(len(x))

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	//Slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	//Muda o tamanho
	slice2 := append(slice, 6)
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 348756
	fmt.Println(slice[3])

	//Range function
	slice3 := []string{"banana", "apple", "mango"}

	for indice, valor := range slice3 {
		fmt.Println("Indice: ", indice, " Value: ", valor)
	}

	//Slice function
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatro queijos", "marguerita"}
	fatia := sabores[0:2]
	fatiaGrande := sabores[2:5]
	pizza := sabores[:]
	fmt.Println(fatia)
	fmt.Println(fatiaGrande)
	fmt.Println(pizza)

	//Slice function delete
	sabores = append(sabores[:2], sabores[4:]...)
	fmt.Println(sabores)

	//Append function
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}
	umaslice = append(umaslice, outraslice...) //... enumeration pega os valores do array
	fmt.Println(umaslice)

	//Make function (Slice when is full, create a new array and copy elements. More computability cost. Make set the range e capacity)
	sliceMake := make([]int, 5, 10)
	sliceMake[0], sliceMake[1], sliceMake[2], sliceMake[3] = 1, 2, 3, 4
	sliceMake[4] = 5
	sliceMake = append(sliceMake, 6)
	sliceMake = append(sliceMake, 7)
	sliceMake = append(sliceMake, 8)
	sliceMake = append(sliceMake, 9)
	sliceMake = append(sliceMake, 10)
	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))
	sliceMake = append(sliceMake, 11)
	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))

	//Slice multidimensional
	ss := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(ss)
	fmt.Println(ss[1][1])

	//Todo Slice tem um array subjacente(ponteiro/endereço) + len + cap -- problema com varios arrays subjacente
	primeiroSlice := []int{1, 2, 3, 4, 5}
	segundoSlice := append(primeiroSlice[:2], primeiroSlice[4:]...)
	fmt.Println(segundoSlice)
	fmt.Println(primeiroSlice)

	//Maps(Chave - Valor)
	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 44444

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"])

	fmt.Println(amigos["romario"])

	if sera, ok := amigos["fantasma"]; !ok {
		fmt.Println("Nao tem!")
		fmt.Println(sera, ok)
	} else {
		fmt.Println(sera)
	}

	//Maps(Ranges - Delete)
	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse e massa",
		18:  "idade de ir pra festa",
	}
	fmt.Println(qualquercoisa)

	for key, value := range qualquercoisa {
		fmt.Println(key, value)
	}

	delete(qualquercoisa, 123)
	fmt.Println(qualquercoisa)
}
