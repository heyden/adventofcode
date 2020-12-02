package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := [][]int{
		[]int{5, 1, 9, 5},
		[]int{7, 5, 3},
		[]int{2, 4, 6, 8},
	}

	expected := 18

	actual := part1(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	input := [][]int{
		[]int{5, 9, 2, 8},
		[]int{9, 4, 7, 3},
		[]int{3, 8, 6, 5},
	}

	expected := 9

	actual := part2(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}
