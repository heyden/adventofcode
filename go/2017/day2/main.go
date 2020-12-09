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
	spreadsheet := [][]int{}
	lines := strings.Split(string(input), "\n")

	for _, l := range lines {
		values := strings.Split(l, "\t")
		data := []int{}
		for _, v := range values {
			value, _ := strconv.Atoi(v)
			data = append(data, value)
		}
		spreadsheet = append(spreadsheet, data)
	}
	result := part1(spreadsheet)
	fmt.Printf("Part 1 = %v\n", result)
	part2(spreadsheet)
}

func part1(data [][]int) int {
	differences := []int{}

	for _, row := range data {
		smallest := 0
		largest := 0

		for _, val := range row {
			if val > largest {
				largest = val
			}

			if smallest == 0 || val < smallest {
				smallest = val
			}
		}

		difference := largest - smallest
		differences = append(differences, difference)
	}

	sum := 0
	for _, val := range differences {
		sum += val
	}
	return sum
}

func part2(data [][]int) int {
	divisions := []int{}

	for _, row := range data {
		fmt.Println(row)
	}

	sum := 0
	for _, val := range divisions {
		sum += val
	}
	return sum
}
