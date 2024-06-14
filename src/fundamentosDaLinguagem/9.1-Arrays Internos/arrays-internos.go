package main

import "fmt"

func main() {
	fmt.Println("Arrays Internos")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))//length
	fmt.Println(cap(slice3))//capacidade

	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 =append(slice4, 12)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}

// Array e uma lista com tamanho fixo e o slice uma lista sem tamanho fixo