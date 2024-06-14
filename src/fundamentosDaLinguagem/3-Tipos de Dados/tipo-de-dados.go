package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 11111
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// UINT8 = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)

	//numeros reais
	var numeroreal1 float32 = 123.45
	var numeroreal2 float64 = 1230000000.45
	numeroreal3 := 1234.54
	fmt.Println(numeroreal1)
	fmt.Println(numeroreal2)
	fmt.Println(numeroreal3)

	//strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'R'
	fmt.Println(char)

	texto := 5
	fmt.Println(texto)

	var boolean bool
	fmt.Println(boolean)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}

/*
int8, int16, int32, int64 
uint8, uint16, uint32, uint64 - tipos de numeros inteiros
a diferença entre cada um e a quantidade de bits que cada
um ocupa e tem o int tmb que vai usar a arquitetura do pc
para basear qual dos tipos de int será usado. Ex: arquitetura 64bits = int64

float32, float64 - tipos de numeros com ponto flutuante
*/