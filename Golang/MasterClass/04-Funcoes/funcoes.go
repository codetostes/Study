package main

import (
	"fmt"
)

// declaração de funçao.
func somar(n1, n2 int) int {
	return n1 + n2
}

// retornar 2 valores é um dos diferenciais do Go.
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subs := n1 - n2
	return soma, subs
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// declaração de função e atribuindo a uma variável.
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado1 := f("Texto da função 1")
	fmt.Println(resultado1)

	resultadoSoma, ResultadoSubs := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, ResultadoSubs)

	// para ignorar um retorno caso a função tenha mais de um
	// basta colocar _ (Underline) no lugar da variável atribuida.
	resultado2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultado2)

	_, resultado3 := calculosMatematicos(10, 15)
	fmt.Println(resultado3)
}
