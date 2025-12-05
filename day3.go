package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func maxJoltage(bank string, numBatteries int) int {
	n := len(bank)
	joltage := 0
	iMax := -1
	for battery := numBatteries; battery > 0; battery-- {
		max := 0
		for i := iMax+1; i < n-battery+1; i++ {
			digit := toInt(string(bank[i]))
			if digit > max {
				max = digit
				iMax = i
			}
		}
		joltage = joltage * 10 + max
	}
	return joltage
}

func totalJoltage(rows []string, numBatteries int) int {
	total := 0
	for _, bank := range rows {
		total += maxJoltage(bank, numBatteries)
	}
	return total
}

func day3() {
	path := filepath.Join("inputs", "day3.in")
	dat, err := os.ReadFile(path)
	check(err)
	rows := strings.Split(string(dat), "\n")
	joltage := totalJoltage(rows, 12)
	fmt.Println(joltage)
}
