package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func dist(a [2]int, b [2]int) float64 {
	return math.Sqrt(math.Pow(float64(a[0]-b[0]), 2) + math.Pow(float64(a[1]-b[1]), 2))
}

func tMove(hPos [2]int, tPos [2]int) [2]int {
	if dist(hPos, tPos) <= math.Sqrt2 {
		return tPos
	}

	possibleMoves := [][2]int{
		{0, 0},
		{0, 1},
		{0, -1},
		{1, 0},
		{1, 1},
		{1, -1},
		{-1, 0},
		{-1, 1},
		{-1, -1},
	}

	minDist := math.MaxFloat64
	tAfterBestMove := [2]int{0, 0}
	for _, move := range possibleMoves {
		tAfterMove := [2]int{tPos[0] + move[0], tPos[1] + move[1]}
		dist := dist(tAfterMove, hPos)

		if dist <= minDist {
			minDist = dist
			tAfterBestMove = tAfterMove
		}
	}

	return tAfterBestMove
}

func part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	visited := map[[2]int]bool{
		{0, 0}: true,
	}

	hPos := [2]int{0, 0}
	tPos := [2]int{0, 0}

	moveMap := map[string][2]int{
		"U": {0, 1},
		"D": {0, -1},
		"L": {-1, 0},
		"R": {1, 0},
	}

	for scanner.Scan() {
		move, sizeStr, _ := strings.Cut(scanner.Text(), " ")

		size, _ := strconv.Atoi(sizeStr)

		for i := 0; i < size; i++ {
			hMove := moveMap[move]

			hPos = [2]int{hPos[0] + hMove[0], hPos[1] + hMove[1]}
			tPos = tMove(hPos, tPos)

			visited[tPos] = true
		}
	}

	return len(visited)
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	visited := map[[2]int]bool{
		{0, 0}: true,
	}

	hPos := [2]int{0, 0}
	knots := [9][2]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	moveMap := map[string][2]int{
		"U": {0, 1},
		"D": {0, -1},
		"L": {-1, 0},
		"R": {1, 0},
	}

	for scanner.Scan() {
		move, sizeStr, _ := strings.Cut(scanner.Text(), " ")

		size, _ := strconv.Atoi(sizeStr)

		for i := 0; i < size; i++ {
			hMove := moveMap[move]
			hPos = [2]int{hPos[0] + hMove[0], hPos[1] + hMove[1]}

			for knotIndex, knot := range knots {
				if knotIndex == 0 {
					knots[0] = tMove(hPos, knots[0])
				} else {
					knots[knotIndex] = tMove(knots[knotIndex-1], knot)
				}
			}

			visited[knots[8]] = true
		}
	}

	return len(visited)
}
