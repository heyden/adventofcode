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
	data := strings.Split(string(input), "\n")
	values := []int{}
	for _, v := range data {
		n, _ := strconv.Atoi(v)
		values = append(values, n)
	}
	result, i := part1(values, 25)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(values, i)
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(input []int, preamble int) (int, int) {
	for i := preamble; i < len(input); i++ {
		found := false
		for n1 := i - preamble; n1 < i; n1++ {
			for n2 := i - preamble; n2 < i; n2++ {
				if n1 == n2 || input[n1] == input[n2] {
					continue
				}
				if input[n1]+input[n2] == input[i] {
					found = true
					break
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return input[i], i
		}
	}

	return 0, 0
}

func part2(input []int, i int) int {
	v := input[i]

	var f []int

	for k1 := 0; k1 < i; k1++ {
		sequence := []int{}
		sequence = append(sequence, input[k1])
		count := input[k1]
		for k2 := k1 + 1; k2 < i; k2++ {
			sequence = append(sequence, input[k2])
			count += input[k2]
			if count == v {
				f = sequence
				break
			}
		}
		if f != nil {
			break
		}
	}

	lowest := 0
	highest := 0
	for _, d := range f {
		if lowest == 0 || d <= lowest {
			lowest = d
		}
		if d >= highest {
			highest = d
		}
	}

	return lowest + highest
}
