package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Game constants - much clearer than magic numbers
const (
	// Shape scores
	ScoreRock = iota + 1
	ScorePaper
	ScoreScissors

	// Outcome scores
	ScoreLose = 0
	ScoreDraw = 3
	ScoreWin  = 6
)

// Shape represents a Rock Paper Scissors choice
type Shape int

const (
	Rock Shape = iota
	Paper
	Scissors
)

// Outcome represents the result of a round
type Outcome int

const (
	Lose Outcome = iota
	Draw
	Win
)

// Round represents a single game round
type Round struct {
	OpponentShape Shape
	PlayerAction  string // Can be Shape or desired Outcome depending on interpretation
}

func main() {
	rounds, err := parseStrategyGuide("files/input_2.txt")
	if err != nil {
		log.Fatal("Failed to parse strategy guide:", err)
	}

	part1Score := calculateTotalScore(rounds, interpretAsPlayerShape)
	part2Score := calculateTotalScore(rounds, interpretAsDesiredOutcome)

	fmt.Println("Part 1 Score:", part1Score)
	fmt.Println("Part 2 Score:", part2Score)

	// Original code for comparison:
	/*
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
	*/
}

// parseStrategyGuide reads the strategy guide file and returns rounds
func parseStrategyGuide(filename string) ([]Round, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	var rounds []Round
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue // Skip empty lines
		}

		if len(line) != 3 || line[1] != ' ' {
			return nil, fmt.Errorf("invalid format at line %d: expected 'A X' format, got '%s'", lineNum, line)
		}

		opponentShape, err := parseOpponentShape(line[0])
		if err != nil {
			return nil, fmt.Errorf("invalid opponent shape at line %d: %c", lineNum, line[0])
		}

		round := Round{
			OpponentShape: opponentShape,
			PlayerAction:  string(line[2]),
		}

		rounds = append(rounds, round)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filename, err)
	}

	return rounds, nil
}

// parseOpponentShape converts opponent's letter to Shape
func parseOpponentShape(letter byte) (Shape, error) {
	switch letter {
	case 'A':
		return Rock, nil
	case 'B':
		return Paper, nil
	case 'C':
		return Scissors, nil
	default:
		return 0, fmt.Errorf("unknown opponent shape: %c", letter)
	}
}

// calculateTotalScore calculates total score using the provided interpretation function
func calculateTotalScore(rounds []Round, interpreter func(Round) int) int {
	totalScore := 0
	for _, round := range rounds {
		totalScore += interpreter(round)
	}
	return totalScore
}

// interpretAsPlayerShape treats X/Y/Z as player's shape choice (Part 1)
func interpretAsPlayerShape(round Round) int {
	playerShape, err := parsePlayerShape(round.PlayerAction)
	if err != nil {
		return 0 // Invalid input, score 0
	}

	shapeScore := getShapeScore(playerShape)
	outcomeScore := getOutcomeScore(round.OpponentShape, playerShape)

	return shapeScore + outcomeScore
}

// interpretAsDesiredOutcome treats X/Y/Z as desired outcome (Part 2)
func interpretAsDesiredOutcome(round Round) int {
	desiredOutcome, err := parseDesiredOutcome(round.PlayerAction)
	if err != nil {
		return 0 // Invalid input, score 0
	}

	playerShape := getShapeForOutcome(round.OpponentShape, desiredOutcome)
	shapeScore := getShapeScore(playerShape)
	outcomeScore := getOutcomeScore(round.OpponentShape, playerShape)

	return shapeScore + outcomeScore
}

// parsePlayerShape converts player's letter to Shape (Part 1 interpretation)
func parsePlayerShape(action string) (Shape, error) {
	switch action {
	case "X":
		return Rock, nil
	case "Y":
		return Paper, nil
	case "Z":
		return Scissors, nil
	default:
		return 0, fmt.Errorf("unknown player shape: %s", action)
	}
}

// parseDesiredOutcome converts player's letter to desired Outcome (Part 2 interpretation)
func parseDesiredOutcome(action string) (Outcome, error) {
	switch action {
	case "X":
		return Lose, nil
	case "Y":
		return Draw, nil
	case "Z":
		return Win, nil
	default:
		return 0, fmt.Errorf("unknown desired outcome: %s", action)
	}
}

// getShapeScore returns the score for playing a specific shape
func getShapeScore(shape Shape) int {
	switch shape {
	case Rock:
		return ScoreRock
	case Paper:
		return ScorePaper
	case Scissors:
		return ScoreScissors
	default:
		return 0
	}
}

// getOutcomeScore determines the outcome score based on opponent and player shapes
func getOutcomeScore(opponent, player Shape) int {
	if opponent == player {
		return ScoreDraw
	}

	// Check winning conditions
	switch {
	case player == Rock && opponent == Scissors,
		player == Paper && opponent == Rock,
		player == Scissors && opponent == Paper:
		return ScoreWin
	default:
		return ScoreLose
	}
}

// getShapeForOutcome determines what shape to play to achieve desired outcome
func getShapeForOutcome(opponent Shape, desired Outcome) Shape {
	switch desired {
	case Draw:
		return opponent // Same shape for draw
	case Win:
		switch opponent {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	case Lose:
		switch opponent {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	}
	return Rock // Default fallback
}

// Original functions preserved for comparison:
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
