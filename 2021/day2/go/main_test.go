package main

import (
	"testing"
)

var input = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestPart1(t *testing.T) {
	expected := 150
	actual := part1(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 900
	actual := part2(input)
	if actual != expected {
		t.Errorf("actual [%v] != expected [%v]", actual, expected)
	}
}
