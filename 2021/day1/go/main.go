package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading input file %v\n", err)
		os.Exit(1)
	}

	// transform input to expected type
	content := strings.Split(string(input), "\n")
	data := []int{}

	for _, v := range content {
		value, _ := strconv.Atoi(v)
		data = append(data, value)
	}

	result := part1(data)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(data)
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(depths []int) int {
	increased := 0
	for i, v := range depths {
		if i != 0 && depths[i-1] < v {
			increased++
		}
	}

	return increased
}

func part2(depths []int) int {
	increased := 0
	psum := 0
	for i, _ := range depths {
		if i+3 > len(depths) {
			break
		}
		sum := depths[i] + depths[i+1] + depths[i+2]
		if i != 0 && sum > psum {
			increased++
		}
		psum = sum
	}

	return increased
}
