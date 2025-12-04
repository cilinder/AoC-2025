package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func mkRotation(data string) int {
	switch string(data[0]) {
	case "R":
		val, err := strconv.Atoi(data[1:])
		check(err)
		return val
	case "L":
		val, err := strconv.Atoi(data[1:])
		check(err)
		return -val
	default:
		panic("Invalid input")
	}
}

func mod(n, d int) int {
	x := n % d
	if x >= 0 {
		return x
	} else {
		return x + d
	}
}

func abs(k int) int {
	if k < 0 {
		return -k
	} else {
		return k
	}
}

func countClicks(init, rot int) int {
	int_div := abs((init + rot) / 100)
	if init+rot <= 0 && init != 0 {
		int_div += 1
	}
	return int_div
}

func foo() {
	path := filepath.Join("inputs", "day1.in")
	dat, err := os.ReadFile(path)
	check(err)
	data := strings.Split(string(dat), "\n")
	count := 0
	clicks := 0
	sum := 50
	for _, d := range data {
		r := mkRotation(d)
		clicks += countClicks(sum, r)
		sum = mod(sum+r, 100)
		if sum == 0 {
			count++
		}
	}
	fmt.Printf("Password: %d\n", count)
	fmt.Printf("Password (CLICKS): %d\n", clicks)
}
