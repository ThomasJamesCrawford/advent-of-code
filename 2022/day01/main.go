package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
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

	max := 0
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if sum > max {
				max = sum
			}

			sum = 0

			continue
		}

		n, _ := strconv.Atoi(line)

		sum += n
	}

	if sum > max {
		max = sum
	}

	return max
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	elves := []int{}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elves = append(elves, sum)

			sum = 0

			continue
		}

		n, _ := strconv.Atoi(line)

		sum += n
	}

    elves = append(elves, sum)

	sort.Ints(elves)

	return elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
}
