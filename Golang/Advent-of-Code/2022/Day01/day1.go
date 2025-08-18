package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// CalorieCounter represents a solution for Advent of Code Day 1 - Calorie Counting
func main() {
	file := "files/input.txt"

	calories, err := ParseElfCalories(file)
	if err != nil {
		log.Fatal("Failed to parse calorie data:", err)
	}

	fmt.Println("Part 1 =", FindMaxCalories(calories))
	fmt.Println("Part 2 =", FindTop3CaloriesSum(calories))

	// Original code for comparison:
	/*
		file := "input.txt"
		lines := readFile(file)

		fmt.Println("Parte 1 =", mostCalories(lines))
		fmt.Println("Parte 2 =", top3Calories(lines))
	*/
}

// ParseElfCalories reads the input file and returns a slice of total calories per elf
func ParseElfCalories(filename string) ([]int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var elfCalories []int
	var currentTotal int

	for lineNum, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			// Empty line indicates end of current elf's inventory
			elfCalories = append(elfCalories, currentTotal)
			currentTotal = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("invalid calorie value at line %d: %s", lineNum+1, line)
			}
			if calories < 0 {
				return nil, fmt.Errorf("negative calorie value not allowed at line %d: %d", lineNum+1, calories)
			}
			currentTotal += calories
		}
	}

	// Don't forget the last elf if file doesn't end with empty line
	if currentTotal > 0 {
		elfCalories = append(elfCalories, currentTotal)
	}

	return elfCalories, nil

	/*
		func readFile(file string) []string {
			data, erro := os.ReadFile(file)
			if erro != nil {
				log.Fatal(erro)
			}

		lines := strings.Split(string(data), "\n")
		return lines
		}
	*/
}

// FindMaxCalories returns the maximum calories carried by any single elf
func FindMaxCalories(elfCalories []int) int {
	if len(elfCalories) == 0 {
		return 0
	}

	maxCalories := elfCalories[0]
	for _, calories := range elfCalories[1:] {
		if calories > maxCalories {
			maxCalories = calories
		}
	}

	return maxCalories

	// Original function for comparison:
	/*
		func mostCalories(lines []string) int {
			total, mostCal := 0, 0
			for _, line := range lines {
				if line != "" {
					cal, _ := strconv.Atoi(line)
					total += cal
				} else if total > mostCal {
					mostCal = total
					total = 0
				} else {
					total = 0
				}
			}
			return mostCal
		}
	*/
}

// FindTop3CaloriesSum returns the sum of calories carried by the top 3 elves
func FindTop3CaloriesSum(elfCalories []int) (sum int) {
	if len(elfCalories) == 0 {
		return 0
	}

	// Create a copy to avoid modifying the original slice
	caloriesCopy := make([]int, len(elfCalories))
	copy(caloriesCopy, elfCalories)

	// Sort in descending order
	sort.Ints(caloriesCopy)
	slices.Reverse(caloriesCopy)

	// Sum the top 3 (or fewer if less than 3 elves)
	top3Count := 3
	if len(caloriesCopy) < top3Count {
		top3Count = len(caloriesCopy)
	}

	for i := 0; i < top3Count; i++ {
		sum += caloriesCopy[i]
	}

	return

	// Original function for comparison:
	/*
		func top3Calories(lines []string) int {
			var (
				total, elf1, elf2, elf3, i int
			)
			for _, line := range lines {
				if line != "" {
					cal, _ := strconv.Atoi(line)
					total += cal
					i++

				} else if total > 60000 {

					switch {
					case total > elf1:
						elf2, elf1 = elf1, total
						total = 0
					case total > elf2:
						elf3, elf2 = elf2, total
						total = 0
					case total > elf3:
						elf3 = total
						total = 0
					default:
						total = 0
					}
				} else {
					total = 0
				}
			}
			total = (elf1 + elf2 + elf3)
			return total
		}
	*/
}

// Alternative implementation using a min-heap approach for better efficiency with larger datasets
func FindTop3CaloriesSumOptimized(elfCalories []int) (sum int) {
	if len(elfCalories) == 0 {
		return 0
	}

	// Keep track of top 3 using a simple array
	top3 := make([]int, 0, 3)

	for _, calories := range elfCalories {
		if len(top3) < 3 {
			top3 = append(top3, calories)
			sort.Ints(top3) // Keep sorted (smallest first)
		} else if calories > top3[0] { // If larger than smallest in top3
			top3[0] = calories
			sort.Ints(top3)
		}
	}

	for _, calories := range top3 {
		sum += calories
	}

	return
}
