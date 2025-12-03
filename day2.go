package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func mkRange(data string) Range {
	vals := strings.Split(data, "-")
	start, err := strconv.Atoi(vals[0])
	check(err)
	end, err := strconv.Atoi(vals[1])
	check(err)
	return Range{start, end}
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

func day2() {
	path := filepath.Join("day2.in")
	dat, err := os.ReadFile(path)
	check(err)
	data := strings.Split(string(dat), ",")
	total := 0
	for _, d := range data {
		rng := mkRange(d)
		start := nextNumber(rng.start)
		end := prevNumber(rng.end)
		repeats := findRepeats(start, end)
		total += repeats
	}
	fmt.Println(total)
}
