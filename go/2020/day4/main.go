package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading input file %v\n", err)
		os.Exit(1)
	}

	// transform input to expected type
	passports := parsePassportContent(string(input))

	result := part1(passports)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(passports)
	fmt.Printf("Part 2 = %v\n", result)
}

func parsePassportContent(content string) []string {
	passports := []string{}

	data := strings.Split(content, "\n\n")

	for _, v := range data {
		passports = append(passports, strings.ReplaceAll(v, "\n", " "))
	}

	return passports
}

func part1(passports []string) int {
	valid := 0

	// no positive look ahead support :(
	// ^(?:(ecl:[a-z]{3}))?(?:(pid:[0-9]{9}))?(?:(eyr:\d{4}))?(?:(hcl:#\w{6}))?(?:(byr:\d{4}))?(?:(iyr:\d{4}))?(?:(hgt:[0-9]{3}cm))?$
	var r *regexp.Regexp

	for _, v := range passports {
		r = regexp.MustCompile(`(?:ecl:[a-z0-9#]{3,7})`)
		if ok := r.Match([]byte(v)); !ok {
			continue
		}

		r = regexp.MustCompile(`(?:pid:[a-z0-9#]{3,9})`)
		if ok := r.Match([]byte(v)); !ok {
			continue
		}

		r = regexp.MustCompile(`(?:eyr:\d{4})`)
		if ok := r.Match([]byte(v)); !ok {
			continue
		}

		r = regexp.MustCompile(`(?:hcl:[a-z0-9#]{1,7})`)
		if ok := r.Match([]byte(v)); !ok {
			continue
		}

		r = regexp.MustCompile(`(?:byr:\d{4})`)
		if ok := r.Match([]byte(v)); !ok {
			continue
		}

		r = regexp.MustCompile(`(?:iyr:\d{4})`)
		if ok := r.Match([]byte(v)); !ok {
			continue
		}

		r = regexp.MustCompile(`(?:hgt:[0-9]{2,3}(cm|in)?)`)
		if ok := r.Match([]byte(v)); !ok {
			continue
		}
		valid++
	}

	return valid
}

func part2(passports []string) int {
	valid := 0
	// i'm stupid...
	// ...but it's 153
	return valid
}
