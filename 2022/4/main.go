package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sRange struct {
	start int
	end   int
}

type pair struct {
	first  sRange
	second sRange
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

func parseInputs(input []string) []pair {
	var l []pair

	toInt := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}

	for _, s := range input {
		sl := strings.Split(s, ",")
		s1l := strings.Split(sl[0], "-")
		s2l := strings.Split(sl[1], "-")
		l = append(l,
			pair{
				sRange{toInt(s1l[0]), toInt(s1l[1])},
				sRange{toInt(s2l[0]), toInt(s2l[1])},
			},
		)
	}
	return l
}

func SolvePart1(inputs []pair) {
	var count int
	for _, p := range inputs {
		if (p.first.start >= p.second.start && p.first.end <= p.second.end) ||
			(p.second.start >= p.first.start && p.second.end <= p.first.end) {
			count++
		}
	}
	fmt.Println("Part 1: Count: ", count)
}

func SolvePart2(inputs []pair) {
	var count int
	for _, p := range inputs {
		if (p.first.start >= p.second.start && p.first.start <= p.second.end) ||
			(p.second.start >= p.first.start && p.second.start <= p.first.end) {
			count++
		}
	}
	fmt.Println("Part 2: Count: ", count)
}

func main() {
	inputs := readInput()
	pairs := parseInputs(inputs)
	for _, v := range pairs {
		fmt.Printf("%v\n", v)
	}
	SolvePart1(pairs)
	SolvePart2(pairs)
}
