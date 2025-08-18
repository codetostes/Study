package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("files/input_2.txt")
	if err != nil {
		log.Fatal(file, err)
	}
	defer file.Close()

	soma1, soma2 := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jokenpoPlay(scanner.Text(), &soma1)
		jokenpoPlayCorrect(scanner.Text(), &soma2)
	}
	fmt.Println(soma1)
	fmt.Println(soma2)
}

// Calculates results from elf jokenpo strategy guide.
func jokenpoPlay(play string, soma *int) {

	switch play[0] {
	case 'A': // PEDRA
		switch play[2] {
		case 'X': // PEDRA
			*soma = *soma + (1 + 3)
		case 'Y': // PAPEL
			*soma = *soma + (2 + 6)
		case 'Z': //TESOURA
			*soma = *soma + (3 + 0)
		}
	case 'B': // PAPEL
		switch play[2] {
		case 'X': // PEDRA
			*soma = *soma + (1 + 0)
		case 'Y': // PAPEL
			*soma = *soma + (2 + 3)
		case 'Z': //TESOURA
			*soma = *soma + (3 + 6)
		}
	case 'C': //TESOURA
		switch play[2] {
		case 'X': // PEDRA
			*soma = *soma + (1 + 6)
		case 'Y': // PAPEL
			*soma = *soma + (2 + 0)
		case 'Z': //TESOURA
			*soma = *soma + (3 + 3)
		}

	}
}

// Calculates results from elf jokenpo strategy guide.
func jokenpoPlayCorrect(play string, soma *int) {

	switch play[0] {
	case 'A': // PEDRA
		switch play[2] {
		case 'X': // LOSE
			*soma = *soma + (3 + 0)
		case 'Y': // DRAW
			*soma = *soma + (1 + 3)
		case 'Z': // WIN
			*soma = *soma + (2 + 6)
		}
	case 'B': // PAPEL
		switch play[2] {
		case 'X': // LOSE
			*soma = *soma + (1 + 0)
		case 'Y': // DRAW
			*soma = *soma + (2 + 3)
		case 'Z': // WIN
			*soma = *soma + (3 + 6)
		}
	case 'C': //TESOURA
		switch play[2] {
		case 'X': // LOSE
			*soma = *soma + (2 + 0)
		case 'Y': // DRAW
			*soma = *soma + (3 + 3)
		case 'Z': // WIN
			*soma = *soma + (1 + 6)
		}

	}
}
