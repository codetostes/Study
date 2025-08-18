package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Advent of Code 2022 - Day 3
// Part 1: Calculate the sum of the values of the duplicate items in each rucksack
func main() {
	file, err := os.Open("files/input_3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	calcPart1, calcPart2 := 0, 0
	i := 1
	var rucksackGroup []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line)%2 != 0 {
			log.Fatal("Wrong line format. Line cannot have odd index")

		}

		calcPart1 += findRucksackDuplicate(line)

		if i%3 != 0 {
			rucksackGroup = append(rucksackGroup, line)
			i++
			continue
		} else {
			rucksackGroup = append(rucksackGroup, line)
			calcPart2 += findRucksackGroup(rucksackGroup)
			rucksackGroup = nil // Reset for the next group
			i = 1
		}

	}
	fmt.Printf("\n\n----------\n")
	fmt.Println(calcPart1)
	fmt.Println(calcPart2)
}

func findRucksackDuplicate(rucksack string) int {
	for i := 0; i < len(rucksack)/2; i++ {
		item1 := rucksack[i]
		for j := len(rucksack) / 2; j < len(rucksack); j++ {
			item2 := rucksack[j]
			if item1 != item2 {
				continue
			} else {
				return calcitemPriority(item1)
			}
		}
	}
	return 0
}

func findRucksackGroup(group []string) int {
	item1, item2, item3 := byte(0), byte(0), byte(0)
	for i := 0; i < len(group[0]); i++ {
		item1 = group[0][i]
		for j := 0; j < len(group[1]); j++ {
			item2 = group[1][j]
			if item1 == item2 {

				for k := 0; k < len(group[2]); k++ {
					item3 := group[2][k]
					if item1 == item3 {
						fmt.Println("Found item:", string(item1), "in group:", group)
						return calcitemPriority(item1)
					}
				}
			} else if item1 != item2 {
				continue
			} else if item1 == item2 && item1 != item3 {
				continue
			} else {
				return calcitemPriority(item1)
			}
		}
	}
	return 0
}

func calcitemPriority(item byte) int {
	if item >= 'A' && item <= 'Z' {
		return int(item - ('A' - 27))
	} else if item >= 'a' && item <= 'z' {
		return int(item - ('a' - 1))
	}
	return 0
}
