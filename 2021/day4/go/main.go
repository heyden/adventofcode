package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// I am not happy with all of the nested looping
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

type board struct {
	rows [][]int
	won  bool
}

func (b *board) wins(number int) bool {
	has := false
	for y, row := range b.rows {
		for x, v := range row {
			if v == number {
				b.rows[y][x] = -1
				has = true
				break
			}
		}
	}

	if has {
		for _, row := range b.rows {
			sum := 0
			for _, v := range row {
				sum += v
			}
			if sum == -5 {
				// winner
				return true
			}
		}
		for x := 0; x < len(b.rows); x++ {
			sum := 0
			for y := 0; y < len(b.rows); y++ {
				sum += b.rows[y][x]
			}
			if sum == -5 {
				//winner
				return true
			}
		}
	}

	return false
}

func (b *board) sum() int {
	sum := 0
	for _, row := range b.rows {
		for _, v := range row {
			if v != -1 {
				sum += v
			}
		}
	}
	return sum
}

func part1(data []string) int {
	numbers := []int{}
	boards := []board{}

	// get the called numbers
	for _, v := range strings.Split(data[0], ",") {
		n, _ := strconv.Atoi(v)
		numbers = append(numbers, n)
	}

	// create board structs
	for i := 2; i < len(data); i += 6 {
		b := board{}
		for j := i; j < i+5; j++ {
			row := strings.Trim(data[j], " ")
			numbers := []int{}
			for _, v := range strings.Split(row, " ") {
				if v == "" {
					continue
				}
				number, _ := strconv.Atoi(v)
				numbers = append(numbers, number)
			}
			b.rows = append(b.rows, numbers)
		}
		boards = append(boards, b)
	}

	// call each number and check for a winner
	// if a winner is found, record the sum of the unmarked numbers
	ws := -1
	wn := -1
	for _, n := range numbers {
		for _, b := range boards {
			// did we have a winner?
			if b.wins(n) {
				ws = b.sum()
				wn = n
				break
			}
		}
		if ws >= 0 {
			break
		}
	}

	return ws * wn
}

func part2(data []string) int {
	numbers := []int{}
	boards := []*board{}

	// get the called numbers
	for _, v := range strings.Split(data[0], ",") {
		n, _ := strconv.Atoi(v)
		numbers = append(numbers, n)
	}

	// create board structs
	for i := 2; i < len(data); i += 6 {
		b := &board{}
		for j := i; j < i+5; j++ {
			row := strings.Trim(data[j], " ")
			numbers := []int{}
			for _, v := range strings.Split(row, " ") {
				if v == "" {
					continue
				}
				number, _ := strconv.Atoi(v)
				numbers = append(numbers, number)
			}
			b.rows = append(b.rows, numbers)
		}
		boards = append(boards, b)
	}

	// call each number and check for a winner
	// if a winner is found, record the sum of the unmarked numbers
	ws := -1
	wn := -1
	winners := 0
	for _, n := range numbers {
		for i, b := range boards {
			// did we have a winner?
			if !b.won && b.wins(n) {
				winners++
				b.won = true
				if winners == len(boards) {
					ws = b.sum()
					wn = n
					break
				}
			}
		}
		if ws >= 0 {
			break
		}
	}

	return ws * wn
}
