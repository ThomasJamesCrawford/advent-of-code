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
	fmt.Println(part2(f))
}

func runCycle(cycle int) bool {
	return cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220
}

func part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	r := 1
	cycle := 1
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		command, val, _ := strings.Cut(line, " ")

		if command == "addx" {
			valInt, _ := strconv.Atoi(val)

			if ok := runCycle(cycle); ok {
				sum += r * cycle
			}

			cycle++

			if ok := runCycle(cycle); ok {
				sum += r * cycle
			}

			cycle++
			r += valInt
		} else {
			if ok := runCycle(cycle); ok {
				sum += r * cycle
			}

			cycle++
		}
	}

	return sum
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func drawSprite(cycle int, x int) string {
	pixel := (cycle - 1) % 40

	nl := false
	if pixel == 39 {
		nl = true
	}

	ret := ""

	if abs(pixel, x) < 2 {
		ret += "#"
	} else {
		ret += "."
	}

	if nl {
		return ret + "\n"
	}

	return ret
}

func part2(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	x := 1
	cycle := 1
	res := ""
	for scanner.Scan() {
		line := scanner.Text()
		command, val, _ := strings.Cut(line, " ")

		if command == "addx" {
			valInt, _ := strconv.Atoi(val)

			res += drawSprite(cycle, x)

			cycle++

			res += drawSprite(cycle, x)

			x += valInt

			cycle++
		} else {
			res += drawSprite(cycle, x)

			cycle++
		}
	}

	fmt.Println()

	return res
}
