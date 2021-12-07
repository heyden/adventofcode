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
	content := strings.Split(string(input), ",")

	result := part1(content)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(content)
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(data []string) int {
	positions := []int{}
	max := 0
	for _, v := range data {
		p, _ := strconv.Atoi(v)
		positions = append(positions, p)
		if p > max {
			max = p
		}
	}

	lf := -1
	for i := 0; i <= max; i++ {
		fuel := 0
		for _, v := range positions {
			if v < i {
				fuel += i - v
			} else if v > i {
				fuel += v - i
			}
		}

		if lf == -1 {
			lf = fuel
		} else if lf > fuel {
			lf = fuel
		}
	}

	return lf
}

func part2(data []string) int {
	positions := []int{}
	max := 0
	for _, v := range data {
		p, _ := strconv.Atoi(v)
		positions = append(positions, p)
		if p > max {
			max = p
		}
	}

	lf := -1
	for i := 0; i <= max; i++ {
		fuel := 0
		for _, v := range positions {
			steps := 0
			if v < i {
				steps += i - v
			} else if v > i {
				steps += v - i
			}
			fuel += (int(math.Pow(float64(steps), float64(2))) + steps) / 2
		}

		if lf == -1 {
			lf = fuel
		} else if lf > fuel {
			lf = fuel
		}
	}

	return lf
}
