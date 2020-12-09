package main

import (
	"testing"
)

var (
	input = []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
)

func TestPart1(t *testing.T) {
	expected := 5
	actual := part1(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 8
	actual := part2(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}
