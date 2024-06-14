package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "Reydner",
		"sobrenome": "Jorge",
		"idade": "28",
		"endereço": "Rua Tanzânia 44",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	//alinhamento de maps
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Reydner",
			"ultimo": "Jorge",
		},
		"curso": {
			"nome": "Ads",
			"faculdade": "Una",
			"turma": "1C",
			"turno": "noturno",
		},
	}
	fmt.Println(usuario2)
}