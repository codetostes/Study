package main

import "fmt"

type Person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type Student struct {
	Person
	course string
	degree string
}

// herança em GO funciona de forma a não precisar especificar o nome para a estrutura dentro de outra.
// dessa forma o acesso aos dados será "estudante.nome" ao invés de "estudante.pessoa.nome"

func main() {
	p1 := person1()
	student1(p1)
}

func person1() Person {
	p1 := Person{"Carlos", "Tostes", 32, 183}
	return p1
}

func student1(p1 Person) {
	e1 := Student{p1, "Computer Science", "Bachelor"}
	fmt.Println(e1.name, e1.surname)
}
