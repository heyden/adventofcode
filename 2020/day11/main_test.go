package main

import (
	"strings"
	"testing"
)

var (
	input = []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
)

func TestPart1(t *testing.T) {
	seats := [][]string{}
	for _, v := range input {
		seats = append(seats, strings.Split(v, ""))
	}

	expected := 37
	actual := part1(seats)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	seats := [][]string{}
	for _, v := range input {
		seats = append(seats, strings.Split(v, ""))
	}
	expected := 26
	actual := part2(seats)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}
