package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func splitData(lines []string) ([][]int, []int){
	freshRanges := make([][]int, 0)
	ingredients := make([]int, 0)
	flag := true
	for _, line := range lines {
		if line == "" {
			flag = false
			continue
		}
		if flag {
			r := strings.Split(line, "-")
			freshRanges = append(freshRanges, []int{toInt(r[0]), toInt(r[1])})
		} else {
			ingredients = append(ingredients, toInt(line))
		}
	}
	return freshRanges, ingredients
}

func isFresh(ingredient int, freshRanges [][]int) bool {
	for _, r := range freshRanges {
		if r[0] <= ingredient && ingredient <= r[1] {
			return true
		}
	}
	return false
}

func countFresh(ingredients []int, freshRanges [][]int) int {
	total := 0
	for _, ingredient := range ingredients {
		if isFresh(ingredient, freshRanges) {
			total++
		}
	}
	return total
}

func mergeRanges(freshRanges [][]int) [][]int {
	slices.SortFunc(freshRanges, func(r1, r2 []int) int {
		if r1[0] == r2[0] {
			return r1[1] - r2[1]
		} else {
			return r1[0] - r2[0]
		}
	})

	merged := make([][]int, 0)
	currRange := freshRanges[0]
	for i := 1; i < len(freshRanges); i++ {
		r := freshRanges[i]
		if r[0] <= currRange[1] {
			currRange[1] = max(currRange[1], r[1])
		} else {
			merged = append(merged, currRange)
			currRange = r
		}
	}
	merged = append(merged, currRange)
	return merged
}

func countPotentialFresh(freshRanges [][]int) int {
	total := 0
	for _, r := range freshRanges {
		total += r[1] - r[0] + 1
	}
	return total
}

func day5() {
	filePath := filepath.Join("inputs", "day5.in")
	data, err := os.ReadFile(filePath)
	check(err)
	lines := strings.Split(string(data), "\n")
	freshRanges, ingredients := splitData(lines)
	numFresh := countFresh(ingredients, freshRanges)
	fmt.Printf("There are %d fresh ingredients\n", numFresh)
	merged := mergeRanges(freshRanges)
	fmt.Printf("Total potential fresh ingredients: %d\n", countPotentialFresh(merged))
}
