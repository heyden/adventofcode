package main

import (
	"testing"
)

var (
	input = []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
)

func TestPart1(t *testing.T) {
	expected := 2
	actual := part1(input)
	if actual != expected {
		t.Errorf("input [%v], actual [%v] != expected [%v]", input, actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 1
	actual := part2(input)
	if actual != expected {
		t.Errorf("input [%v], actual [%v] != expected [%v]", input, actual, expected)
	}
}
