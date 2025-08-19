package main

import "fmt"

func main() {
	// ===== ARITMÉTICOS =====
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// Não é possível fazer operações aritméticas com tipos diferentes de dados
	// var numero1 int16 = 10
	// var numero2 int32 = 25
	// soma2 := numero1 + numero2
	// fmt.Println(soma2)
	// ===== FIM ARITMETICOS =====

	//===== ATRIBUIÇÃO =====
	var variavel1 string = "String"
	variavel2 := "String"
	fmt.Println(variavel1, variavel2)
	// ===== FIM ATRIBUIÇÃO =====

	// ===== OPERADORES RELACIONAIS =====
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	// ===== FIM OPERADORES RELACIONAIS =====

	// ==== OPERADORES LOGICOS =====
	fmt.Println("===== LOGICOS=====")
	verdadeiro, falso := true, false
	fmt.Println(falso && false)      // AND
	fmt.Println(falso || verdadeiro) // OR
	fmt.Println(!true)               // NEGATE
	fmt.Println(!false)
	// ===== FIM OPERADORES LOGICOS =====

	// ===== OPERADORES UNARIOS =====
	numero := 10
	numero++     // SOMA 1
	numero += 15 // SOMA 15
	numero--     // SUBTRAI 1
	numero -= 5  // SUBTRAI 5
	numero *= 3  // MULTIPLICA POR 3
	numero /= 10 // DIVIDE POR 10
	numero %= 3  // RESTO DA DIVISAO POR 3

	fmt.Println(numero)
	//===== FIM OPERADORES UNARIOS =====

	// ===== OPERADOR TERNARIO =====

	// Go não utiliza operadores ternários. comparação de condições apenas com if..else

	// Exemplo operador ternário:
	// texto := numero > 5 ? "Maior que 5" : "Menor que 5"

	// ===== FIM OPERADOR TERNARIO =====s
}
