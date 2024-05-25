package main

import (
	"fmt"
)

func main() {
	x := 10 //operador de declaração
	y := "Olá boa tarde"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x = 20 // operador de atribuição
	fmt.Printf("x: %v, %T\n", x, x)
}

/*
O package printf define um analisador que verifica a consistência de strings e argumentos de formato Printf
*/