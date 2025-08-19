package main

import (
	"fmt"
)

func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println("Incrementando i ", i)
	// 	time.Sleep(time.Second)
	// 	i++
	// }

	// fmt.Println("---------")

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	fmt.Println("---------")

	nomes := [3]string{"Carlos", "Cris", "Alexandre"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Carlos",
		"sobrenome": "Tostes",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
