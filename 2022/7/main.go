package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type dir struct {
	name  string
	dirs  map[string]dir
	files map[string]file
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

func (d dir) addDir(n string) {
	d.dirs[n] = dir{n, make(map[string]dir), make(map[string]file)}
}

func (d dir) addFile(n string, size int) {
	d.files[n] = file{n, size}
}

func parseInput(input []string) dir {
	var root dir = dir{"/", make(map[string]dir), make(map[string]file)}
	var s []dir = []dir{root}
	for _, v := range input {
		switch {
		case strings.HasPrefix(v, "$ ls"):
			continue
		case strings.HasPrefix(v, "$ cd /"):
			s = []dir{root}
		case strings.HasPrefix(v, "$ cd .."):
			s = s[:len(s)-1]
		case strings.HasPrefix(v, "$ cd "):
			s = append(s, s[len(s)-1].dirs[strings.TrimPrefix(v, "$ cd ")])
		case strings.HasPrefix(v, "dir "):
			s[len(s)-1].addDir(strings.TrimPrefix(v, "dir "))
		default:
			l := strings.Split(v, " ")
			n := l[1]
			sz, _ := strconv.Atoi(l[0])
			s[len(s)-1].addFile(n, sz)
		}
	}
	return root
}

// Recursive function
func checkSize(d dir) (int, int) {
	var t int
	var sum int
	for _, d := range d.dirs {
		dsum, tt := checkSize(d)
		sum += dsum
		t += tt
	}
	for _, f := range d.files {
		sum += f.size
	}
	if sum <= 100000 {
		t += sum
	}
	return sum, t
}

// Recursive function
func calcSize(d dir, sizes *[]int) int {
	var sum int
	for _, d := range d.dirs {
		sum += calcSize(d, sizes)
	}
	for _, f := range d.files {
		sum += f.size
	}
	*sizes = append(*sizes, sum)
	return sum
}

func SolvePart1(d dir) {
	_, t := checkSize(d)
	fmt.Println("Part 1: Result", t)
}

func SolvePart1New(d dir) {
	var sizes []int
	_ = calcSize(d, &sizes)

	var sum int
	for _, v := range sizes {
		if v < 100000 {
			sum += v
		}
	}

	fmt.Println("Part 1 new: Result", sum)
}

func SolvePart2(d dir) {
	var sizes []int
	s := calcSize(d, &sizes)

	a := 70000000 - s
	if a < 0 {
		panic("Unexpected negative amount")
	}

	l := 30000000 - a
	if l < 0 {
		panic("Unexpected negative limit")
	}

	sort.Ints(sizes)

	var sz int
	for _, v := range sizes {
		if v > l {
			sz = v
			break
		}
	}

	fmt.Println("Part 2: Result", sz)
}

func main() {
	inputs := readInput()
	dir := parseInput(inputs)
	// fmt.Printf("Input: %+v\n", dir)
	SolvePart1(dir)
	SolvePart1New(dir)
	SolvePart2(dir)
}
