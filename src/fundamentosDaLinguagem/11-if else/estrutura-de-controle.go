package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle if/else")

	numero := 0
	if numero > 15{
		fmt.Println("Maior que 15")
	}else {
		fmt.Println("Menor ou igual a 10")
	}

	if outroNumero := numero; outroNumero > 0 { //if init 
		fmt.Println("Número é maior que 0")
	} else if numero < 10 {
		fmt.Println("Numero e menor que 10")
	} else {
		fmt.Println("Número e maior que 10")
	}
}
