package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("unable to read input.txt %s", err)
		os.Exit(1)
	}

	claims := strings.Split(string(content), "\n")
	part1(claims)
	part2()
}

func part1(claims []string) {

}

func part2() {

}
