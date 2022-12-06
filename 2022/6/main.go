package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

type item string

func readInput() string {
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

	return inputs[0]
}

func SolvePart1(s string) {
	var i int = 3
loop:
	for {
		switch {
		case s[i] == s[i-1]:
			i += 3
		case s[i] == s[i-2] || s[i-1] == s[i-2]:
			i += 2
		case s[i] == s[i-3] || s[i-1] == s[i-3] || s[i-2] == s[i-3]:
			i += 1
		default:
			break loop
		}
	}
	fmt.Println("Part 1: Result", i+1)
}

func SolvePart1New(s string) {
	var c int = 4
	var i int = 0
	var n int = 0
	for n < c-1 {
		idx := strings.LastIndex(s[i+1:i+c-n], string(s[i]))
		if idx == -1 {
			n++
			i++
		} else {
			n = 0
			i += idx + 1
		}
	}
	fmt.Println("Part 1: Result", i+1)
}

func SolvePart2(s string) {
	var c int = 14
	var i int = 0
	var n int = 0
	for n < c-1 {
		idx := strings.LastIndex(s[i+1:i+c-n], string(s[i]))
		if idx == -1 {
			n++
			i++
		} else {
			n = 0
			i += idx + 1
		}
	}
	fmt.Println("Part 2: Result", i+1)
}

func SolvePart2Alternative(b []byte) {
	var c int = 14
	var i int = 0
	var n int = 0
	for n < c-1 {
		idx := slices.Index(b[i+1:i+c-n], b[i])
		if idx == -1 {
			n++
			i++
		} else {
			n = 0
			i += idx + 1
		}
	}
	fmt.Println("Part 2: Result", i+1)
}

func main() {
	input := readInput()
	SolvePart1(input)
	SolvePart1New(input)
	SolvePart2(input)
	SolvePart2Alternative([]byte(input))
}
