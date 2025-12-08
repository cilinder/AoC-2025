package main

import (
	"fmt"
	"os"
	"path/filepath"
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

func day5() {
	filePath := filepath.Join("inputs", "day5.in")
	data, err := os.ReadFile(filePath)
	check(err)
	lines := strings.Split(string(data), "\n")
	freshRanges, ingredients := splitData(lines)
	numFresh := countFresh(ingredients, freshRanges)
	fmt.Printf("There are %d fresh ingredients\n", numFresh)
}
