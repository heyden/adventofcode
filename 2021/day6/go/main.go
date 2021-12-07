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
	content := string(input)

	result := part1(content)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(content)
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(input string) int {
	days := 80

	vals := strings.Split(input, ",")
	fish := []int{}

	for _, v := range vals {
		timer, _ := strconv.Atoi(v)
		fish = append(fish, timer)
	}

	for i := 1; i <= days; i++ {
		fa := 0

		for i, v := range fish {
			if v == 0 {
				fish[i] = 6
				fa++
			} else {
				fish[i] = v - 1
			}
		}

		for ; fa != 0; fa-- {
			fish = append(fish, 8)
		}
	}

	return len(fish)
}

func part2(input string) int {
	days := 256
	vals := strings.Split(input, ",")

	f := make(map[int]int)
	for _, a := range vals {
		age, _ := strconv.Atoi(a)
		if _, ok := f[age]; !ok {
			f[age] = 1
		} else {
			f[age] = f[age] + 1
		}
	}

	for i := 1; i <= days; i++ {
		nf := make(map[int]int)
		for k, v := range f {
			if k == 0 {
				if _, ok := nf[8]; !ok {
					nf[8] = v
				} else {
					nf[8] = nf[8] + v
				}
				if _, ok := nf[6]; !ok {
					nf[6] = v
				} else {
					nf[6] = nf[6] + v
				}
			} else {
				if _, ok := nf[k-1]; !ok {
					nf[k-1] = v
				} else {
					nf[k-1] = nf[k-1] + v
				}
			}
		}
		f = nf
	}

	sum := 0
	for _, v := range f {
		sum += v
	}
	return sum
}
