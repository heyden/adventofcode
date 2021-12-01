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
	data := strings.Split(string(input), "\n")
	matrix := buildMatrix(data)

	result := part1(matrix, 3, 1)
	fmt.Printf("Part 1 = %v\n", result)

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	result = part2(matrix, slopes)
	fmt.Printf("Part 2 = %v\n", result)
}

func buildMatrix(input []string) [][]string {
	matrix := [][]string{}
	for _, v := range input {
		row := []string{}
		for _, c := range v {
			row = append(row, string(c))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func part1(input [][]string, right, down int) int {
	t := 0
	x := 0
	for i, v := range input {
		if i != 0 && v[x] == "#" {
			t++
		}
		x += 3

		l := len(v)
		if x >= l {
			x = x - l
		}
	}
	return t
}

func part2(matrix [][]string, slopes [][]int) int {
	t := []int{}

	for _, slope := range slopes {
		fmt.Printf("Evaluating slope %v %v\n", slope[0], slope[1])
		s := evaluateSlope(matrix, slope[0], slope[1])
		t = append(t, s)
	}

	result := 1
	for _, v := range t {
		result *= v
	}
	return result
}

func evaluateSlope(matrix [][]string, x, y int) int {
	t := 0
	posx := 0
	posy := 0

	for {
		posx += x
		posy += y

		if posy >= len(matrix) {
			break
		}

		fmt.Printf("%v %v\n", posx, posy)

		row := matrix[posy]
		l := len(row)
		if posx >= l {
			posx = posx - l
		}

		if matrix[posy][posx] == "#" {
			t++
		}
	}

	return t
}
