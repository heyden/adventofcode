package main

import (
	"testing"
)

var input = []string{
	"2199943210",
	"3987894921",
	"9856789892",
	"8767896789",
	"9899965678",
}

func TestPart1(t *testing.T) {
	expected := 15
	actual := part1(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	actual := 0
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}
