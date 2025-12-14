package main

import (
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func toStr(i int) string {
	s := strconv.Itoa(i)
	return s
}

func main() {
	day8()
}
