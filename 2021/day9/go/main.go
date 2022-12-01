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

	result := part1(content)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(content)
	fmt.Printf("Part 2 = %v\n", result)
}

func isLower(p int, c []int) bool {
	for _, v := range c {
		if v <= p {
			return false
		}
	}
	return true
}

func part1(data []string) int {
	points := [][]int{}

	for _, p := range data {
		r := strings.Split(p, "")
		rp := []int{}
		for _, point := range r {
			val, _ := strconv.Atoi(point)
			rp = append(rp, val)
		}
		points = append(points, rp)
	}

	lp := []int{}

	for i, r := range points {
		for j, p := range r {
			adj := []int{}

			// top
			if i != 0 {
				adj = append(adj, points[i-1][j])
			}

			// bottom
			if i != len(points)-1 {
				adj = append(adj, points[i+1][j])
			}

			// left
			if j != 0 {
				adj = append(adj, r[j-1])
			}

			// right
			if j != len(r)-1 {
				adj = append(adj, r[j+1])
			}

			if isLower(p, adj) {
				lp = append(lp, p)
			}
		}
	}

	fmt.Println(lp)

	sum := 0
	for _, v := range lp {
		sum += v + 1
	}
	return sum
}

func part2(data []string) int {
	points := [][]int{}

	for _, p := range data {
		r := strings.Split(p, "")
		rp := []int{}
		for _, point := range r {
			val, _ := strconv.Atoi(point)
			rp = append(rp, val)
		}
		points = append(points, rp)
	}

	return 0
}
