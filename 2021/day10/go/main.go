package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	result := part1(content)
	fmt.Printf("Part 1 = %v\n", result)

	//result = part2()
	//fmt.Printf("Part 2 = %v\n", result)
}

func part1(data []string) int {
	return 0
}

func part2() int {
	return 0
}
