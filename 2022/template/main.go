package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputs := readInput()
	problem1(inputs)
}

func problem1(inputs []string) {
	// parse inputs
	for _, v := range inputs {
		fmt.Println(v)
	}
}

func readInput() []string {
	file, err := os.Open("input.txt")
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
