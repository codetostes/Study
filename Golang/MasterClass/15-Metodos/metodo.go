package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiordeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Carlos Tostes", 31}
	usuario1.salvar()

	usuario2 := usuario{"Victoria Tostes", 25}

	usuario2.salvar()
	fmt.Println(usuario2.maiordeIdade())

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}

// Métodos são bastente parecidos com funções, porém estão obrigatoriamente associado a algo.
// seja um struct, uma interface, etc.
