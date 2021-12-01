package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	E = iota + 1
	S
	W
	N
)

var (
	directions = []int{E, S, W, N}
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

type pos struct {
	x int
	y int
}

func part1(d []string) int {
	ship := &pos{x: 0, y: 0}
	direction := E

	for _, v := range d {
		op := v[0:1]
		n, _ := strconv.Atoi(v[1:])

		switch op {
		case "F":
			moved := move(direction, n)
			ship.x += moved[0]
			ship.y += moved[1]
		case "R":
			direction = turn(op, direction, n)
		case "L":
			direction = turn(op, direction, n)
		case "N":
			ship.y += n
		case "E":
			ship.x += n
		case "S":
			ship.y -= n
		case "W":
			ship.x -= n
		}
	}

	return int(math.Abs(float64(ship.x)) + math.Abs(float64(ship.y)))
}

func part2(d []string) int {
	wp := &pos{x: 10, y: 1}
	movement := []int{10, 1}
	ship := &pos{x: 0, y: 0}

	for _, v := range d {
		op := v[0:1]
		n, _ := strconv.Atoi(v[1:])

		switch op {
		case "F":
			m := make([]int, n)

			for range m {
				ship.x = wp.x
				ship.y = wp.y
				wp.x = ship.x + movement[0]
				wp.y = ship.y + movement[1]
			}
		case "R":
			movement = rotateWaypoint(wp, ship, movement, op, n)
		case "L":
			movement = rotateWaypoint(wp, ship, movement, op, n)
		case "N":
			wp.y += n
			movement[1] += n
		case "E":
			wp.x += n
			movement[0] += n
		case "S":
			wp.y -= n
			movement[1] -= n
		case "W":
			wp.x -= n
			movement[0] -= n
		}
	}

	return int(math.Abs(float64(ship.x)) + math.Abs(float64(ship.y)))
}

func move(c int, u int) []int {
	switch c {
	case N:
		return []int{0, u}
	case E:
		return []int{u, 0}
	case S:
		return []int{0, -u}
	case W:
		return []int{-u, 0}
	}
	// should never happen, ideally maybe return an error
	return nil
}

func turn(dir string, c, degrees int) int {
	turn := degrees / 90
	if dir == "L" {
		turn *= -1
	}

	var n int
	r := c + turn
	if r <= 0 {
		n = r + 4
	} else if r > 4 {
		n = r - 4
	} else {
		n = r
	}

	return n
}

func rotateWaypoint(p, s *pos, movement []int, dir string, degrees int) []int {
	turns := int(math.Abs(float64(degrees) / 90))

	t := make([]int, turns)
	m := []int{movement[0], movement[1]}
	for range t {
		m1 := m[0]
		m2 := m[1]
		if dir == "L" {
			m[0] = m2 * -1
			m[1] = m1
		} else {
			m[0] = m2
			m[1] = m1 * -1
		}
		p.x = s.x + m[0]
		p.y = s.y + m[1]
	}
	return m
}
