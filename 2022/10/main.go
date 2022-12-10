package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func SolvePart1(inputs []string) {
	var sum int
	var c int
	var x int = 1
	for _, v := range inputs {
		var n, dx int
		switch {
		case strings.HasPrefix(v, "noop"):
			n = 1
		case strings.HasPrefix(v, "addx"):
			n = 2
			dx, _ = strconv.Atoi(v[5:])

		}
		for i := 0; i < n; i++ {
			c++
			if ((c - 20) % 40) == 0 {
				sum += c * x
			}
		}
		x += dx
	}
	fmt.Println("Part 1: Result", sum)
}

func SolvePart2(inputs []string) {
	rows := 6
	columns := 40
	var crt [][]bool = make([][]bool, rows)
	var c int
	var x int = 1
	for _, v := range inputs {
		var n, dx int
		switch {
		case strings.HasPrefix(v, "noop"):
			n = 1
		case strings.HasPrefix(v, "addx"):
			n = 2
			dx, _ = strconv.Atoi(v[5:])
		}
		for i := 0; i < n; i++ {
			row := (c) / columns
			pos := c % columns
			if (pos) == 0 {
				crt[row] = make([]bool, columns)
			}
			crt[row][pos] = (x-1 <= pos && pos <= x+1)
			c++
		}
		x += dx
	}
	fmt.Println("Part 2:")
	for _, v := range crt {
		for _, b := range v {
			if b {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	inputs := readInput()
	SolvePart1(inputs)
	SolvePart2(inputs)
}
