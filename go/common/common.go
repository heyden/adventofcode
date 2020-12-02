package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ReadInput reads the input.txt file for input.
func ReadInput() []string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading input file %v\n", err)
		os.Exit(1)
	}

	content := strings.Split(string(input), "\n")
	return content
}
