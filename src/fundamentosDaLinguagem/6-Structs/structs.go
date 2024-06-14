package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo Structs")
// primeira forma de declara uma structs
	var u usuario 
	u.nome = "Reydner"
	u.idade = 28
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Itumbiara", 44}

// segunda forma por inferencia 
	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

// terceira forma
	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)
}


