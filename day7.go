package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printManifold(manifold []string) {
	for _, row := range manifold {
		fmt.Println(row)
	}
}

func printCounts(counts [][]int) {
	for _, row := range counts {
		fmt.Println(row)
	}
}

func replaceAtIndex(str string, replacement rune, index int) string {
    return str[:index] + string(replacement) + str[index+1:]
}

func simulateBeam(tachyonManifold []string) int {
	splits := 0
	tachyonManifold[0] = strings.Replace(tachyonManifold[0], "S", "|", 1)
	for i := 1; i < len(tachyonManifold); i++ {
		prev := tachyonManifold[i-1]
		curr := tachyonManifold[i]
		for j := 0; j < len(curr); j++ {
			if curr[j] == '.' && prev[j] == '|' {
				tachyonManifold[i] = replaceAtIndex(tachyonManifold[i], '|', j)
			} else if curr[j] == '^' && prev[j] == '|' {
				tachyonManifold[i] = replaceAtIndex(tachyonManifold[i], '|', j-1)
				tachyonManifold[i] = replaceAtIndex(tachyonManifold[i], '|', j+1)
				splits += 1
			}
		}
	}
	return splits
}

func manyWorlds(tachyonManifold []string) [][]int {
	tachyonManifold[0] = strings.Replace(tachyonManifold[0], "S", "|", 1)
	counts := make([][]int, len(tachyonManifold))
	for i := 0; i < len(tachyonManifold); i++ {
		counts[i] = make([]int, len(tachyonManifold[0]))
	}
	for i := 0; i < len(counts[0]); i++ {
		if tachyonManifold[0][i] == '|' {
			counts[0][i] = 1
		}
	}
	for i := 1; i < len(tachyonManifold); i++ {
		prev := tachyonManifold[i-1]
		curr := tachyonManifold[i]
		for j := 0; j < len(curr); j++ {
			count := 0
			if curr[j] == '|' {
				if prev[j] == '|' {
					count += counts[i-1][j]
				}
				if j > 0 && curr[j-1] == '^' {
					count += counts[i-1][j-1]
				}
				if j < len(curr)-1 && curr[j+1] == '^' {
					count += counts[i-1][j+1]
				}
			}
			counts[i][j] = count
		}
	}
	return counts
}

func day7() {
	fileName := filepath.Join("inputs", "day7.in")
	contents, err := os.ReadFile(fileName)
	check(err)
	rows := strings.Split(string(contents), "\n")
	splits := simulateBeam(rows)
	fmt.Printf("Total splits: %d\n", splits)
	counts := manyWorlds(rows)
	total := 0
	for i := 0; i < len(counts[0]); i++ {
		total += counts[len(counts)-1][i]
	}
	fmt.Println(total)
}