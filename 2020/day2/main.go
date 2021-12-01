package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	r = regexp.MustCompile(`^(?P<min>\d+)-(?P<max>\d+) (?P<char>\w): (?P<password>[a-zA-Z]+)$`)
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading input file %v\n", err)
		os.Exit(1)
	}

	// transform input to expected type
	passwords := strings.Split(string(input), "\n")

	result := part1(passwords)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(passwords)
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(passwords []string) int {
	valid := 0

	for _, passwordPolicy := range passwords {
		matches := r.FindStringSubmatch(passwordPolicy)

		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		character := matches[3]
		password := matches[4]

		cr := regexp.MustCompile(fmt.Sprintf("(%v)", character))
		c := len(cr.FindAll([]byte(password), -1))
		if c >= min && c <= max {
			valid++
		}
	}

	return valid
}

func part2(passwords []string) int {
	valid := 0

	for _, passwordPolicy := range passwords {
		matches := r.FindStringSubmatch(passwordPolicy)

		pos1, _ := strconv.Atoi(matches[1])
		pos2, _ := strconv.Atoi(matches[2])
		character := matches[3]
		password := matches[4]

		count := 0

		if string(password[pos1-1]) == character {
			count++
		}

		if string(password[pos2-1]) == character {
			count++
		}

		if count == 1 {
			valid++
		}
	}

	return valid
}
