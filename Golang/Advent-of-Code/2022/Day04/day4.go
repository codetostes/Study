package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Range representa um intervalo numérico
type Range struct {
	Start, End int
}

// Contains verifica se este range contém completamente o outro
func (r Range) Contains(other Range) bool {
	return r.Start <= other.Start && r.End >= other.End
}

// Overlaps verifica se há sobreposição entre os ranges
func (r Range) Overlaps(other Range) bool {
	return r.Start <= other.End && r.End >= other.Start
}

func main() {
	file, err := os.Open("files/input_04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fullOverlaps, partialOverlaps := countOverlaps(file)

	fmt.Println("Part 1:", fullOverlaps)
	fmt.Println("Part 2:", partialOverlaps)
}

func countOverlaps(file *os.File) (fullOverlaps, partialOverlaps int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		range1, range2, ok := parseAssignmentPair(line)
		if !ok {
			continue // Skip invalid lines
		}

		if hasFullOverlap(range1, range2) {
			fullOverlaps++
		}

		if range1.Overlaps(range2) {
			partialOverlaps++
		}
	}

	return fullOverlaps, partialOverlaps
}

// parseAssignmentPair converte uma linha no formato "1-3,2-4" em dois ranges
func parseAssignmentPair(line string) (Range, Range, bool) {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return Range{}, Range{}, false
	}

	range1, ok1 := parseRange(parts[0])
	range2, ok2 := parseRange(parts[1])

	return range1, range2, ok1 && ok2
}

// parseRange converte uma string no formato "1-3" em um Range
func parseRange(s string) (Range, bool) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return Range{}, false
	}

	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return Range{}, false
	}

	return Range{Start: start, End: end}, true
}

// hasFullOverlap verifica se um range contém completamente o outro
func hasFullOverlap(r1, r2 Range) bool {
	return r1.Contains(r2) || r2.Contains(r1)
}
