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
		fmt.Printf("unable to read da input.txt %s", err)
		os.Exit(1)
	}

	ids := strings.Split(string(content), "\n")

	part1(ids)
	part2(ids)
}

func part1(ids []string) {
	twoCount := 0
	threeCount := 0

	for _, s := range ids {
		s = strings.ReplaceAll(s, " ", "")
		counts := map[string]int{}
		for _, r := range s {
			c := fmt.Sprintf("%c", r)
			if _, ok := counts[c]; ok {
				counts[c] = counts[c] + 1
				continue
			}
			counts[c] = 1
		}

		for _, v := range counts {
			if v == 2 {
				twoCount++
				break
			}
		}

		for _, v := range counts {
			if v == 3 {
				threeCount++
				break
			}
		}
	}

	checksum := twoCount * threeCount
	fmt.Println(checksum)
}

func part2(ids []string) {
	words := map[string][]string{}
	for _, s := range ids {
		words[s] = strings.Split(s, "")
	}

	var id1 string
	var id2 string
	found := false

	for k, v := range words {
		for k1, v1 := range words {
			if k == k1 {
				// skip when it's the same key
				continue
			}

			differs := 0
			for i := 0; i < len(v); i++ {
				if v[i] != v1[i] {
					differs++
				}
			}

			if differs == 1 {
				id1 = k
				id2 = k1
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	id := ""

	if id1 != "" && id2 != "" {
		for i := 0; i < len(id1); i++ {
			if id1[i] != id2[i] {
				id = id1[0:i] + id1[i+1:]
			}
		}
	}
	fmt.Println(id1)
	fmt.Println(id2)
	fmt.Println(id)
}
