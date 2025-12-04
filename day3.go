package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func maxJoltage(bank string) int {
	n := len(bank)
	max := 0
	sndMax := 0
	var iMax int
	for i := 0; i < n-1; i++ {
		digit := toInt(string(bank[i]))
		if digit > max {
			max = digit
			iMax = i
		}
	}
	for i := iMax + 1; i < n; i++ {
		digit := toInt(string(bank[i]))
		if digit > sndMax {
			sndMax = digit
		}
	}
	return max*10 + sndMax

}

func totalJoltage(rows []string) int {
	total := 0
	for _, bank := range rows {
		total += maxJoltage(bank)
	}
	return total
}

func day3() {
	path := filepath.Join("inputs", "day3.in")
	dat, err := os.ReadFile(path)
	check(err)
	rows := strings.Split(string(dat), "\n")
	joltage := totalJoltage(rows)
	fmt.Println(joltage)
}
