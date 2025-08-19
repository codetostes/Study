package main

import (
	"errors"
	"fmt"
)

func main() {
	// ===== NÚMEROS INTEIROS =====
	// int8, int16, int32, int64 varieades do tipo INT. Altera quantidade de bits
	// ao usar apenas 'int', será utilizado a arquitura do processador como base.
	var numero int16 = 32767
	fmt.Println(numero)

	// uint = unsigned int. Número inteiro sem sinal (aceita apenas números positivos)
	// assim como o int possui as variedade de 8,16,32,64 bits.
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// ALLIASES
	// RUNE = INT32
	var numero3 rune = 65
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 100
	fmt.Println(numero4)

	// ===== FIM NÚMEROS INTEIROS =====

	// ===== NÚMEROS REAIS =====
	// float32, float64.
	var numeroReal1 float32 = 60.7
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1120731.43
	fmt.Println(numeroReal2)

	// ===== FIM NÚMEROS REAIS =====

	// ===== STRINGS =====
	// string sempre utilizar aspas duplas,
	var nome string = "Carlos Tostes"
	fmt.Println(nome)

	char := 'B'
	fmt.Println(char)

	// ===== FIM STRINGS =====

	// ===== BOOLEANO e ERROR =====
	// true or false
	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)

	// ==== VALOR (ZERO) =====
	// todo variável em GO tem um valor default.
	// 0 para int; "" para string; nulo (<nil>) para error; false para bool
	var texto0 string
	fmt.Println(texto0)

	var numero0 int
	fmt.Println(numero0)

	var float0 float64
	fmt.Println(float0)

	var erro0 error
	fmt.Println(erro0)

	var booleano0 bool
	fmt.Println(booleano0)

	// ===== FIM VALOR (ZERO) =====

}
