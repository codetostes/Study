package main

import "fmt"

func main() {

	variavel1 := 10
	variavel2 := variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERÊNCIA PARA UM ENDEREÇO DE MEMÓRIA
	var variavel3 int16 = 100
	ponteiro1 := &variavel3

	// ao mandar imprimir o ponteiro, ele vai imprimir o endereço de memória
	fmt.Println(variavel3, ponteiro1)

	// para pegar o valor que está sendo apontado pelo ponteiro se usa a DESREFERENCIAÇÃO.
	fmt.Println(variavel3, *ponteiro1)
	variavel3 += 50
	fmt.Println(variavel3, *ponteiro1)
	fmt.Println(variavel3, ponteiro1)
}
