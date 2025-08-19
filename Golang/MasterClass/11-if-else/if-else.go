package main

import "fmt"

func main() {

	numero := -3

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// variáveis criadas usando IF INIT são validas apenas para o escopo da estrutura
	if numero2 := numero; numero2 > 0 {
		fmt.Println("Número é maior que 0")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número está entre 0 e -10")
	}
}
