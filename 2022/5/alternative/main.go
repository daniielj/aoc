package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack []byte

type move struct {
	count int
	from  int
	to    int
}

type inputT struct {
	stacks []stack
	moves  []move
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

func parseInput(input []string) inputT {
	var data inputT
	for _, v := range input {
		if v == "" {
			continue
		} else if strings.HasPrefix(v, "move") {
			s := strings.Split(v, " ")
			c, _ := strconv.Atoi(s[1])
			f, _ := strconv.Atoi(s[3])
			t, _ := strconv.Atoi(s[5])
			data.moves = append(data.moves, move{c, f, t})
		} else {
			var n int
			for i := 1; i < len(v); i += 4 {
				n++
				if len(data.stacks) < n {
					data.stacks = append(data.stacks, stack{})
				}
				if v[i] >= 'A' {
					data.stacks[n-1] = append(data.stacks[n-1], v[i])
				}
			}
		}
	}
	// Reverse stacks
	for i := range data.stacks {
		n := len(data.stacks[i]) - 1
		for j := 0; j < n/2; j++ {
			t := data.stacks[i][j]
			data.stacks[i][j] = data.stacks[i][n-j]
			data.stacks[i][n-j] = t
		}
	}
	return data
}

func SolvePart1(d inputT) {
	for _, v := range d.moves {
		for i := 0; i < v.count; i++ {
			l := len(d.stacks[v.from-1])
			d.stacks[v.to-1] = append(d.stacks[v.to-1], d.stacks[v.from-1][l-1])
			d.stacks[v.from-1] = d.stacks[v.from-1][:l-1]
		}
	}
	fmt.Printf("Part 1: Result = ")
	for _, v := range d.stacks {
		fmt.Printf("%c", v[len(v)-1])
	}
	fmt.Println("")
}

func SolvePart2(d inputT) {
	for _, v := range d.moves {
		l := len(d.stacks[v.from-1])
		d.stacks[v.to-1] = append(d.stacks[v.to-1], d.stacks[v.from-1][l-v.count:]...)
		d.stacks[v.from-1] = d.stacks[v.from-1][:l-v.count]
	}
	fmt.Printf("Part 2: Result = ")
	for _, v := range d.stacks {
		fmt.Printf("%c", v[len(v)-1])
	}
	fmt.Println("")
}

func main() {
	inputs := readInput()
	data := parseInput(inputs)
	SolvePart1(data)
	data = parseInput(inputs)
	SolvePart2(data)
}
