package main

import "fmt"

// PARAMETRO POR CÓPIA
func inverterSinal(num int) int {
	return num * -1
}

// PAREMETRO POR REFERÊNCIA
func interverSinalPonteiro(num *int) {
	*num = *num * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	fmt.Println("----------")

	novoNumero := 40
	fmt.Println(novoNumero)
	interverSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
