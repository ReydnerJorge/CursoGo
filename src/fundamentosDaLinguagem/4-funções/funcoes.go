package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1 int8, n2 int8) (int8, int8){
	adicao := n1 + n2
	subtracao := n1 - n2
	return adicao, subtracao
}

func main() {
	soma := somar(10, 30)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função 12")
	fmt.Println(resultado)

	resultadoAdicao, resultadoSubtracao := calculosMatematicos(10, -15)
	fmt.Println(resultadoAdicao, resultadoSubtracao)
}

/*
Funções são uma serie de passos que vão ser seguidos pelo
programa
*/