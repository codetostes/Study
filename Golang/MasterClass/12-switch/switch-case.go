package main

import "fmt"

func diaDaSemana(num uint8) string {
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(num uint8) string {
	var diaDaSemana string

	switch {
	case num == 1:
		diaDaSemana = "Domingo"
		fallthrough // faz com que a próx condição seja executada indepentende de ser verdadeira
	case num == 2:
		diaDaSemana = "Segunda-Feira"
	case num == 3:
		diaDaSemana = "Terça-Feira"
	case num == 4:
		diaDaSemana = "Quarta-Feira"
	case num == 5:
		diaDaSemana = "Quinta-Feira"
	case num == 6:
		diaDaSemana = "Sexta-Feira"
	case num == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println(diaDaSemana(1))

	fmt.Println(diaDaSemana2(4))
}
