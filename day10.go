package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func trim(s string) string {
	n := len(s)
	return s[1 : n-1]
}

func parseData(data string) ([]bool, [][]int, []int) {
	parts := strings.Split(data, " ")
	lightsData := trim(parts[0])
	lights := make([]bool, 0)
	for _, x := range lightsData {
		if x == '#' {
			lights = append(lights, true)
		} else {
			lights = append(lights, false)
		}
	}

	buttonsData := parts[1 : len(parts)-1]
	buttons := make([][]int, 0)
	for _, buttonConfig := range buttonsData {
		x := make([]int, 0)
		for _, button := range strings.Split(trim(buttonConfig), ",") {
			x = append(x, toInt(button))
		}
		buttons = append(buttons, x)
	}

	joltageData := parts[len(parts)-1]
	joltage := make([]int, 0)
	for _, jolt := range strings.Split(trim(joltageData), ",") {
		joltage = append(joltage, toInt(jolt))
	}

	return lights, buttons, joltage
}

func toggleLights(lights []bool, button []int) {
	for _, b := range button {
		lights[b] = !lights[b]
	}
}

func boolArrayEq(arr1, arr2 []bool) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func fewestPresses(targetLights []bool, buttons [][]int) int {
	var solve func(state []bool, i int) int

	solve = func(state []bool, i int) int {
		if i >= len(buttons) {
			return -1
		}
		if boolArrayEq(state, targetLights) {
			return 0
		}
		// Try not toggling
		notToggle := solve(state, i+1)
		toggleLights(state, buttons[i])
		toggle := solve(state, i+1)
		toggleLights(state, buttons[i]) // Reverse
		if notToggle < 0 && toggle < 0 {
			return -1
		} else if notToggle < 0 || toggle < 0 {
			return max(toggle, notToggle) + 1
		} else {
			return min(toggle, notToggle) + 1
		}

	}

	init := make([]bool, len(targetLights))
	return solve(init, 0)
}

func day10() {
	fileName := path.Join("inputs", "day10_sample.in")
	content, err := os.ReadFile(fileName)
	check(err)
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		target, buttons, _ := parseData(line)
		fmt.Println(fewestPresses(target, buttons))
	}
}
