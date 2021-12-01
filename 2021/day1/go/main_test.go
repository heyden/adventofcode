package main

import (
	"testing"
)

var input = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func TestPart1(t *testing.T) {
	expected := 7
	actual := part1(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 5
	actual := part2(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}
