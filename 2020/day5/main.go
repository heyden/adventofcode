package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

	result := part1(data)
	fmt.Printf("Part 1 = %v\n", result)
	result = part2(data)
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(b []string) int {
	result := 0

	for _, v := range b {
		r := evaluate(v[0:7], [2]int{0, 127}, "F", "B")
		c := evaluate(v[7:], [2]int{0, 7}, "L", "R")
		seatID := r*8 + c

		if result < seatID {
			result = seatID
		}
	}
	return result
}

func part2(b []string) int {
	seats := []int{}

	for _, v := range b {
		r := evaluate(v[0:7], [2]int{0, 127}, "F", "B")
		c := evaluate(v[7:], [2]int{0, 7}, "L", "R")
		seatID := r*8 + c
		seats = append(seats, seatID)
	}

	sort.Ints(seats)

	possible := 0
	for i, v := range seats {
		if i == 0 {
			continue
		}

		if i != len(seats)-1 && v != seats[i+1]-1 {
			possible = v + 1
			if possible+1 == seats[i+1] {
				break
			}
		}
	}

	return possible
}

func evaluate(t string, set [2]int, lower, upper string) int {
	v := 0

	r := make([]int, len(t))

	for i := range r {
		d := string(t[i])

		m := (set[1] + set[0]) / 2

		if d == upper {
			set[0] = m + 1
		} else if d == lower {
			set[1] = m
		}

		if i == len(r)-1 {
			switch d {
			case upper:
				v = set[1]
			case lower:
				v = set[0]
			}
		}
	}
	return v
}
