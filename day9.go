package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type RedTile struct {
	x int
	y int
}

func parseTiles(rows []string) []RedTile {
	tiles := make([]RedTile, 0)
	for _, row := range rows {
		parts := strings.Split(row, ",")
		tiles = append(tiles, RedTile{toInt(parts[0]), toInt(parts[1])})
	}
	return tiles
} 

func area(tile1, tile2 RedTile) int {
	return (abs(tile2.x - tile1.x) + 1) * (abs(tile2.y - tile1.y) + 1)
}

func findMaxArea(tiles []RedTile) int {
	maxArea := 0
	for i, tile1 := range tiles {
		for j := i+1; j < len(tiles); j++ {
			tile2 := tiles[j]
			tileArea := area(tile1, tile2)
			if tileArea > maxArea {
				maxArea = tileArea
			}
		}
	}
	return maxArea
}

func makeGrid(tiles []RedTile) 

func day9() {
	fileName := filepath.Join("inputs", "day9.in")
	contents, err := os.ReadFile(fileName)
	check(err)
	rows := strings.Split(string(contents), "\n")
	redTiles := parseTiles(rows)
	maxArea := findMaxArea(redTiles)
	fmt.Printf("Largest area is %d\n", maxArea)
}