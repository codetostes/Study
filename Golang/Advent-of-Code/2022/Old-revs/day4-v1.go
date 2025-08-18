package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("files/input_04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1Total, part2Total := processJobsAssignments(file)

	fmt.Println("Part 1:", part1Total)
	fmt.Println("Part 2:", part2Total)

}

func processJobsAssignments(file *os.File) (part1Total, part2Total int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Split the line into two parts
		pair := strings.Split(line, ",")
		if len(pair) != 2 {
			continue // Skip invalid lines
		}

		part1Total += checkFullOverlap(pair[0], pair[1])
		part2Total += checkPartialOverlap(pair[0], pair[1])
	}

	return part1Total, part2Total
}

func checkFullOverlap(job1, job2 string) int {
	start1, end1 := parseJobRange(job1)
	start2, end2 := parseJobRange(job2)

	if (start1 <= start2 && end1 >= end2) || (start2 <= start1 && end2 >= end1) {
		return 1 // Full overlap
	}
	return 0 // No full overlap
}

func checkPartialOverlap(job1, job2 string) int {
	start1, end1 := parseJobRange(job1)
	start2, end2 := parseJobRange(job2)

	if (start1 <= end2 && end1 >= start2) || (start2 <= end1 && end2 >= start1) {
		return 1 // Partial overlap
	}
	return 0 // No partial overlap
}

func parseJobRange(job string) (int, int) {
	parts := strings.Split(job, "-")
	if len(parts) != 2 {
		return 0, 0 // Invalid job range
	}

	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return 0, 0 // Invalid job range
	}
	return start, end
}
