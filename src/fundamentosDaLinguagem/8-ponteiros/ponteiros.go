package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = 10

	fmt.Println(variavel1, variavel2)

	variavel1++ 
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // & para ver o endereço referente a memoria

	fmt.Println(variavel3, ponteiro) // utilizasse o *ponteiro para ver a desreferenciação
}