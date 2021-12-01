package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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
	adapters := []int{}
	for _, v := range data {
		a, _ := strconv.Atoi(v)
		adapters = append(adapters, a)
	}

	sort.Ints(adapters)

	result := part1(adapters)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(adapters)
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(n []int) int {
	a := make([]int, len(n))
	copy(a, n)
	a = append(a, 0)
	sort.Ints(a)

	d := map[int]int{}

	for i, j := range a {
		if i == len(a)-1 {
			d[3]++
			break
		}
		diff := a[i+1] - j
		d[diff]++
	}

	return d[1] * d[3]
}

func part2(n []int) int {
	a := make([]int, len(n))
	copy(a, n)
	a = append(a, 0)
	highest := 0
	for _, v := range a {
		if v >= highest {
			highest = v
		}
	}
	highest += 3
	a = append(a, highest)
	sort.Ints(a)

	counts := []int{}

	for i, v := range a {
		if i == len(a)-1 {
			break
		}

		p := 0
		for _, o := range a[i+1:] {
			if o-v <= 3 {
				p++
			}
		}
		counts = append(counts, p)
	}

	values := []int{}

	i := 0
	for e := len(counts) - 1; e >= 0; e-- {
		v := counts[e]
		if len(values) == 0 {
			values = append(values, v)
			i++
			continue
		}
		switch v {
		case 1:
			p := values[i-1 : i]
			values = append(values, sum(p))
		case 2:
			p := values[i-2 : i]
			values = append(values, sum(p))
		case 3:
			p := values[i-3 : i]
			values = append(values, sum(p))
		}
		i++
	}

	return values[len(values)-1]
}

func sum(a []int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}
