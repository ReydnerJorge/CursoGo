package main

import (
	"fmt"
)

func main() {
	x := 13 //operador de declaração
	y := "Olá boa tarde"
	z := 17

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)
	fmt.Printf("z: %v, %T\n", z, z)

	x = 20 // operador de atribuição
	fmt.Printf("x: %v, %T\n", x, x)
}

/*
O package printf define um analisador que verifica a consistência de strings e argumentos de formato Printf
*/