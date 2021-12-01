package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
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
	data := strings.Split(string(input), "\n")
	result := part1(data)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(data[1])
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(input []string) int {
	timestamp, _ := strconv.Atoi(input[0])

	busIDs := strings.Split(input[1], ",")
	closestTimestamp := 0
	closestBusID := 0

	for _, d := range busIDs {
		busID := string(d)
		if busID == "x" {
			continue
		}

		interval, _ := strconv.Atoi(busID)

		i := 0
		nearest := 0

		for {
			i += interval

			if i >= timestamp {
				nearest = i
				d1 := int(math.Abs(float64(timestamp - closestTimestamp)))
				d2 := int(math.Abs(float64(timestamp - nearest)))
				if d1 > d2 {
					closestTimestamp = nearest
					closestBusID = interval
				}
				break
			}
		}
	}
	return (closestTimestamp - timestamp) * closestBusID
}

func part2(input string) int {
	data := strings.Split(input, ",")

	busIDs := []int{}

	for _, busID := range data {
		v := 0
		if busID != "x" {
			v, _ = strconv.Atoi(busID)
		}
		busIDs = append(busIDs, v)
	}

	t := 0

	return t
}

/**
Brute force way

t := 0
for {
	found := true
	t += busIDs[0]

	for i, v := range busIDs {
		if v == 0 {
			continue
		}
		d := i
		if (t+d)%v != 0 {
			found = false
			break
		}
	}
	if found {
		break
	}
}
*/
