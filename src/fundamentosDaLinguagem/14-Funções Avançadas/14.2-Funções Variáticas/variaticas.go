package main

import "fmt"

func soma(numeros ...int) int { //função variática e com um slice
	total  := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
 totalDaSoma := soma(1, 2, 3, 10, 15, 25, 55, 101)
 fmt.Println(totalDaSoma)

 escrever("Olá Mundo", 10, 2, 3, 4, 5, 6, 7)
}