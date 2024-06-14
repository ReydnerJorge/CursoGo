package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudande struct {
	pessoa 
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Reydner", "Oliveira", 28, 160}
	fmt.Println(p1)

	e1 := estudande{p1, "Ads", "Una"}
	fmt.Println(e1.nome)
}