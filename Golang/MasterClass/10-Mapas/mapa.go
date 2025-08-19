package main

import "fmt"

func main() {
	// o map é mais rigido que o sctruct. Só aceita 1 tipo de dado por vez, seja na chave ou valor.
	usuario := map[string]string{
		"nome":      "Carlos",
		"sobrenome": "Tostes",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Victoria",
			"ultimo":   "Tostes",
		},
		"Curso": {
			"nome":   "Pedagogia",
			"campus": "UFPB",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}
