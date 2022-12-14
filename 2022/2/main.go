package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type (
	shape   int8
	outcome int8
)

const (
	ROCK     shape = 1
	PAPER    shape = 2
	SCISSORS shape = 3
)

const (
	LOSE outcome = iota
	DRAW
	WIN
)

type round struct {
	opponent shape
	me       shape
	outcome  outcome
}

func (r round) score() int {
	var score int = r.me.score()
	switch {
	case r.opponent == r.me:
		score += 3
	case r.me == PAPER && r.opponent == ROCK,
		r.me == ROCK && r.opponent == SCISSORS,
		r.me == SCISSORS && r.opponent == PAPER:
		score += 6
	}
	return score
}

func (s shape) score() int {
	return int(s)
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

func parseInputs(input []string) []round {
	var r []round

	o := map[byte]shape{
		'A': ROCK,
		'B': PAPER,
		'C': SCISSORS,
	}
	m := map[byte]shape{
		'X': ROCK,
		'Y': PAPER,
		'Z': SCISSORS,
	}

	for _, v := range input {
		r = append(r, round{opponent: o[v[0]], me: m[v[2]]})
	}
	return r
}

func parseInputs2(input []string) []round {
	var r []round

	c := func(b byte) shape {
		switch b {
		case 'A', 'X':
			return ROCK
		case 'B', 'Y':
			return PAPER
		case 'C', 'Z':
			return SCISSORS
		}
		return -1
	}

	for _, v := range input {
		r = append(r, round{opponent: c(v[0]), me: c(v[2])})
	}
	return r
}

func parseInputs3(input []string) []round {
	var r []round

	s := map[byte]shape{
		'A': ROCK,
		'B': PAPER,
		'C': SCISSORS,
	}
	o := map[byte]outcome{
		'X': LOSE,
		'Y': DRAW,
		'Z': WIN,
	}

	for _, v := range input {
		r = append(r, round{opponent: s[v[0]], outcome: o[v[2]]})
	}
	return r
}

func SolvePart1(input []round) {
	var sum int = 0

	for _, r := range input {
		sum += r.score()
	}
	fmt.Println("Part 1: Sum score =", sum)
}

func SolvePart2(input []round) {
	var sum int

	f := func(r round) shape {
		var me shape = r.opponent
		switch r.outcome {
		case LOSE:
			me = (r.opponent+1)%3 + 1
		case WIN:
			me = r.opponent%3 + 1
		}
		return me
	}

	for _, r := range input {
		r.me = f(r)
		sum += r.score()
	}

	fmt.Println("Part 2: Sum score =", sum)
}

func main() {
	inputs := readInput()
	rounds := parseInputs2(inputs)
	// for i, v := range rounds {
	// 	fmt.Printf("%d: %+v\n", i, v)
	// }
	SolvePart1(rounds)

	rounds2 := parseInputs3(inputs)
	// for i, v := range rounds2 {
	// 	fmt.Printf("%d: %+v\n", i, v)
	// }
	SolvePart2(rounds2)
}
