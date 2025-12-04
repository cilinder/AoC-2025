package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
	repeats []int
}

func mkRange(data string) Range {
	vals := strings.Split(data, "-")
	start, err := strconv.Atoi(vals[0])
	check(err)
	end, err := strconv.Atoi(vals[1])
	check(err)
	return Range{start, end, make([]int, 0)}
}

func nextNumber(num int) int {
	m := len(strconv.Itoa(num))
	if m%2 == 0 {
		return num
	} else {
		return int(math.Pow10(m))
	}
}

func prevNumber(num int) int {
	m := len(strconv.Itoa(num))
	if m%2 == 0 {
		return num
	} else {
		return int(math.Pow10(m-1)) - 1
	}
}

func findTop(num int) int {
	s := toStr(num)
	t := s[:len(s)/2]
	r := toInt(t + t)
	if r <= num {
		return r
	}
	k := toInt(t)
	m := toInt(toStr(k-1) + toStr(k-1))
	return m
}

func findBot(num int) int {
	s := toStr(num)
	t := s[:len(s)/2]
	r := toInt(t + t)
	if r >= num {
		return r
	}
	k := toInt(t)
	m := toInt(toStr(k+1) + toStr(k+1))
	return m
}

func leftHalf(num int) int {
	s := toStr(num)
	t := s[:len(s)/2]
	return toInt(t)
}

func findRepeats(start, end int) int {
	bot := findBot(start)
	top := findTop(end)
	if top < bot {
		return 0
	}
	botLeft := leftHalf(bot)
	topLeft := leftHalf(top)
	total := 0
	for i := botLeft; i <= topLeft; i++ {
		total += toInt(toStr(i) + toStr(i))
	}
	return total
}

func maxOfRanges(ranges []Range) int {
	max := 0
	for _, r := range ranges {
		if r.end >= max {
			max = r.end
		}
	}
	return max
}

func (rng Range) contains(candidate int) bool {
	return rng.start <= candidate && candidate <= rng.end
}

func generateDuplicates(width int, ranges []Range) {
	m := maxOfRanges(ranges)
	mLen := len(toStr(m))
	first := int(math.Pow10(width-1)) 
	last := int(math.Pow10(width)) - 1
	fmt.Printf("Generating duplicates from %d to %d\n", first, last)
	for number := first; number <= last; number++ {
		for i := 2; i <= mLen / width; i++ {
			candidate := toInt(strings.Repeat(toStr(number), i))
			
			for j := range ranges {
				if ranges[j].contains(candidate) && !slices.Contains(ranges[j].repeats, candidate) {
					ranges[j].repeats = append(ranges[j].repeats, candidate)
				}
			}
		}
	}
}

func generateAllDuplicates(ranges []Range) {
	m := maxOfRanges(ranges)
	mLen := len(toStr(m))
	for width := 1; width <= mLen / 2; width++ {
		generateDuplicates(width, ranges)
	}
}

func sumRepeats(ranges []Range) int {
	total := 0
	for	_, rng := range ranges {
		for _, num := range rng.repeats {
			total += num
		}
	} 
	return total
}

func printRanges(ranges []Range) {
	for _, rng := range ranges {
		fmt.Printf("{%d, %d, %v}\n", rng.start, rng.end, rng.repeats)
	}
}

func day2() {
	path := filepath.Join("day2.in")
	dat, err := os.ReadFile(path)
	check(err)
	data := strings.Split(string(dat), ",")
	total := 0
	ranges := make([]Range, 0)
	for _, d := range data {
		rng := mkRange(d)
		ranges = append(ranges, rng)
		start := nextNumber(rng.start)
		end := prevNumber(rng.end)
		repeats := findRepeats(start, end)
		total += repeats
	}
	generateAllDuplicates(ranges)
	printRanges(ranges)
	fmt.Printf("Password is: %d\n", sumRepeats(ranges))
}
