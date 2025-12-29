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
	width int
	height int
	flags   [][]Directions
}

func (grid DirectionGrid) get(i, j int) Directions {
	return grid.flags[i][j]
}

func (grid DirectionGrid) set(i, j int, dir Directions) {
	grid.flags[i][j] = dir
}

func (grid DirectionGrid) setNorth(i, j int, val bool) {
	grid.flags[i][j].north = val
}

func (grid DirectionGrid) setSouth(i, j int, val bool) {
	grid.flags[i][j].south = val
}

func (grid DirectionGrid) setEast(i, j int, val bool) {
	grid.flags[i][j].east = val
}

func (grid DirectionGrid) setWest(i, j int, val bool) {
	grid.flags[i][j].west = val
}

func (grid DirectionGrid) print() {
	for i := range grid.height {
		for j := range grid.width {
			flags := grid.get(i, j)
			if flags.east && flags.west && flags.north && flags.south {
				fmt.Printf("X")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
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
	for i := range tiles {
		tiles[i].x = tiles[i].x - minX
		tiles[i].y = tiles[i].y - minY
	}
	dimX := maxX - minX + 1
	dimY := maxY - minY + 1
	grid := make([][]Directions, dimY)
	for i := range dimY {
		grid[i] = make([]Directions, dimX)
	}
	fmt.Printf("dimX: %d, dimY: %d\n", dimX, dimY)
	return DirectionGrid{dimX, dimY, grid}
}

func fillInterior(grid DirectionGrid, tiles []RedTile) {
	for i := 1; i < len(tiles); i++ {
		t1 := tiles[i-1]
		t2 := tiles[i]

		if t1.x == t2.x {
			for j := min(t1.y, t2.y); j < max(t1.y, t2.y); j++ {
				for k := 0; k < grid.width; k++ {
					if k < t1.x {
						grid.setEast(j, k, !grid.get(j, k).east)
					}
					if k > t1.x {
						grid.setWest(j, k, !grid.get(j, k).west)
					}
				}
			}
		} else if t1.y == t2.y {
			for j := min(t1.x, t2.x); j < max(t1.x, t2.x); j++ {
				for k := 0; k < grid.height; k++ {
					if k < t1.y {
						grid.setNorth(k, j, !grid.get(k, j).north)
					}
					if k > t1.y {
						grid.setSouth(k, j, !grid.get(k, j).south)
					}
				}
			}
		} else {
			panic("Invalid state")
		}
	}
}

func day9() {
	fileName := filepath.Join("inputs", "day9_sample.in")
	contents, err := os.ReadFile(fileName)
	check(err)
	rows := strings.Split(string(contents), "\n")
	redTiles := parseTiles(rows)
	maxArea := findMaxArea(redTiles)
	fmt.Printf("Largest area is %d\n", maxArea)
	grid := makeGrid(redTiles)
	grid.print()
	fmt.Println(redTiles)
	fillInterior(grid, redTiles)
	grid.print()
}
