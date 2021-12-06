package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// incomplete
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

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
	t  string
}

func part1(data []string) int {
	lines := []line{}

	for _, v := range data {
		points := strings.Split(v, " -> ")
		p1 := strings.Split(points[0], ",")
		p2 := strings.Split(points[1], ",")

		x1, _ := strconv.Atoi(p1[0])
		y1, _ := strconv.Atoi(p1[1])
		x2, _ := strconv.Atoi(p2[0])
		y2, _ := strconv.Atoi(p2[1])

		t := "d"

		if x2-x1 == 0 {
			t = "v"
		} else if y2-y1 == 0 {
			t = "h"
		}

		line := line{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
			t:  t,
		}

		lines = append(lines, line)
	}

	values := make(map[string]string)
	overlaps := 0

	for _, l := range lines {
		c1xmax := int(math.Max(float64(l.x1), float64(l.x2)))
		c1xmin := int(math.Min(float64(l.x1), float64(l.x2)))
		c1ymax := int(math.Max(float64(l.y1), float64(l.y2)))
		c1ymin := int(math.Min(float64(l.y1), float64(l.y2)))

		if l.t == "h" {
			for i := c1xmin; i <= c1xmax; i++ {
				if v, ok := values[fmt.Sprintf("%v,%v", i, l.y1)]; v == "." && ok {
					overlaps++
					values[fmt.Sprintf("%v,%v", i, l.y1)] = "x"
				} else {
					values[fmt.Sprintf("%v,%v", i, l.y1)] = "."
				}
			}
		} else if l.t == "v" {
			for i := c1ymin; i <= c1ymax; i++ {
				if v, ok := values[fmt.Sprintf("%v,%v", l.x1, i)]; v == "." && ok {
					overlaps++
					values[fmt.Sprintf("%v,%v", l.x1, i)] = "x"
				} else {
					values[fmt.Sprintf("%v,%v", l.x1, i)] = "."
				}
			}
		}
	}

	return overlaps
}

func step() {

}

func part2(data []string) int {
	lines := []line{}

	for _, v := range data {
		points := strings.Split(v, " -> ")
		p1 := strings.Split(points[0], ",")
		p2 := strings.Split(points[1], ",")

		x1, _ := strconv.Atoi(p1[0])
		y1, _ := strconv.Atoi(p1[1])
		x2, _ := strconv.Atoi(p2[0])
		y2, _ := strconv.Atoi(p2[1])

		t := "d"

		if x2-x1 == 0 {
			t = "v"
		} else if y2-y1 == 0 {
			t = "h"
		}

		line := line{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
			t:  t,
		}

		lines = append(lines, line)
	}

	values := make(map[string]string)
	overlaps := 0

	for _, l := range lines {
		x := 1
		y := 1
	}

	return overlaps
}
