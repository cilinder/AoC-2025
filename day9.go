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

type Directions struct {
	north bool
	south bool
	east  bool
	west  bool
}

type DirectionGrid struct {
	offsetX int
	offsetY int
	flags   [][]Directions
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
	return (abs(tile2.x-tile1.x) + 1) * (abs(tile2.y-tile1.y) + 1)
}

func findMaxArea(tiles []RedTile) int {
	maxArea := 0
	for i, tile1 := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			tile2 := tiles[j]
			tileArea := area(tile1, tile2)
			if tileArea > maxArea {
				maxArea = tileArea
			}
		}
	}
	return maxArea
}

func makeGrid(tiles []RedTile) DirectionGrid {
	minX := tiles[0].x
	maxX := tiles[0].x
	minY := tiles[0].y
	maxY := tiles[0].y
	for _, tile := range tiles {
		minX = min(minX, tile.x)
		minY = min(minY, tile.y)
		maxX = max(maxX, tile.x)
		maxY = max(maxY, tile.y)
	}
	dimX := maxX - minX + 1
	dimY := maxY - minY + 1
	grid := make([][]Directions, dimY)
	for i := range dimX {
		grid[i] = make([]Directions, dimX)
	}
	return DirectionGrid{minX, minY, grid}
}

func (grid DirectionGrid) fill(tiles []RedTile) {

}

func day9() {
	fileName := filepath.Join("inputs", "day9.in")
	contents, err := os.ReadFile(fileName)
	check(err)
	rows := strings.Split(string(contents), "\n")
	redTiles := parseTiles(rows)
	maxArea := findMaxArea(redTiles)
	fmt.Printf("Largest area is %d\n", maxArea)
}
