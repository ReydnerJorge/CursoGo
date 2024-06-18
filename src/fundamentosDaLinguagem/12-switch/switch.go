package main

import "fmt"

// primeira forma de se fazer um switch
func diaDaDSemana(numero int) string {
	switch numero {
	case 1:
			return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5: 
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7: 
		return "Sabado"
	default:
		return "Número Inválido"
}
}
// segunda forma de criar um switch
func diaDaSemana2 (numero int) string {
	var diaDaSemana3 string

	switch {
	case numero == 1:
		diaDaSemana3 =  "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana3 =  "Segunda"
	case numero == 3:
		diaDaSemana3 = "Terça"
	case numero == 4:
		diaDaSemana3 = "Quarta"
	case numero == 5:
		diaDaSemana3 = "Quinta"
	case numero == 6:
		diaDaSemana3 = "Sextou"
	case numero == 7:
		diaDaSemana3 = "Sabado"
	default:
		return "Número Inválido"
	}

	return diaDaSemana3
}

func main() {
	fmt.Println("Switch")

	fmt.Println("-----------------")

	dia := diaDaDSemana(7)
	fmt.Println(dia)

	//fmt.Println("-----------------")

	//dia2 := diaDaSemana2(6)
	//fmt.Println(dia2)
	
	fmt.Println("------------------")
	dia3 := diaDaSemana2(1)
	fmt.Println(dia3)
}