package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	h int
	v bool
	s int
}

type grid [][]tree

func (g grid) String() string {
	var b strings.Builder
	for _, r := range g {
		for _, v := range r {
			b.WriteString(fmt.Sprintf("%v", v))
		}
		b.WriteString(fmt.Sprintf("\n"))
	}
	return b.String()
}

func (g grid) count() int {
	var c int
	for _, r := range g {
		for _, v := range r {
			if v.v {
				c++
			}
		}
	}
	return c
}

func (g grid) maxS() int {
	var m int
	for _, r := range g {
		for _, v := range r {
			if v.s > m {
				m = v.s
			}
		}
	}
	return m
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

func parseInput(input []string) grid {
	var g grid
	rows := len(input)
	columns := len(input[0])
	g = make([][]tree, rows)
	for i, s := range input {
		g[i] = make([]tree, columns)
		for j, v := range s {
			n, _ := strconv.Atoi(string(v))
			g[i][j] = tree{h: n, s: 1}
		}
	}
	return g
}

func SolvePart1(g grid) {
	rows := len(g)
	columns := len(g[0])

	for i := 0; i < columns; i++ {
		var m int
		var c int
		var midx int
		for j := 0; j < rows; j++ {
			if j == 0 {
				m = g[j][i].h
				g[j][i].v = true
				c++
				continue
			} else if g[j][i].h == 9 {
				g[j][i].v = true
				c++
				midx = j
				break
			} else if g[j][i].h > m {
				m = g[j][i].h
				g[j][i].v = true
				c++
				midx = j
				continue
			}
		}
		m = 0
		for j := columns - 1; j > midx; j-- {
			if j == columns-1 {
				m = g[j][i].h
				g[j][i].v = true
				c++
				continue
			} else if g[j][i].h == 9 {
				g[j][i].v = true
				c++
				break
			} else if g[j][i].h > m {
				m = g[j][i].h
				g[j][i].v = true
				c++
				continue
			}
		}
	}
	for i, r := range g {
		var m int
		var c int
		var midy int
		for j, _ := range g[i] {
			if j == 0 {
				m = r[j].h
				r[j].v = true
				c++
				continue
			} else if r[j].h == 9 {
				r[j].v = true
				c++
				midy = j
				break
			} else if r[j].h > m {
				m = r[j].h
				r[j].v = true
				c++
				midy = j
				continue
			}
		}
		m = 0
		for j := columns - 1; j > midy; j-- {
			if j == columns-1 {
				m = r[j].h
				r[j].v = true
				c++
				continue
			} else if r[j].h == 9 {
				r[j].v = true
				c++
				break
			} else if r[j].h > m {
				m = r[j].h
				r[j].v = true
				c++
				continue
			}
		}
	}
	fmt.Println("Part 1: Result", g.count())
}

func SolvePart2(g grid) {
	rows := len(g)
	columns := len(g[0])

	var dist []int = make([]int, 10)
	f := func(h int) {
		for i := range dist {
			if i <= h {
				dist[i] = 1
			} else {
				dist[i]++
			}
		}
	}

	f(0)

	for i := 0; i < rows; i++ {
		dist = make([]int, 10)
		f(0)
		for j := 0; j < columns; j++ {
			v := g[i][j]
			if j == 0 {
				g[i][j].s = 0
			} else {
				g[i][j].s = v.s * dist[v.h]
				f(v.h)
			}

		}
		dist = make([]int, 10)
		f(0)
		for j := columns - 1; j >= 0; j-- {
			v := g[i][j]
			if j == columns-1 {
				g[i][j].s = 0
			} else {
				g[i][j].s = v.s * dist[v.h]
				f(v.h)
			}
		}
	}
	dist = make([]int, 10)
	f(0)
	for i := 0; i < columns; i++ {
		dist = make([]int, 10)
		f(0)
		for j := 0; j < rows; j++ {
			v := g[j][i]
			if j == 0 {
				g[j][i].s = 0
			} else {
				g[j][i].s = v.s * dist[v.h]
				f(v.h)
			}

		}
		dist = make([]int, 10)
		f(0)
		for j := rows - 1; j >= 0; j-- {
			v := g[j][i]
			if j == rows-1 {
				g[j][i].s = 0
			} else {
				g[j][i].s = v.s * dist[v.h]
				f(v.h)
			}

		}
	}
	fmt.Println("Part 2: Result", g.maxS())
}

func main() {
	inputs := readInput()
	g := parseInput(inputs)
	SolvePart1(g)
	SolvePart2(g)
}
