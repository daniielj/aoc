package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type direction int8

const (
	R direction = iota
	U
	L
	D
)

func (d direction) String() string {
	switch d {
	case R:
		return "R"
	case U:
		return "U"
	case L:
		return "L"
	case D:
		return "D"
	}
	return "Error"
}

type motion struct {
	dir   direction
	steps int
}

func (m motion) String() string {
	return fmt.Sprintf("{%v %d}", m.dir, m.steps)
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

func parseInput(input []string) []motion {
	var m []motion
	for _, v := range input {
		l := strings.Split(v, " ")
		var d direction
		switch l[0] {
		case "R":
			d = R
		case "U":
			d = U
		case "L":
			d = L
		case "D":
			d = D
		default:
			fmt.Println("Error")
		}
		n, _ := strconv.Atoi(l[1])
		m = append(m, motion{d, n})
	}
	return m
}

type position struct {
	x int
	y int
}

func (p position) dist(o position) int {
	var xd int
	var yd int
	if p.x < o.x {
		xd = o.x - p.x
	} else {
		xd = p.x - o.x
	}
	if p.y < o.y {
		yd = o.y - p.y
	} else {
		yd = p.y - o.y
	}
	if xd > yd {
		return xd
	} else {
		return yd
	}
}

func (p *position) move(m motion) {
	switch m.dir {
	case R:
		p.x += m.steps
	case U:
		p.y += m.steps
	case L:
		p.x -= m.steps
	case D:
		p.y -= m.steps
	}
}

func (p *position) follow(o position) {
	if p.x < o.x && o.x-p.x > 0 {
		p.x++
	} else if p.x-o.x > 0 {
		p.x--
	}
	if p.y < o.y && o.y-p.y > 0 {
		p.y++
	} else if p.y-o.y > 0 {
		p.y--
	}
}

func SolvePart1(input []motion) {
	var posMap map[position]int = make(map[position]int)
	var tailPos position = position{}
	var headPos position = position{}
	posMap[position{0, 0}] = 1

	for _, m := range input {
		headPos.move(m)
		for headPos.dist(tailPos) > 1 {
			tailPos.follow(headPos)
			n, ok := posMap[tailPos]
			if ok {
				posMap[tailPos] = n + 1
			} else {
				posMap[tailPos] = 1
			}
		}
	}
	fmt.Println("Part 1: Result", len(posMap))
}

func SolvePart2(input []motion) {
	var n int = 9
	var posMap map[position]int = make(map[position]int)
	var posList []position = make([]position, n+1) // Head = 0
	posMap[position{0, 0}] = 1

	for _, m := range input {
		for s := 0; s < m.steps; s++ {
			posList[0].move(motion{m.dir, 1})
			for i := 1; i <= n; i++ {
				if posList[i].dist(posList[i-1]) > 1 {
					posList[i].follow(posList[i-1])
					if i == n {
						n, ok := posMap[posList[i]]
						if ok {
							posMap[posList[i]] = n + 1
						} else {
							posMap[posList[i]] = 1
						}
					}
				} else {
					break
				}
			}
		}
	}
	fmt.Println("Part 2: Result", len(posMap))
}

func main() {
	inputs := readInput()
	motions := parseInput(inputs)
	SolvePart1(motions)
	SolvePart2(motions)
}
