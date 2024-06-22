package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2)//buffer de canais
	canal <- "Olá Mundo!"
	canal <- "Programando em Go"

	mensagem := <-canal
	mensagem1 := <- canal 
	fmt.Println(mensagem, mensagem1)
}