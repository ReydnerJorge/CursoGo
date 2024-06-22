package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}

/* Função que vai encapsular uma chamada para
um goroutine e deolver em forma de um canal */
