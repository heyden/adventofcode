package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading input file %v\n", err)
		os.Exit(1)
	}

	// transform input to expected type

	result := part1(string(input))
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(string(input))
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(input string) int {
	var matches []int

	for i, c := range input {
		num1, _ := strconv.Atoi(string(c))

		nextPos := i + 1
		if i == len(input)-1 {
			nextPos = 0
		}

		num2, _ := strconv.Atoi(string(input[nextPos]))
		if num1 == num2 {
			matches = append(matches, num1)
		}
	}

	sum := 0
	for _, val := range matches {
		sum += val
	}
	return sum
}

func part2(input string) int {
	var matches []int

	for i, c := range input {
		num1, _ := strconv.Atoi(string(c))

		skip := len(input) / 2
		nextPos := i + skip

		l := len(input)
		if nextPos >= l {
			nextPos = nextPos - l
		}

		num2, _ := strconv.Atoi(string(input[nextPos]))
		if num1 == num2 {
			matches = append(matches, num1)
		}
	}

	sum := 0
	for _, val := range matches {
		sum += val
	}
	return sum
}
