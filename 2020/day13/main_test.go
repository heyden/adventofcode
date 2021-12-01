package main

import (
	"strconv"
	"testing"
)

var (
	input = []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}
)

func TestPart1(t *testing.T) {
	expected := 295
	actual := part1(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	input := map[string]string{
		"1068781":    "7,13,x,x,59,x,31,19",
		"3417":       "17,x,13,19",
		"754018":     "67,7,59,61",
		"779210":     "67,x,7,59,61",
		"1261476":    "67,7,x,59,61",
		"1202161486": "1789,37,47,1889",
	}

	for k, v := range input {
		expected, _ := strconv.Atoi(k)
		actual := part2(v)
		if actual != expected {
			t.Errorf("actual [%v] != expected [%v]", actual, expected)
		}
	}
}
