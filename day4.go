package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func countNeighbours(grid [][]bool, i, j int) int {
	total := 0
	for di := max(0, i-1); di <= min(i+1, len(grid)-1); di++ {
		for dj := max(0, j-1); dj <= min(j+1, len(grid[0])-1); dj++ {
			if grid[di][dj] && (di != i || dj != j) {
				total += 1
			}
		}
	}
	return total
}

func countAccessible(grid [][]bool) int {
	accessible := 0
	for i, row := range grid {
		for j, el := range row {
			if el && countNeighbours(grid, i, j) < 4 {
				accessible += 1
			}
		}
	}

	return accessible
}

func removeAccessible(grid [][]bool) int {
	accessible := 0
	for i, row := range grid {
		for j, el := range row {
			if el && countNeighbours(grid, i, j) < 4 {
				accessible += 1
				grid[i][j] = false
			}
		}
	}
	return accessible
}

func clearRolls(grid [][]bool) int {
	totalRemoved := 0
	numRemoved := 1000
	for numRemoved > 0 {
		numRemoved = removeAccessible(grid)
		totalRemoved += numRemoved
	}
	return totalRemoved
}

func boolToPaper(b bool) string {
	switch b {
	case true:
		return "@"
	default:
		return "."
	}
}

func printGrid(grid [][]bool) {
	for _, row := range grid {
		for _, el := range row {
			fmt.Print(boolToPaper(el))
		}
		fmt.Println()
	}
	fmt.Println()
}

func day4() {
	filePath := filepath.Join("inputs", "day4.in")
	data, err := os.ReadFile(filePath)
	check(err)
	grid := make([][]bool, 0)
	for dataRow := range strings.SplitSeq(string(data), "\n") {
		row := make([]bool, 0)
		for _, chr := range dataRow {
			row = append(row, chr == '@')
		}
		grid = append(grid, row)
	}

	accessible := countAccessible(grid)
	fmt.Printf("Number of accessible rolls of paper: %d\n", accessible)

	removed := clearRolls(grid)
	fmt.Printf("Total rows removed: %d\n", removed)
}
