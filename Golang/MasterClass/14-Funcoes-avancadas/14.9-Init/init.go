package main

import "fmt"

func main() {
	fmt.Println("Função main sendo executada")
}

// função INIT é executada sempre antes da função MAIN
// independente de ter sido declarada antes ou após a função MAIN
func init() {
	fmt.Println("Executando a função init")
}
