package main

import (
	"testing"
)

var (
	input = `abc

a
b
c

ab
ac

a
a
a
a

b`
)

func TestPart1(t *testing.T) {
	data := parseInput(input)

	expected := 11
	actual := part1(data)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	data := parseInput(input)

	expected := 6
	actual := part2(data)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}
