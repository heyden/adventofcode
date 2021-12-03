package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	result := part1(content)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(content)
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(data []string) int {
	x := 0
	y := 0

	for _, v := range data {
		d := strings.Split(v, " ")
		a, _ := strconv.Atoi(d[1])
		switch d[0] {
		case "forward":
			x += a
		case "down":
			y -= a
		case "up":
			y += a
		}
	}

	return int(math.Abs(float64(x * y)))
}

func part2(data []string) int {
	x := 0
	y := 0
	aim := 0

	for _, v := range data {
		d := strings.Split(v, " ")
		a, _ := strconv.Atoi(d[1])
		switch d[0] {
		case "forward":
			x += a
			y += a * aim
		case "down":
			aim += a
		case "up":
			aim -= a
		}
	}

	return int(x * y)
}
