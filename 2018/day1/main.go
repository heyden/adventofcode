package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("unable to read input.txt %s", err)
	}

	part1(content)
	part2(content)
}

func part1(content []byte) {
	frequency := 0

	adjustors := strings.Split(string(content), "\r\n")

	for _, s := range adjustors {
		value, _ := strconv.Atoi(s)
		frequency += value
	}
	fmt.Print(frequency)
}

func part2(content []byte) {

}
