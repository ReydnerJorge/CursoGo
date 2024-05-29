package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = (42 + 56 / 2) / 2 * 4 
	fmt.Println(x)
}

/*
exercício #4
Crie um tipo. O tipo subjacente deve ser int.
Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
Na função main:

    Demonstre o valor da variável "x"
    Demonstre o tipo da variável "x"
    Atribua 42 à variável "x" utilizando o operador "="
    Demonstre o valor da variável "x"


*/