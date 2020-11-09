package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("unable to read input.txt %s", err)
		os.Exit(1)
	}

	adjustors := strings.Split(string(content), "\n")
	part1(adjustors)
	part2(adjustors)
}

func part1(adjustors []string) {
	frequency := 0
	for _, s := range adjustors {
		value, _ := strconv.Atoi(s)
		frequency += value
	}
	fmt.Println(frequency)
}

func part2(adjustors []string) {
	frequency := 0
	noDuplicate := true
	frequencies := map[int]int{}
	var duplicate int

	for noDuplicate == true {
		for _, s := range adjustors {
			value, _ := strconv.Atoi(s)
			frequency += value

			if _, ok := frequencies[frequency]; ok {
				duplicate = frequency
				noDuplicate = false
				break
			} else {
				frequencies[frequency] = 1
			}
		}
	}

	fmt.Println(duplicate)
}
