package main

import (
	"testing"
)

var (
	input = []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
)

func TestPart1(t *testing.T) {
	expected := 25
	actual := part1(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 286
	actual := part2(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}
