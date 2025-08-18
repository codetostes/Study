package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// CalorieCounter represents a solution for Advent of Code Day 1 - Calorie Counting
type CalorieCounter struct {
	filename string
}

// NewCalorieCounter creates a new CalorieCounter instance
func NewCalorieCounter(filename string) *CalorieCounter {
	return &CalorieCounter{filename: filename}
}

func main() {
	counter := NewCalorieCounter("files/input.txt")

	calories, err := counter.ParseElfCalories()
	if err != nil {
		log.Fatal("Failed to parse calorie data:", err)
	}

	fmt.Println("Part 1 =", counter.FindMaxCalories(calories))
	fmt.Println("Part 2 =", counter.FindTop3CaloriesSum(calories))
}

// ParseElfCalories reads the input file and returns a slice of total calories per elf
func (cc *CalorieCounter) ParseElfCalories() ([]int, error) {
	data, err := os.ReadFile(cc.filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", cc.filename, err)
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
}

// FindMaxCalories returns the maximum calories carried by any single elf
func (cc *CalorieCounter) FindMaxCalories(elfCalories []int) int {
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
}

// FindTop3CaloriesSum returns the sum of calories carried by the top 3 elves
func (cc *CalorieCounter) FindTop3CaloriesSum(elfCalories []int) int {
	if len(elfCalories) == 0 {
		return 0
	}

	// Create a copy to avoid modifying the original slice
	caloriesCopy := make([]int, len(elfCalories))
	copy(caloriesCopy, elfCalories)

	// Sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesCopy)))

	// Sum the top 3 (or fewer if less than 3 elves)
	sum := 0
	top3Count := 3
	if len(caloriesCopy) < top3Count {
		top3Count = len(caloriesCopy)
	}

	for i := 0; i < top3Count; i++ {
		sum += caloriesCopy[i]
	}

	return sum
}

// Alternative implementation using a min-heap approach for better efficiency with larger datasets
func (cc *CalorieCounter) FindTop3CaloriesSumOptimized(elfCalories []int) int {
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

	sum := 0
	for _, calories := range top3 {
		sum += calories
	}

	return sum
}
