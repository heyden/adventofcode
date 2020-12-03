package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading input file %v\n", err)
		os.Exit(1)
	}

	// transform input to expected type

	part1()

	part2()
}

func part1() {

}

func part2() {

}
