package main

import (
	"fmt"
	"time"
)

func main() {
	 fmt.Println("Loops")
	 	 i := 0

	 for i < 10 {
 	 	i++
	  		fmt.Println("Incrementando i")
	 	 	time.Sleep(time.Second)
	 	 }
	 	 fmt.Println(i)

	 	 for j := 0; j < 10; j++ {
	 	 	fmt.Println("Incrementando j", j)
	 	 	time.Sleep(time.Second)
		 }


		//range 
		nomes := [3]string{"Reydner", "Juan", "Rafael"}

		for indices, nome := range nomes {
			fmt.Println(indices, nome)
		}

		for indice, letra := range "Reydner" {
			fmt.Println(indice, string(letra))
		}

		usuario := map[string]string{
			"nome": "Reydner",
			"sobrenome": "Oliveira",
		}
		for chave, valor := range usuario {
			fmt.Println(chave, valor)
		}
}