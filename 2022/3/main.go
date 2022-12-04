package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/slices"
)

type item byte

type rucksack struct {
	items []item
	c1    []item
	c2    []item
}

func (b item) prio() int {
	if b >= 'a' {
		return int(b-'a') + 1
	} else {
		return int(b-'A') + 27
	}
}

func readInput() []string {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputs []string

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		inputs = append(inputs, scan.Text())
	}

	return inputs
}

func parseInputs(input []string) []rucksack {
	var l []rucksack
	for _, v := range input {
		n := len(v)
		m := n / 2
		l = append(l, rucksack{[]item(v), []item(v)[0:m], []item(v)[m:n]})
	}
	return l
}

func SolvePart1(input []rucksack) {
	var sum int

	for _, v := range input {
		for _, b := range v.c1 {
			if slices.Contains(v.c2, b) {
				sum += b.prio()
				break
			}
		}
	}
	fmt.Println("Part 1: Sum = ", sum)
}

func SolvePart2(input []rucksack) {
	var sum int

	var groupSize int = 3
	for i := 0; i < len(input); i += groupSize {
		for _, b := range input[i].items {
			if slices.Contains(input[i+1].items, b) &&
				slices.Contains(input[i+2].items, b) {
				sum += b.prio()
				break
			}
		}
	}
	fmt.Println("Part 2: Sum = ", sum)
}

func main() {
	inputs := readInput()
	rucksacks := parseInputs(inputs)
	for i, v := range rucksacks {
		fmt.Printf("%d: %+v\n", i, v)
	}
	SolvePart1(rucksacks)
	SolvePart2(rucksacks)
}
