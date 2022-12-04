package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")

	defer f.Close()

	fmt.Println("part1: ", part1(f))

	f.Seek(0, io.SeekStart)
	fmt.Println("part2: ", part2(f))
}

func part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	score := map[string]map[string]int{
		"A": {
			"X": 1 + 3,
			"Y": 2 + 6,
			"Z": 3 + 0,
		},
		"B": {
			"X": 1 + 0,
			"Y": 2 + 3,
			"Z": 3 + 6,
		},
		"C": {
			"X": 1 + 6,
			"Y": 2 + 0,
			"Z": 3 + 3,
		},
	}

	sum := 0
	for scanner.Scan() {
		a, b, _ := strings.Cut(scanner.Text(), " ")

		sum += score[a][b]
	}

	return sum
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	score := map[string]map[string]int{
		"A": {
			"X": 0 + 3,
			"Y": 3 + 1,
			"Z": 6 + 2,
		},
		"B": {
			"X": 0 + 1,
			"Y": 3 + 2,
			"Z": 6 + 3,
		},
		"C": {
			"X": 0 + 2,
			"Y": 3 + 3,
			"Z": 6 + 1,
		},
	}

	sum := 0
	for scanner.Scan() {
		a, b, _ := strings.Cut(scanner.Text(), " ")

		sum += score[a][b]
	}

	return sum
}
