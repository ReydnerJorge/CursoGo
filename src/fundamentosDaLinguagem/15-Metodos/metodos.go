package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário(a) %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) aniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Juan", 23}
	usuario1.salvar()

	usuario2 := usuario{"Reydner", 17}
	usuario2.salvar()

	usuario3 := usuario{"Rafael", 27}
	usuario3.salvar()

	usuario4 := usuario{"Nathália", 19}
	usuario4.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println("Usuário e maior de idade: ", maiorDeIdade)

	usuario3.aniversario()
	fmt.Println(usuario3.idade)
}