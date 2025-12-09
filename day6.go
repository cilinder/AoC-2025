package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func splitLine(text string) []string {
	data := make([]string, 0)
	currElt := ""
	for _, chr := range text {
		if chr == ' ' && len(currElt) != 0 {
			data = append(data, currElt)
			currElt = ""
		} else if chr != ' ' {
			currElt += string(chr)
		}
	}
	if len(currElt) > 0 {
		data = append(data, currElt)
	}
	return data
}

func parse(rows []string) ([][]int, []string) {
	nRows := len(rows)
	nCols := len(splitLine(rows[0]))

	data := make([][]int, nCols)
	for i := 0; i < nRows-1; i++ {
		for j, elt := range splitLine(rows[i]) {
			data[j] = append(data[j], toInt(elt))
		}
	}
	return data, splitLine(rows[nRows-1])
}

func parse2(rows []string) ([][]int, []string) {
	nRows := len(rows)
	nCols := len(rows[0])
	problems := make([][]int, 0)
	problem := make([]int, 0)
	for i := nCols - 1; i >= 0; i-- {
		num := 0
		anyNonEmpty := false
		for j := 0; j < nRows-1; j++ {
			if rows[j][i] != ' ' {
				elt := toInt(string(rows[j][i]))
				num = num*10 + elt
				anyNonEmpty = true
			}
		}
		if anyNonEmpty {
			problem = append(problem, num)
		} else {
			problems = append(problems, problem)
			problem = make([]int, 0)
		}
	}
	problems = append(problems, problem)
	ops := splitLine(rows[nRows-1])
	slices.Reverse(ops)
	return problems, ops
}

func aggregate(data []int, op string) int {
	var total int
	if op == "+" {
		total = 0
	} else {
		total = 1
	}
	for _, num := range data {
		if op == "+" {
			total += num
		} else {
			total *= num
		}
	}
	return total
}

func solve(data [][]int, ops []string) int {
	total := 0
	for i := range data {
		total += aggregate(data[i], ops[i])
	}
	return total
}

func day6() {
	fileName := filepath.Join("inputs", "day6.in")
	contents, err := os.ReadFile(fileName)
	check(err)
	rows := strings.Split(string(contents), "\n")
	data, ops := parse(rows)
	fmt.Printf("Solution: %d\n", solve(data, ops))

	problems, ops := parse2(rows)
	fmt.Printf("Really correct solution: %d\n", solve(problems, ops))
}
