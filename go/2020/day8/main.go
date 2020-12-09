package main

import (
	"fmt"
	"io/ioutil"
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
	inst := strings.Split(string(input), "\n")

	result := part1(inst)
	fmt.Printf("Part 1 = %v\n", result)

	result = part2(inst)
	fmt.Printf("Part 2 = %v\n", result)
}

func part1(instructions []string) int {
	b := newBootLoader(instructions)
	result, _ := b.run()
	return result
}

func part2(instructions []string) int {
	result := 0

	for i, v := range instructions {
		if strings.Contains(v, "jmp") {
			n := append([]string{}, instructions...)
			n[i] = strings.ReplaceAll(n[i], "jmp", "nop")
			b := newBootLoader(n)
			if result, terminated := b.run(); terminated {
				return result
			}
		} else if strings.Contains(v, "nop") {
			n := append([]string{}, instructions...)
			n[i] = strings.ReplaceAll(n[i], "nop", "jmp")
			b := newBootLoader(n)
			if result, terminated := b.run(); terminated {
				return result
			}
		}
	}
	return result
}

type BootLoader struct {
	// instructions
	inst []string

	// called instructions
	invoked map[string]int

	// accumulator value
	accumulator int

	// Current instruction index
	c int
}

func newBootLoader(i []string) *BootLoader {
	return &BootLoader{
		inst:    i,
		invoked: make(map[string]int),
	}
}

func (b *BootLoader) run() (int, bool) {
	for b.c < len(b.inst) {
		instruction := b.inst[b.c]

		op := strings.Split(instruction, " ")[0]
		v, _ := strconv.Atoi(strings.Split(instruction, " ")[1])

		i := fmt.Sprintf(fmt.Sprintf("%v:%v", b.c, instruction))

		if _, ok := b.invoked[i]; ok {
			return b.accumulator, false
		}

		b.invoked[i] = 1

		switch op {
		case "nop":
			b.nop(v)
		case "acc":
			b.acc(v)
		case "jmp":
			b.jmp(v)
		}
	}
	return b.accumulator, true
}

func (b *BootLoader) nop(v int) {
	b.c++
}

func (b *BootLoader) acc(v int) {
	b.accumulator += v
	b.c++
}

func (b *BootLoader) jmp(v int) {
	b.c += v
}
