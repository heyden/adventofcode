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

	r := part2(content)
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(data []string) int {
	gr := ""
	er := ""

	i := 0
	for i < len(data[0]) {
		c := []int{0, 0}
		for _, v := range data {
			switch string(v[i]) {
			case "0":
				c[0] = c[0] + 1
			case "1":
				c[1] = c[1] + 1
			}
		}

		if c[0] < c[1] {
			gr = gr + "1"
			er = er + "0"
		} else {
			gr = gr + "0"
			er = er + "1"
		}
		i++
	}

	grv, _ := strconv.ParseInt(gr, 2, 64)
	erv, _ := strconv.ParseInt(er, 2, 64)
	return int(grv * erv)
}

func part2(data []string) int {
	ogr := 0
	csr := 0

	return ogr * csr
}
