package main

import "testing"

func TestPart1(t *testing.T) {
	testInput := map[string]int{
		"1122":     3,
		"1111":     4,
		"1234":     0,
		"91212129": 9,
	}

	for input, expected := range testInput {
		actual := part1(input)
		if actual != expected {
			t.Errorf("input [%v], actual [%v] != expected [%v]", input, actual, expected)
		}
	}
}

func TestPart2(t *testing.T) {
	testInput := map[string]int{
		"1212":     6,
		"1221":     0,
		"123425":   4,
		"123123":   12,
		"12131415": 4,
	}

	for input, expected := range testInput {
		actual := part2(input)
		if actual != expected {
			t.Errorf("input [%v], actual [%v] != expected [%v]", input, actual, expected)
		}
	}
}
