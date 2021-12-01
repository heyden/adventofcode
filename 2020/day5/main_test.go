package main

import (
	"testing"
)

var (
	input = []string{
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}
)

func TestPart1(t *testing.T) {
	expected := 820
	actual := part1(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	actual := part2(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}
