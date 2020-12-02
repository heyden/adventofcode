package main

import "testing"

var (
	input = []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
)

func TestPart1(t *testing.T) {
	expected := 514579

	actual := part1(input)
	if actual != expected {
		t.Errorf("input [%v], actual [%v] != expected [%v]", input, actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 241861950

	actual := part2(input)
	if actual != expected {
		t.Errorf("input [%v], actual [%v] != expected [%v]", input, actual, expected)
	}
}
