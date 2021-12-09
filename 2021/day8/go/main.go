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
	content := strings.Split(string(input), "\n")
	result := part1(content)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(content)
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(data []string) int {
	result := 0

	for _, v := range data {
		line := strings.Split(v, " | ")
		outputs := strings.Split(line[1], " ")
		for _, o := range outputs {
			switch len(o) {
			case 2, 3, 4, 7:
				result++
			}
		}
	}

	return result
}

func alpha(s string) string {
	runes := strings.Split(s, "")
	sort.Strings(runes)
	return strings.Join(runes, "")
}

func part2(data []string) int {
	result := 0

	for _, v := range data {
		line := strings.Split(v, " | ")
		signals := strings.Split(line[0], " ")
		outputs := strings.Split(line[1], " ")
		patterns := buildPatternKnowlege(signals)
		output := ""
		for _, o := range outputs {
			number := getNumber(patterns, alpha(o))
			output = fmt.Sprintf("%v%v", output, number)
		}
		value, _ := strconv.Atoi(output)
		result += value
	}

	return result
}

func getNumber(patterns map[int]string, s string) int {
	for k, v := range patterns {
		if v == s {
			return k
		}
	}
	return -1
}

func buildPatternKnowlege(signals []string) map[int]string {
	patterns := make(map[int]string)

	for _, o := range signals {
		switch len(o) {
		case 2:
			patterns[1] = alpha(o)
		case 3:
			patterns[7] = alpha(o)
		case 4:
			patterns[4] = alpha(o)
		case 7:
			patterns[8] = alpha(o)
		}
	}

	for _, sig := range signals {
		sig = alpha(sig)
		//fmt.Printf("%v %v = %v\n", patterns[8], sig, diff(patterns[8], sig))
		if len(patterns[8])-len(sig) == 1 {
			// 0, 6, 9
			if len(contains(sig, patterns[4])) == 4 {
				patterns[9] = sig
			} else if len(contains(sig, patterns[1])) == 2 {
				patterns[0] = sig
			} else {
				patterns[6] = sig
			}
		} else if len(patterns[8])-len(sig) == 2 {
			// 2, 3, 5
			if len(contains(sig, patterns[1])) == 2 {
				patterns[3] = sig
			} else if len(contains(sig, patterns[4])) == 2 {
				patterns[2] = sig
			} else {
				patterns[5] = sig
			}
		}
	}

	return patterns
}

func contains(s1 string, s2 string) []string {
	contains := []string{}
	for _, v1 := range strings.Split(s1, "") {
		for _, v2 := range strings.Split(s2, "") {
			if v1 == v2 {
				contains = append(contains, v1)
			}
		}
	}
	return contains
}
