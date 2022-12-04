package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type item string

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

func parseInput(input []string) []item {
	var items []item
	for _, v := range input {
		items = append(items, item(v))
	}
	return items
}

func SolvePart1(inputs []item) {
	for _, v := range inputs {
		fmt.Println(v)
	}
	fmt.Println("Part 1: Result")
}

func SolvePart2(inputs []item) {
	for _, v := range inputs {
		fmt.Println(v)
	}
	fmt.Println("Part 2: Result")
}

func main() {
	inputs := readInput()
	items := parseInput(inputs)
	SolvePart1(items)
	SolvePart2(items)
}
