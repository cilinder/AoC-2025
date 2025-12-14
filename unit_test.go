package main

import (
	"slices"
	"testing"
)

func TestFindRoot(t *testing.T) {
	testCases := []struct {
		setsIn, setsOut []int
		idx, root       int
	}{
		{[]int{0, 0, 0, 2, 3, 1}, []int{0, 0, 0, 2, 3, 1}, 0, 0},
		{[]int{0, 0, 0, 2, 3, 1}, []int{0, 0, 0, 0, 0, 1}, 4, 0},
		{[]int{0, 0, 0, 2, 3, 1}, []int{0, 0, 0, 2, 3, 0}, 5, 0},
		{[]int{0, 0, 2, 2, 3, 1}, []int{0, 0, 2, 2, 2, 1}, 4, 2},
	}

	for i, testCase := range testCases {
		r := findRoot(testCase.idx, testCase.setsIn)
		if r != testCase.root {
			t.Errorf("Test %d: Expected %d, got %d\n", i, testCase.root, r)
		}
		if slices.Compare(testCase.setsIn, testCase.setsOut) != 0 {
			t.Errorf("Test %d: Expected testSets to become %v, got %v\n", i, testCase.setsOut, testCase.setsOut)
		}
	}

}
