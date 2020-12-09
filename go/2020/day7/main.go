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
	data := parseInput(strings.Split(string(input), "\n"))

	result := part1(data)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(data)
	fmt.Printf("Part 2 = %v\n", result)
}

func parseInput(input []string) map[string][]string {
	data := make(map[string][]string)

	for _, s := range input {
		s = strings.ReplaceAll(s, "bags", "")
		s = strings.ReplaceAll(s, "bag", "")

		c := strings.Split(s, " contain ")

		v := strings.ReplaceAll(c[1], ".", "")
		k := strings.TrimSpace(c[0])

		if strings.Contains(c[1], "no other") {
			data[k] = []string{}
		} else {
			data[k] = strings.Split(v, ", ")
		}
	}

	return data
}

func part1(input map[string][]string) int {
	result := 0

	fmt.Println(input)

	for k, v := range input {
		y := false
		if len(v) > 0 {
			y = recurse(input, k)
		}
		if y {
			result++
		}
	}

	return result
}

func recurse(bags map[string][]string, k string) bool {
	y := false
	b := bags[k]
	if len(b) == 0 {
		return false
	}
	for _, v := range b {
		v = strings.TrimSpace(v)
		c := strings.SplitN(v, " ", 2)
		bag := c[1]
		if strings.Contains(bag, "shiny gold") {
			return true
		}
		y = recurse(bags, bag)
		if y {
			break
		}
	}
	return y
}

func part2(input map[string][]string) int {
	fmt.Println(input)

	result := part2Recurse(input, "shiny gold")

	return result
}

func part2Recurse(input map[string][]string, b string) int {
	count := 0

	bags := input[b]

	for _, bag := range bags {
		bag = strings.TrimSpace(bag)

		n, _ := strconv.Atoi(strings.SplitN(bag, " ", 2)[0])
		k := strings.SplitN(bag, " ", 2)[1]

		count += n

		nb := input[k]
		if len(nb) > 0 {
			count += n * part2Recurse(input, k)
		}
	}

	return count
}
