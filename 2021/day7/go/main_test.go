package main

import (
	"testing"
)

var input = []string{
	"16",
	"1",
	"2",
	"0",
	"4",
	"2",
	"7",
	"1",
	"2",
	"14",
}

func TestPart1(t *testing.T) {
	expected := 37
	actual := part1(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 168
	actual := part2(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}
