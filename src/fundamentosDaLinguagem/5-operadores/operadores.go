package main

import "fmt"

func main() {
	// aritméticos
	soma := 10 + 10
	subracao := 15 - 10
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subracao, divisao, multiplicacao, restoDaDivisao)
	fmt.Println("---------------------")

	// atribuição
	var variavel string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel, variavel2)

	// operadores relacionais
	fmt.Println("---------------------")
	fmt.Println(1 > 2) //maior
	fmt.Println(1 >= 2) //maior ou igual
	fmt.Println(1 == 2) //igual
	fmt.Println(1 <= 2) // menor ou igual
	fmt.Println(1 > 2) // maior que
	fmt.Println(1 < 2) // menor que
	fmt.Println(1 != 2) //diferente

	//operadores logicos
	fmt.Println("---------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//operadores unarios
	fmt.Println("---------------------")
	numero := 10
	numero++
	numero +=15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20

	numero *= 3 //numero = numero * 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)
}

/*
+ = adição
- = subtração
/ = divisão
* = multiplicação
% = resto da divisão
*/