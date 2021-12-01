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
	s := strings.Split(string(input), "\n")

	s1 := make([][]string, 0)
	s2 := make([][]string, 0)

	for _, v := range s {
		s1 = append(s1, strings.Split(v, ""))
		s2 = append(s2, strings.Split(v, ""))
	}

	result := part1(s1)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(s2)
	fmt.Printf("Part 2 = %v\n", result)
}

func printSeats(s [][]string) {
	fmt.Println("---")
	for _, r := range s {
		fmt.Println(r)
	}
}

func part1(s [][]string) int {
	for {
		count := board(s, true)
		if count == 0 {
			break
		}
	}

	f := 0

	for _, r := range s {
		for _, s := range r {
			if s == "#" {
				f++
			}
		}
	}

	return f
}

func part2(s [][]string) int {
	for {
		count := board(s, false)
		if count == 0 {
			break
		}
	}

	f := 0

	for _, r := range s {
		for _, s := range r {
			if s == "#" {
				f++
			}
		}
	}

	return f
}

func board(s [][]string, adjacent bool) int {
	var c []string

	for y, r := range s {
		for x, seat := range r {
			if seat == "." {
				continue
			}

			if adjacent {
				a := getAdjacencies(x, y)
				if state := sit(s, seat, a, 4); state != "" {
					c = append(c, fmt.Sprintf("%v-%v-%v", x, y, state))
				}
			} else {
				a := getNearestSeen(s, x, y)
				if state := sit(s, seat, a, 5); state != "" {
					c = append(c, fmt.Sprintf("%v-%v-%v", x, y, state))
				}
			}
		}
	}

	for _, change := range c {
		newState := strings.Split(change, "-")
		x, _ := strconv.Atoi(newState[0])
		y, _ := strconv.Atoi(newState[1])
		state := newState[2]
		s[y][x] = state
	}

	return len(c)
}

func sit(seats [][]string, v string, a [][]int, max int) string {
	occupied := 0
	for _, adjSeat := range a {
		x := adjSeat[0]
		y := adjSeat[1]

		if x < 0 || y < 0 {
			continue
		}
		if y >= len(seats) || x >= len(seats[0]) {
			continue
		}

		if o := seats[y][x]; o == "#" {
			occupied++
		}
	}

	if v == "#" && occupied >= max {
		return "L"
	} else if v == "L" && occupied == 0 {
		return "#"
	}

	return ""
}

func getAdjacencies(x, y int) [][]int {
	a := [][]int{}
	for ax := x - 1; ax <= x+1; ax++ {
		for ay := y - 1; ay <= y+1; ay++ {
			if x == ax && y == ay {
				continue
			}
			a = append(a, []int{ax, ay})
		}
	}
	return a
}

func getNearestSeen(seats [][]string, x, y int) [][]int {
	d := [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
	}

	n := [][]int{}

	for _, m := range d {
		ax := x
		ay := y
		for {
			ax += m[0]
			ay += m[1]
			if ax >= len(seats[0]) || ay >= len(seats) || ax < 0 || ay < 0 {
				break
			}
			s := seats[ay][ax]
			if s == "L" || s == "#" {
				n = append(n, []int{ax, ay})
				break
			}
		}
	}

	return n
}
