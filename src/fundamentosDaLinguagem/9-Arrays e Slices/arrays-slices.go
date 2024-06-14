package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")
	
	//array
	// 3 maneiras de se utilizar um array
	var array1[5] string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5] int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...] int {100, 200, 300, 400, 500, 600} //[...] deixa o tamanho baseado na qtd. de valores
	fmt.Println(array3)

	//slice
	slice := [] int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18, 19, 20) //append para adicionar campos ao slice existente
	fmt.Println(slice)

	slice2 := array3[1:4]
	fmt.Println(slice2)

	array3[2] = 350
	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(slice))//TypeOf devolve o tipo 
	fmt.Println(reflect.TypeOf(array3))
}

/*
- arrays e uma lista de valores
- dentro de um array todos os valores tem q ser do mesmo tipo
- quantidde de posições fixas

- slice não tem tamanho fixo. 
- o tipo de dados tem q ser o mesmo para todas as posições
*/