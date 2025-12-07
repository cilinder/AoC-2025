package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func countNeighbours(grid []string, i, j int) int {
	total := 0
	for di := max(0, i-1); di <= min(i+1, len(grid)-1); di++ {
		for dj := max(0, j-1); dj <= min(j+1, len(grid[0])-1); dj++ {
			if grid[di][dj] == '@' && (di != i || dj != j) {
				total += 1
			}
		}
	}
	return total
}

func countAccessible(grid []string) int {
	accessible := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '@' && countNeighbours(grid, i, j) < 4 {
				accessible += 1
			}
		}
	}

	return accessible
}

func day4() {
	filePath := filepath.Join("inputs", "day4.in")
	data, err := os.ReadFile(filePath)
	check(err)
	rows := strings.Split(string(data), "\n")
	accessible := countAccessible(rows)
	fmt.Printf("Number of accessible rolls of paper: %d\n", accessible)

}
