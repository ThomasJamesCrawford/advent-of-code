package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")

	defer f.Close()

	fmt.Println("part1: ", part1(f))

	f.Seek(0, io.SeekStart)
	fmt.Println("part2: ", part2(f))
}

type Assignment struct {
	lower int
	upper int
}

func (a *Assignment) contains(b Assignment) bool {
	return a.lower <= b.lower && a.upper >= b.upper
}

func (a *Assignment) intersects(b Assignment) bool {
	return a.upper >= b.lower && b.upper >= a.lower
}

func part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	sum := 0
	for scanner.Scan() {
		a, b := parseLine(scanner.Text())

		if a.contains(b) || b.contains(a) {
			sum++
		}
	}

	return sum
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	sum := 0
	for scanner.Scan() {
		a, b := parseLine(scanner.Text())

		if a.intersects(b) {
			sum++
		}
	}

	return sum
}

func parseAssignment(s string) Assignment {
	lower, upper, _ := strings.Cut(s, "-")

	upperInt, _ := strconv.Atoi(upper)
	lowerInt, _ := strconv.Atoi(lower)

	return Assignment{upper: upperInt, lower: lowerInt}
}

func parseLine(s string) (Assignment, Assignment) {
	a, b, _ := strings.Cut(s, ",")

	return parseAssignment(a), parseAssignment(b)
}
