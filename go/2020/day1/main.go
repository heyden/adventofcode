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

func part1(input []int) int {
	found := false
	result := 0

	for i, v1 := range input {
		for j, v2 := range input {
			if i == j {
				continue
			}
			if v1+v2 == 2020 {
				found = true
				result = v1 * v2
				break
			}
		}
		if found {
			break
		}
	}

	return result
}

func part2(input []int) int {
	found := false
	result := 0
	for i, v1 := range input {
		for i2, v2 := range input {
			if i == i2 {
				continue
			}
			for i3, v3 := range input {
				if i3 == i || i3 == i2 {
					continue
				}
				if v1+v2+v3 == 2020 {
					found = true
					result = v1 * v2 * v3
					break
				}
			}
		}
		if found {
			break
		}
	}

	return result
}
