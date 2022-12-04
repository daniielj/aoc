package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	itemCalories []int
}

// Recursive function
func sum(s []int) int {
	if l := len(s); l == 0 {
		return 0
	} else {
		return sum(s[1:l]) + s[0]
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

func parseInputs(input []string) []elf {
	var e []elf = []elf{{}} // Dummy elf

	var itemCalories []int
	for _, v := range input {
		if v == "" {
			e = append(e, elf{itemCalories: itemCalories})
			itemCalories = []int{}
			continue
		}
		i, _ := strconv.Atoi(v)
		itemCalories = append(itemCalories, i)
	}
	return e
}

func SolvePart1(input []elf) {
	var max int
	var n int
	for i, e := range input {
		if s := sum(e.itemCalories); s > max {
			max = s
			n = i
		}
	}
	fmt.Printf("Part 1: Max calories = %d, elf %d\n", max, n)
}

func SolvePart2(input []elf) {
	var sums []int
	for _, e := range input {
		sums = append(sums, sum(e.itemCalories))
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})
	fmt.Println("Part 2: Sum calories top three =", sums[0]+sums[1]+sums[2])
}

func main() {
	inputs := readInput()
	elves := parseInputs(inputs)
	for i, v := range elves {
		fmt.Printf("%d: %+v\n", i, v)
	}
	SolvePart1(elves)
	SolvePart2(elves)
}
