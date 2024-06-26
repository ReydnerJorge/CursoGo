package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(2)
	go func() {
		escrever("Olá Mundo!")
		waitgroup.Done()
	}()

	go func() {
		escrever("Programando em go")
		waitgroup.Done()
	}()

	waitgroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 10; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

/*
waitgroup e uma forma de sincronizar as goroutines
*/