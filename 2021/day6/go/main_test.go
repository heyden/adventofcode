package main

import (
	"testing"
)

var input = "3,4,3,1,2"

func TestPart1(t *testing.T) {
	expected := 5934
	actual := part1(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 26984457539
	actual := part2(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}
