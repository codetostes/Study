package main

import (
	"fmt"
)

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	// o DEFER adia para o último momento porém sempre antes do RETURN.
	defer fmt.Println("Média calculada. Resultado será retornado!")

	fmt.Println("Função para verificação de aprovação do Aluno")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	} else {
		return false
	}
}

func main() {
	// defer funcao1()
	// // DEFER = ADIAR. Adia a função até o último momento possível.
	// funcao2()

	fmt.Println(alunoAprovado(7, 8))
}
