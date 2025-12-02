package main

import (
	// "bufio"
	// "fmt"
	// "io"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


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

func foo() {
	path := filepath.Join("day1.in")
    dat, err := os.ReadFile(path)
	check(err)
	data := strings.Split(string(dat), "\n")
	count := 0
	sum := 50
	for _, d := range data {
		r := mkRotation(d)
		sum = mod(sum + r, 100)
		if sum == 0 {
			count++
		}
	}
	fmt.Printf("Password: %d\n", count)
}
