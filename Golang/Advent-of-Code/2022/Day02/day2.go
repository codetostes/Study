package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Score constants - clearer than magic numbers
const (
	ScoreRock     = 1
	ScorePaper    = 2
	ScoreScissors = 3
	ScoreLose     = 0
	ScoreDraw     = 3
	ScoreWin      = 6
)

func main() {
	file, err := os.Open("files/input_2.txt")
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	part1Total := 0
	part2Total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) != 3 {
			continue // Skip invalid lines
		}

		part1Total += calculatePart1Score(line[0], line[2])
		part2Total += calculatePart2Score(line[0], line[2])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}

	fmt.Println("Part 1:", part1Total)
	fmt.Println("Part 2:", part2Total)

	// Original code for comparison:
	/*
		soma1, soma2 := 0, 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jokenpoPlay(scanner.Text(), &soma1)
			jokenpoPlayCorrect(scanner.Text(), &soma2)
		}
		fmt.Println(soma1)
		fmt.Println(soma2)
	*/
}

// calculatePart1Score treats X/Y/Z as player shapes
func calculatePart1Score(opponent, player byte) int {
	// Opponent: A=Rock, B=Paper, C=Scissors
	// Player: X=Rock, Y=Paper, Z=Scissors

	var shapeScore int
	switch player {
	case 'X':
		shapeScore = ScoreRock
	case 'Y':
		shapeScore = ScorePaper
	case 'Z':
		shapeScore = ScoreScissors
	}

	var outcomeScore int
	switch {
	// Draw cases
	case (opponent == 'A' && player == 'X') || // Rock vs Rock
		(opponent == 'B' && player == 'Y') || // Paper vs Paper
		(opponent == 'C' && player == 'Z'): // Scissors vs Scissors
		outcomeScore = ScoreDraw

	// Win cases
	case (opponent == 'A' && player == 'Y') || // Rock vs Paper
		(opponent == 'B' && player == 'Z') || // Paper vs Scissors
		(opponent == 'C' && player == 'X'): // Scissors vs Rock
		outcomeScore = ScoreWin

	// Lose cases (everything else)
	default:
		outcomeScore = ScoreLose
	}

	return shapeScore + outcomeScore
}

// calculatePart2Score treats X/Y/Z as desired outcomes
func calculatePart2Score(opponent, outcome byte) int {
	// Opponent: A=Rock, B=Paper, C=Scissors
	// Outcome: X=Lose, Y=Draw, Z=Win

	var outcomeScore int
	switch outcome {
	case 'X':
		outcomeScore = ScoreLose
	case 'Y':
		outcomeScore = ScoreDraw
	case 'Z':
		outcomeScore = ScoreWin
	}

	// Determine what shape to play
	var shapeScore int
	switch outcome {
	case 'Y': // Draw - play same shape
		switch opponent {
		case 'A':
			shapeScore = ScoreRock
		case 'B':
			shapeScore = ScorePaper
		case 'C':
			shapeScore = ScoreScissors
		}

	case 'Z': // Win - play winning shape
		switch opponent {
		case 'A': // Rock -> play Paper
			shapeScore = ScorePaper
		case 'B': // Paper -> play Scissors
			shapeScore = ScoreScissors
		case 'C': // Scissors -> play Rock
			shapeScore = ScoreRock
		}

	case 'X': // Lose - play losing shape
		switch opponent {
		case 'A': // Rock -> play Scissors
			shapeScore = ScoreScissors
		case 'B': // Paper -> play Rock
			shapeScore = ScoreRock
		case 'C': // Scissors -> play Paper
			shapeScore = ScorePaper
		}
	}

	return shapeScore + outcomeScore
}

// Original functions for comparison:
/*
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
*/
