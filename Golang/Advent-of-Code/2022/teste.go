package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := `31-31,32-40
26-92,13-91
9-90,29-91
72-72,25-73`

	// Método 1: Processando linha por linha
	fmt.Println("=== Método 1: Linha por linha ===")
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		fmt.Printf("Linha %d: %s\n", i+1, line)

		// Separa os pares por vírgula
		pairs := strings.Split(line, ",")

		for j, pair := range pairs {
			// Separa cada par por hífen
			numbers := strings.Split(pair, "-")

			if len(numbers) == 2 {
				// Converte para inteiros
				num1, err1 := strconv.Atoi(numbers[0])
				num2, err2 := strconv.Atoi(numbers[1])

				if err1 == nil && err2 == nil {
					fmt.Printf("  Par %d: %d e %d\n", j+1, num1, num2)
				}
			}
		}
		fmt.Println()
	}
}
