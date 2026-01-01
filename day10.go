package main

import (
	"fmt"
	"os"
	"path"
	"slices"
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

func toggleLights(lights []bool, button []int) []bool {
	newLights := make([]bool, len(lights))
	copy(newLights, lights)
	for _, b := range button {
		newLights[b] = !lights[b]
	}
	return newLights
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
		if boolArrayEq(state, targetLights) {
			return 0
		}
		if i >= len(buttons) {
			return -1
		}
		// Try not toggling
		notToggle := solve(state, i+1)
		// Try toggling
		toggledState := toggleLights(state, buttons[i])
		toggle := solve(toggledState, i+1)
		if notToggle < 0 && toggle < 0 {
			return -1
		} else if toggle >= 0 && (notToggle < 0 || notToggle > toggle) {
			return toggle + 1
		} else {
			return notToggle
		}

	}

	init := make([]bool, len(targetLights))
	return solve(init, 0)
}

func joltTooBig(state, target []int) bool {
	for i := range target {
		if state[i] > target[i] {
			return true
		}
	}
	return false
}

func increaseJoltage(state []int, button []int) []int {
	newState := make([]int, len(state))
	copy(newState, state)
	for _, b := range button {
		newState[b]++
	}
	return newState
}

func makeString(state []int) string {
	str := ""
	for _, elt := range state {
		str += toStr(elt) + " "
	}
	return str
}

type memoKey struct {
	state  string
	button int
}

func fewestJoltagePresses(targetJoltage []int, buttons [][]int) int {
	memo := make(map[memoKey]int)
	var solve func(state []int, i int) int
	solve = func(state []int, i int) int {
		key := memoKey{makeString(state), i}
		val, ok := memo[key]
		if ok {
			fmt.Println("using existing value")
			return val
		}
		if slices.Equal(state, targetJoltage) {
			memo[key] = 0
			return 0
		}
		if joltTooBig(state, targetJoltage) || i >= len(buttons) {
			memo[key] = -1
			return -1
		}

		// Skip the button, go to the next one.
		noPress := solve(state, i+1)

		// Press the button, don't move to the next one.
		newState := increaseJoltage(state, buttons[i])
		press := solve(newState, i)

		if noPress < 0 && press < 0 {
			memo[key] = -1
			return -1
		} else if press >= 0 && (noPress < 0 || press < noPress) {
			memo[key] = press + 1
			return press + 1
		} else {
			memo[key] = noPress
			return noPress
		}
	}
	init := make([]int, len(targetJoltage))
	return solve(init, 0)
}

func day10() {
	fileName := path.Join("inputs", "day10.in")
	content, err := os.ReadFile(fileName)
	check(err)
	lines := strings.Split(string(content), "\n")
	totalPresses := 0
	totalJoltPresses := 0
	for _, line := range lines {
		target, buttons, joltage := parseData(line)
		totalPresses += fewestPresses(target, buttons)
		totalJoltPresses += fewestJoltagePresses(joltage, buttons)
		// init := make([]int, len(joltage))
		// init = increaseJoltage(init, buttons[0])
		// init = increaseJoltage(init, buttons[1])
		// init = increaseJoltage(init, buttons[1])
		// init = increaseJoltage(init, buttons[1])
		// fmt.Println(init, slices.Equal(init, joltage))
		// init = increaseJoltage(init, buttons[3])
		// init = increaseJoltage(init, buttons[3])
		// init = increaseJoltage(init, buttons[3])
		// init = increaseJoltage(init, buttons[4])
		// init = increaseJoltage(init, buttons[5])
		// init = increaseJoltage(init, buttons[5])
		// fmt.Println(init, slices.Equal(init, joltage))
		break
	}
	fmt.Println("Total minimal number of presses:", totalPresses)
	fmt.Println("Total minimal number of presses for joltage match:", totalJoltPresses)
}
