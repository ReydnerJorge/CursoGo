package main

import "fmt"

const (
	a = iota 
	_ = iota
	c = iota * 2
	_ = iota
	e = iota
	_ = iota
	g = iota 
)

func main() {
	fmt.Println(a, c, e, g)
}

/*
Numa declaração de constantes, o identificador iota representa numeros sequenciais.
*/