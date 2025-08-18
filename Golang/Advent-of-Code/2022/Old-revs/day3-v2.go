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

func processRucksacks(file *os.File) (int, int) {
	scanner := bufio.NewScanner(file)

	part1Total := 0
	part2Total := 0
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

	return part1Total, part2Total
}

func findDuplicateInRucksack(rucksack string) int {
	if len(rucksack)%2 != 0 {
		return 0 // Skip invalid lines instead of crashing
	}

	mid := len(rucksack) / 2
	firstHalf := rucksack[:mid]
	secondHalf := rucksack[mid:]

	// Create set for first half
	firstHalfItems := make(map[byte]bool)
	for i := 0; i < len(firstHalf); i++ {
		firstHalfItems[firstHalf[i]] = true
	}

	// Find first common item in second half
	for i := 0; i < len(secondHalf); i++ {
		if firstHalfItems[secondHalf[i]] {
			return calculatePriority(secondHalf[i])
		}
	}

	return 0
}

func findCommonItemInGroup(group []string) int {
	if len(group) != 3 {
		return 0
	}

	// Create sets for first two rucksacks
	firstSet := createItemSet(group[0])
	secondSet := createItemSet(group[1])

	// Find common item in all three
	for i := 0; i < len(group[2]); i++ {
		item := group[2][i]
		if firstSet[item] && secondSet[item] {
			return calculatePriority(item)
		}
	}

	return 0
}

func createItemSet(rucksack string) map[byte]bool {
	itemSet := make(map[byte]bool)
	for i := 0; i < len(rucksack); i++ {
		itemSet[rucksack[i]] = true
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
