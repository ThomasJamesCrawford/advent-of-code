package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	scanner.Scan()

	line := scanner.Text()

	for i := 4; i < len(line); i++ {
		dedupe := map[rune]bool{}
		for _, r := range line[i-4 : i] {
			dedupe[r] = true
		}

		if len(dedupe) == 4 {
			return i
		}
	}

	return 0
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	scanner.Scan()

	line := scanner.Text()

	for i := 14; i < len(line); i++ {
		dedupe := map[rune]bool{}
		for _, r := range line[i-14 : i] {
			dedupe[r] = true
		}

		if len(dedupe) == 14 {
			return i
		}
	}

	return 0
}
