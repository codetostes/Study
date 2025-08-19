package main

import "fmt"

func manyMathCalcs(numeros ...int) (total int) {
	for _, numero := range numeros {
		total += numero
	}
	return
}

func variaticFunction(nome string, numeros ...int) {

	total := 0
	for _, numero := range numeros {
		total += numero
	}

	fmt.Println(nome, total)
}

func main() {
	sumTotal := manyMathCalcs(1, 2, 3, 4, 5, 6, 7, 101, 204, 309)
	fmt.Println(sumTotal)

	variaticFunction("Tostes", 1, 2, 3, 4, 5)
}
