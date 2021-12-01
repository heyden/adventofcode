package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading input file %v\n", err)
		os.Exit(1)
	}

	// transform input to expected type
	data := parseInput(string(input))

	result := part1(data)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(data)
	fmt.Printf("Part 2 = %v\n", result)
}

func parseInput(content string) [][]string {
	result := [][]string{}
	data := strings.Split(content, "\n\n")

	for _, v := range data {
		result = append(result, strings.Split(v, "\n"))
	}

	return result
}

func part1(input [][]string) int {
	result := 0

	for _, v := range input {
		var b bytes.Buffer
		for _, a := range v {
			for _, c := range a {
				if !bytes.Contains(b.Bytes(), []byte(string(c))) {
					b.WriteString(string(c))
				}
			}
		}
		result += len(b.Bytes())
	}

	return result
}

func part2(input [][]string) int {
	result := 0

	for _, s := range input {
		count := 0
		i := s[0]

		for _, c := range i {
			has := true
			for _, a := range s[1:] {
				if !strings.Contains(a, string(c)) {
					has = false
				}
			}
			if has {
				count++
			}
		}

		result += count
	}

	return result
}
