package main

import (
	"fmt"
	"time"
)

func main() {
		canal := make(chan string)

		go escrever("Olá Mundo", canal)

		for {
			mensagem, aberto := <-canal
			if !aberto {
				break
			}
			fmt.Println(mensagem)
		}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		canal <- texto
		time.Sleep(time.Second)
	}
}

/*

*/