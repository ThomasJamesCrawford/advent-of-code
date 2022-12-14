package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	//"time"
)

func main() {
	f, _ := os.Open("input.txt")

	defer f.Close()

	fmt.Println("part1: ", part1(f))

	f.Seek(0, io.SeekStart)
	fmt.Println("part2: ", part2(f))
}

type Grid struct {
	Items [][]rune
	X     int
	Y     int
	Sand  [2]int
}

func NewGrid(x int, y int, sand [2]int) *Grid {
	g := &Grid{
		X:    x,
		Y:    y,
		Sand: sand,
	}

	for j := 0; j < y; j++ {
		g.Items = append(g.Items, []rune{})
		for i := 0; i < x; i++ {

			if i == sand[0] && j == sand[1] {
				g.Items[j] = append(g.Items[j], '+')
			} else {
				g.Items[j] = append(g.Items[j], '.')
			}
		}
	}

	return g
}

func (g *Grid) SandCount() int {
	sum := 0
	for y := 0; y < g.Y; y++ {
		for x := 0; x < g.X; x++ {
			if g.Items[y][x] == 'o' {
				sum += 1
			}
		}
	}

	return sum
}

func (g *Grid) String() string {
	res := ""
	for y := 0; y < len(g.Items); y++ {
		for x := 0; x < len(g.Items[0]); x++ {
			res += fmt.Sprintf("%c", g.Items[y][x])
		}
		res += "\n"
	}

	return res
}

func (g *Grid) AddBlock(b [2]int) {
	g.Items[b[1]][b[0]] = '#'
}

func (g *Grid) MoveSand(x, y int) bool {
	//time.Sleep(time.Millisecond * 50)

	if y == g.Y-1 {
		fmt.Println("HIT BOTTOM")
		return true
	}

	if g.Items[y+1][x] == '.' {
		g.Items[y][x] = '.'
		y = y + 1
		g.Items[y][x] = 'o'
	} else if x-1 < 0 {
		g.Items[y][x] = '.'
		return true
	} else if g.Items[y+1][x-1] == '.' {
		g.Items[y][x] = '.'
		x, y = x-1, y+1
		g.Items[y][x] = 'o'
	} else if x >= g.X {
		g.Items[y][x] = '.'
		return true
	} else if g.Items[y+1][x+1] == '.' {
		g.Items[y][x] = '.'
		x, y = x+1, y+1
		g.Items[y][x] = 'o'
	} else if g.Sand[0] == x && g.Sand[1] == y {
		return false
	} else {
		g.Items[y][x] = 'o'
		return g.MoveSand(g.Sand[0], g.Sand[1])
	}

	return g.MoveSand(x, y)
}

func (g *Grid) MoveSandPart2(x, y int) bool {
	if g.Items[y+1][x] == '.' {
		g.Items[y][x] = '.'
		y = y + 1
		g.Items[y][x] = 'o'
	} else if g.Items[y+1][x-1] == '.' {
		g.Items[y][x] = '.'
		x, y = x-1, y+1
		g.Items[y][x] = 'o'
	} else if g.Items[y+1][x+1] == '.' {
		g.Items[y][x] = '.'
		x, y = x+1, y+1
		g.Items[y][x] = 'o'
	} else if g.Sand[0] == x && g.Sand[1] == y {
		g.Items[y][x] = 'o'
		return true
	} else {
		g.Items[y][x] = 'o'
		return g.MoveSand(g.Sand[0], g.Sand[1])
	}

	return g.MoveSand(x, y)
}

func parseXY(s string) (int, int) {
	x, y, _ := strings.Cut(s, ",")

	xInt, _ := strconv.Atoi(x)
	yInt, _ := strconv.Atoi(y)

	return xInt, yInt
}

func minYMaxY(f [][2]int) (int, int) {
	sort.Slice(f, func(i, j int) bool {
		return f[i][1] < f[j][1]
	})

	return 0, f[len(f)-1][1]
}

func minXMaxX(f [][2]int) (int, int) {
	sort.Slice(f, func(i, j int) bool {
		return f[i][0] < f[j][0]
	})

	return f[0][0], f[len(f)-1][0]
}

func normalise(minX int, minY int, f ...[2]int) [][2]int {
	for i, square := range f {
		normalisedX, normalisedY := square[0]-minX, square[1]
		f[i] = [2]int{normalisedX, normalisedY}
	}

	return f
}

func parseLine(l string) [][2]int {
	sections := strings.Split(l, " -> ")

	filled := [][2]int{}

	for i := 0; i < len(sections)-1; i += 1 {
		sx, sy := parseXY(sections[i])
		ex, ey := parseXY(sections[i+1])

		if sx != ex {
			if sx > ex {
				for x := ex; x <= sx; x++ {
					filled = append(filled, [2]int{x, sy})
				}
			} else {
				for x := sx; x <= ex; x++ {
					filled = append(filled, [2]int{x, sy})
				}
			}
		} else if sy != ey {
			if sy > ey {
				for y := ey; y <= sy; y++ {
					filled = append(filled, [2]int{sx, y})
				}
			} else {
				for y := sy; y <= ey; y++ {
					filled = append(filled, [2]int{sx, y})
				}
			}
		}

	}

	return filled
}

func part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	filled := [][2]int{}
	for scanner.Scan() {
		filled = append(filled, parseLine(scanner.Text())...)
	}

	minX, maxX := minXMaxX(filled)
	minY, maxY := minYMaxY(filled)

	normalisedF := normalise(minX, minY, filled...)

	g := NewGrid(maxX-minX+1, maxY-minY+1, [2]int{500 - minX, 0 - minY})

	for _, block := range normalisedF {
		g.AddBlock(block)
	}

	fmt.Println(g)

	for !g.MoveSand(g.Sand[0], g.Sand[1]) {
	}

	fmt.Println(g)

	return g.SandCount()
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	filled := [][2]int{}
	for scanner.Scan() {
		filled = append(filled, parseLine(scanner.Text())...)
	}

	minX, maxX := minXMaxX(filled)
	minY, maxY := minYMaxY(filled)

	maxY += 2

    minX -= 200
    maxX += 200

	normalisedF := normalise(minX, minY, filled...)

    fmt.Println(minX, maxX)

	gridX, gridY := maxX-minX+1, maxY-minY+1

	g := NewGrid(gridX, gridY, [2]int{500 - minX, 0 - minY})

	for _, block := range normalisedF {
		g.AddBlock(block)
	}

	for i := 0; i < gridX; i++ {
		g.AddBlock([2]int{i, gridY - 1})
	}

	fmt.Println(g)

	for !g.MoveSandPart2(g.Sand[0], g.Sand[1]) {
	}

	fmt.Println(g)

	return g.SandCount()
    
}
