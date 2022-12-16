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

	fmt.Println("part1: ", part1(f, 2000000))

	f.Seek(0, io.SeekStart)
	fmt.Println("part2: ", part2(f, 4000000))
}

type Sensor struct {
	Pos   [2]int
	Power int
}

func (s *Sensor) Touches(x, y int) bool {
	sx, sy := s.Pos[0], s.Pos[1]

	distance := dist(sx, sy, x, y)

	if distance <= s.Power {
		return true
	}

	return false
}

func parseLine(l string) (*Sensor, [2]int) {
	s := &Sensor{}

	splits := strings.Split(l, "=")

	sx, sy := splits[1][:strings.IndexRune(splits[1], ',')], splits[2][:strings.IndexRune(splits[2], ':')]

	bx, by := splits[3][:strings.IndexRune(splits[3], ',')], splits[4]

	sX, _ := strconv.Atoi(sx)
	sY, _ := strconv.Atoi(sy)
	bX, _ := strconv.Atoi(bx)
	bY, _ := strconv.Atoi(by)

	power := dist(sX, sY, bX, bY)

	s.Power = power
	s.Pos = [2]int{sX, sY}

	return s, [2]int{bX, bY}
}

func dist(x, y, x2, y2 int) int {
	dx := 0
	if x > x2 {
		dx = x - x2
	} else {
		dx = x2 - x
	}

	dy := 0
	if y > y2 {
		dy = y - y2
	} else {
		dy = y2 - y
	}

	return dx + dy
}

func part1(input io.Reader, y int) int {
	scanner := bufio.NewScanner(input)

	sensors := []*Sensor{}
	beaconPositions := [][2]int{}
	for scanner.Scan() {
		sensor, beaconPos := parseLine(scanner.Text())
		sensors = append(sensors, sensor)
		beaconPositions = append(beaconPositions, beaconPos)
	}

	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	maxPow := math.MinInt

	for _, sensor := range sensors {
		x := sensor.Pos[0]
		y := sensor.Pos[1]
		pow := sensor.Power

		if pow > maxPow {
			maxPow = pow
		}

		if x < minX {
			minX = x
		}

		if x > maxX {
			maxX = x
		}

		if y < minY {
			minY = y
		}

		if y > maxY {
			maxY = y
		}
	}

	count := 0
	for x := minX - maxPow; x <= maxX+maxPow; x++ {
		touched := false
		for _, s := range sensors {
			if s.Touches(x, y) {
				touched = true
			}
		}

		isBeacon := false
		for _, b := range beaconPositions {
			if b == [2]int{x, y} {
				isBeacon = true
			}
		}

		if touched && !isBeacon {
			count++
		}
	}

	return count
}

func part2(input io.Reader, max int) int {
	scanner := bufio.NewScanner(input)

	beaconPositions := map[[2]int]bool{}
	sensors := []*Sensor{}
	for scanner.Scan() {
		sensor, beaconPos := parseLine(scanner.Text())
		beaconPositions[beaconPos] = true
		sensors = append(sensors, sensor)
	}

	for y := 0; y <= max; y++ {
		for x := 0; x <= max; x++ {
			var touched *Sensor

			for _, s := range sensors {
				if s.Touches(x, y) {
					touched = s
				}
			}

			_, isBeacon := beaconPositions[[2]int{x, y}]

			if touched == nil && !isBeacon {
				return x*4000000 + y
			} else {
				distance := dist(touched.Pos[0], touched.Pos[1], x, y)

				diff := touched.Power - distance

				x += diff
			}
		}
	}

	return 0
}
