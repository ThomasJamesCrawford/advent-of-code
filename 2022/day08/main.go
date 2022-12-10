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

	trees := [][]int{}
	y := 0
	for scanner.Scan() {
		trees = append(trees, []int{})
		for _, c := range scanner.Text() {
			trees[y] = append(trees[y], int(c-'0'))
		}

		y++
	}

	visible := 0
	for y, row := range trees {
		for x, height := range row {
			if y == 0 || y == len(trees)-1 {
				visible++
				continue
			}

			if x == 0 || x == len(row)-1 {
				visible++
				continue
			}

			canSee := true
			for i := 1; i <= x; i++ {
				if height <= row[i-1] {
					canSee = false
				}
			}

			if canSee {
				visible++
				continue
			}

			canSee = true
			for i := len(row) - 2; i >= x; i-- {
				if height <= row[i+1] {
					canSee = false
				}
			}

			if canSee {
				visible++
				continue
			}

			canSee = true
			for j := 1; j <= y; j++ {
				if height <= trees[j-1][x] {
					canSee = false
				}
			}

			if canSee {
				visible++
				continue
			}

			canSee = true
			for j := len(trees) - 2; j >= y; j-- {
				if height <= trees[j+1][x] {
					canSee = false
				}
			}

			if canSee {
				visible++
				continue
			}
		}
	}

	return visible
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	trees := [][]int{}
	y := 0
	for scanner.Scan() {
		trees = append(trees, []int{})
		for _, c := range scanner.Text() {
			trees[y] = append(trees[y], int(c-'0'))
		}

		y++
	}

	max := 0
	for y, row := range trees {
		for x, height := range row {
			score := 0

			left := 0
			for i := x - 1; i >= 0; i-- {
				left++
				if height <= row[i] {
					break
				}
			}

			right := 0
			for i := x + 1; i < len(row); i++ {
				right++
				if height <= row[i] {
					break
				}
			}

			up := 0
			for j := y - 1; j >= 0; j-- {
				up++
				if height <= trees[j][x] {
					break
				}
			}

			down := 0
			for j := y + 1; j < len(trees); j++ {
				down++
				if height <= trees[j][x] {
					break
				}
			}

			score = left * right * up * down

			if score >= max {
				max = score
			}
		}
	}

	return max
}
