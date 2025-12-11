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

func day7() {
	fileName := filepath.Join("inputs", "day7.in")
	contents, err := os.ReadFile(fileName)
	check(err)
	rows := strings.Split(string(contents), "\n")
	splits := simulateBeam(rows)
	// printManifold(rows)
	fmt.Printf("Total splits: %d\n", splits)
}