package main

import "fmt"

// Go não é OO então não tem conceitos de classe e polimorfismo
// Structs em go é o que seria mais próximo de de uma classe em uma linguágem OO
// Stricts são utilizados como uma colecao de campos. Cada campo tem um nome e um tipo

type User struct {
	name    string
	age     uint8
	address Address // um struct pode ser referenciado dentro de outro usando o mesmo nome

}
type Address struct {
	street string
	number uint16
}

func main() {
	var u User // uma variável não pode ser declarada usando mesmo nome de uma struct
	u.name = "Carlos Tostes"
	u.age = 31
	fmt.Println(u)

	enderecoExemplo := Address{"Rua dos sonhadores", 1003}

	u2 := User{"Victoria Tostes", 24, enderecoExemplo}
	fmt.Println(u2)

	u3 := User{age: 21}
	fmt.Println(u3)
}
