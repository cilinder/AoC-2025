package main

import (
	"fmt"
	"iter"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

type Node struct {
	pointId int
	next    int
}

func parsePoint(row string) Point {
	elts := strings.Split(row, ",")
	x := toInt(elts[0])
	y := toInt(elts[1])
	z := toInt(elts[2])
	return Point{x, y, z}
}

func initMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range n {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func square(num int) int {
	return num * num
}

func (p Point) dist(q Point) int {
	d := square(p.x-q.x) + square(p.y-q.y) + square(p.z-q.z)
	return d
}

func pairDists(points []Point) [][]int {
	matrix := initMatrix(len(points))
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			matrix[i][j] = points[i].dist(points[j])
			matrix[j][i] = matrix[i][j]
		}
	}
	return matrix
}

func distances(points []Point) [][]int {
	n := len(points)
	dists := make([][]int, 0)
	for i := range points {
		for j := i + 1; j < n; j++ {
			d := points[i].dist(points[j])
			dists = append(dists, []int{d, i, j})
		}
	}
	slices.SortFunc(dists, func(d1, d2 []int) int { return d1[0] - d2[0] })
	return dists
}

func initSets(n int) []int {
	sets := make([]int, n)
	for i := 0; i < n; i++ {
		sets[i] = i
	}
	return sets
}

func findRoot(p int, sets []int) int {
	for p != sets[p] {
		sets[p] = findRoot(sets[p], sets)
		p = sets[p]
	}
	return p
}

func connect(points []Point, maxConns int) []int {
	dists := distances(points)
	sets := initSets(len(points))
	for i := 0; i < min(maxConns, len(dists)); i++ {
		p := dists[i][1]
		q := dists[i][2]
		pRoot := findRoot(p, sets)
		qRoot := findRoot(q, sets)
		sets[pRoot] = qRoot
	}
	for i := range sets {
		sets[i] = findRoot(i, sets)
	}
	return sets
}

func countComponents(sets []int) int {
	for i := range sets {
		sets[i] = findRoot(i, sets)
	}
	counted := make([]int, 0)
	count := 0
	for _, el := range sets {
		if !slices.Contains(counted, el) {
			count += 1
			counted = append(counted, el)
		}
	}
	return count
}

func connectAll(points []Point) int {
	dists := distances(points)
	sets := initSets(len(points))
	var p, q int
	for i := 0; i < len(dists) && countComponents(sets) > 1; i++ {
		p = dists[i][1]
		q = dists[i][2]
		pRoot := findRoot(p, sets)
		qRoot := findRoot(q, sets)
		// fmt.Printf("Connecting %d (%d, %d, %d) and %d (%d, %d, %d)\n",
		// p, points[p].x, points[p].y, points[p].z, q, points[q].x, points[q].y, points[q].z)
		sets[pRoot] = qRoot
		// fmt.Println(sets)
		// fmt.Println()
	}
	return points[p].x * points[q].x
}

func countDistinct(lst []int) map[int]int {
	groupCounts := make(map[int]int)
	for _, el := range lst {
		groupCounts[el] += 1
	}
	return groupCounts
}

func find3Largest(lst iter.Seq[int]) (int, int, int) {
	first := -1
	second := -1
	third := -1
	for el := range lst {
		if el > first {
			third = second
			second = first
			first = el
		} else if el > second {
			third = second
			second = el
		} else if el > third {
			third = el
		}
	}
	return first, second, third
}

func day8() {
	fileName := filepath.Join("inputs", "day8.in")
	contents, err := os.ReadFile(fileName)
	check(err)
	rows := strings.Split(string(contents), "\n")
	points := make([]Point, len(rows))
	for i, row := range rows {
		points[i] = parsePoint(row)
	}
	sets := connect(points, 1000)
	numbers := make([]int, 0)
	for i := range sets {
		numbers = append(numbers, i)
	}
	counts := countDistinct(sets)
	fst, snd, trd := find3Largest(maps.Values(counts))
	fmt.Printf("Result part 1: %d\n", fst*snd*trd)
	part2 := connectAll(points)
	fmt.Printf("Result part 2: %d\n", part2)

}
