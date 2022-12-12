package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

type Node struct {
	Position [2]int
	Height   int
}

func (n *Node) Neighbors(values map[[2]int]*Node) []*Node {
	neighbors := []*Node{}

	dydxNeighbor := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	for _, dydx := range dydxNeighbor {
		if neighbor, ok := values[[2]int{dydx[0] + n.Position[0], dydx[1] + n.Position[1]}]; ok {
			if n.CanClimb(neighbor) {
				neighbors = append(neighbors, neighbor)
			}
		}
	}

	return neighbors
}

func (n *Node) CanClimb(n2 *Node) bool {
	if n.Height < n2.Height-1 {
		return false
	}

	return true
}

func (n *Node) BFS(g *Node, nodes map[[2]int]*Node) int {
	depths := map[int]map[[2]int]*Node{0: {n.Position: n}}
	visited := map[[2]int]bool{}

	for depth := 0; len(depths[depth]) > 0; depth++ {
		for _, node := range depths[depth] {
			visited[node.Position] = true
			for _, neighbor := range node.Neighbors(nodes) {
				if neighbor == g {
					return depth + 1
				}

				if _, ok := visited[neighbor.Position]; ok {
					continue
				}

				if _, ok := depths[depth+1]; !ok {
					depths[depth+1] = map[[2]int]*Node{neighbor.Position: neighbor}
				} else {
					depths[depth+1][neighbor.Position] = neighbor
				}
			}
		}
	}

	return -1
}

func main() {
	f, _ := os.Open("input.txt")

	defer f.Close()

	fmt.Println("part1: ", part1(f))

	f.Seek(0, io.SeekStart)
	fmt.Println("part2: ", part2(f))
}

func part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	y := 0
	nodes := map[[2]int]*Node{}
	var start *Node
	var goal *Node
	for scanner.Scan() {
		for x, c := range scanner.Text() {
			node := &Node{
				Position: [2]int{x, y},
				Height:   int(c),
			}

			if c == 'S' {
				node.Height = int('a')
				start = node
			}

			if c == 'E' {
				node.Height = int('z')
				goal = node
			}

			nodes[node.Position] = node
		}

		y++
	}

	return start.BFS(goal, nodes)
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	y := 0
	nodes := map[[2]int]*Node{}
	var goal *Node
	for scanner.Scan() {
		for x, c := range scanner.Text() {
			node := &Node{
				Position: [2]int{x, y},
				Height:   int(c),
			}

			if c == 'S' {
				node.Height = int('a')
			}

			if c == 'E' {
				node.Height = int('z')
				goal = node
			}

			nodes[node.Position] = node
		}

		y++
	}

	min := math.MaxInt

	for _, n := range nodes {
		if n.Height == 'a' {
			depth := n.BFS(goal, nodes)

			if depth != -1 && depth <= min {
				min = depth
			}
		}
	}

	return min
}
