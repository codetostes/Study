package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Advent of Code 2022 - Day 3
func main() {
	file, err := os.Open("files/input_3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1Total, part2Total := processRucksacks(file)

	fmt.Printf("Part 1: %d\n", part1Total)
	fmt.Printf("Part 2: %d\n", part2Total)
}

func processRucksacks(file *os.File) (part1Total, part2Total int) {
	scanner := bufio.NewScanner(file)

	groupLines := make([]string, 0, 3)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Part 1: Find duplicate in each rucksack
		part1Total += findDuplicateInRucksack(line)

		// Part 2: Collect lines for group processing
		groupLines = append(groupLines, line)

		if len(groupLines) == 3 {
			part2Total += findCommonItemInGroup(groupLines)
			groupLines = groupLines[:0] // Reset slice
		}
	}
	return
}

func findDuplicateInRucksack(rucksack string) int {
	if len(rucksack)%2 != 0 {
		return 0 // Skip invalid lines instead of crashing
	}

	mid := len(rucksack) / 2
	firstHalf := rucksack[:mid]
	secondHalf := rucksack[mid:]

	// Create set for first half using struct{} for zero memory overhead
	firstHalfItems := make(map[byte]struct{}, len(firstHalf))
	for i := 0; i < len(firstHalf); i++ {
		firstHalfItems[firstHalf[i]] = struct{}{}
	}

	// Find first common item in second half
	for i := 0; i < len(secondHalf); i++ {
		if _, exists := firstHalfItems[secondHalf[i]]; exists {
			return calculatePriority(secondHalf[i])
		}
	}

	return 0
}

func findCommonItemInGroup(group []string) int {
	if len(group) != 3 {
		return 0
	}

	// Create sets for first two rucksacks using struct{} for memory efficiency
	firstSet := createItemSet(group[0])
	secondSet := createItemSet(group[1])

	// Find common item in all three
	for i := 0; i < len(group[2]); i++ {
		item := group[2][i]
		if _, existsInFirst := firstSet[item]; existsInFirst {
			if _, existsInSecond := secondSet[item]; existsInSecond {
				return calculatePriority(item)
			}
		}
	}

	return 0
}

func createItemSet(rucksack string) map[byte]struct{} {
	itemSet := make(map[byte]struct{}, len(rucksack))
	for i := 0; i < len(rucksack); i++ {
		itemSet[rucksack[i]] = struct{}{}
	}
	return itemSet
}

func calculatePriority(item byte) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	}
	if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27)
	}
	return 0
}
