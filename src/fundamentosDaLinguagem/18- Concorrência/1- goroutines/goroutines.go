package main

import (
	"fmt"
	"time"
)

func main() {
	//Concorrência != Paralelismo
	go escrever("Olá Mundo!") //goroutines
	escrever("Programa em Go!")
}

func escrever(texto string) {
	for i := 0; i < 10; i++{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}