package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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

type Valve struct {
	Name      string
	Rate      int
	Neighbors []string
	Open      bool
}

func parseLine(l string) Valve {
	v := Valve{}

	v.Name = l[6:8]

	v.Rate, _ = strconv.Atoi(strings.Split(l, "=")[1][:strings.IndexRune(strings.Split(l, "=")[1], ';')])

	r := strings.NewReplacer("valves", "valve")

	rSpace := strings.NewReplacer(" ", "")

	for _, name := range strings.Split(rSpace.Replace(strings.Split(r.Replace(l), "valve")[1]), ",") {
		v.Neighbors = append(v.Neighbors, name)
	}

	return v
}

func djikstra(nodes []*Node) {
	byName := map[string]*Node{}
	for _, node := range nodes {
		byName[node.Name] = node
	}

	for _, source := range nodes {

		prev := map[string]*Node{}

		dist := map[string]int{}

		q := []*Node{}

		for _, node := range nodes {
			dist[node.Name] = math.MaxInt
			q = append(q, node)
		}

		dist[source.Name] = 0

		for len(q) != 0 {
			sort.Slice(q, func(i, j int) bool {
				return dist[q[i].Name] > dist[q[j].Name]
			})

			u := q[len(q)-1]
			q = q[:len(q)-1]

			for _, name := range u.Neighbors {
				for _, inQ := range q {
					if inQ.Name == name {
						v := byName[name]

						alt := dist[u.Name] + 1

						if alt < dist[v.Name] {
							dist[v.Name] = alt
							prev[v.Name] = u
						}
						break
					}
				}
			}
		}

		source.Edges = map[string]*Edge{}

		for name, distance := range dist {
			dest := byName[name]

			if dest.Rate == 0 {
				continue
			}

			if dest.Name == source.Name {
				continue
			}

			source.Edges[dest.Name] = &Edge{
				Dist: distance,
				Dest: dest,
			}

		}
	}
}

func (v *Valve) Node() *Node {
	return &Node{
		Name:      v.Name,
		Rate:      v.Rate,
		Neighbors: v.Neighbors,
	}
}

type Edge struct {
	Dist int
	Dest *Node
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d -> %s", e.Dist, e.Dest.Name)
}

type Node struct {
	Name string
	Rate int

	Edges     map[string]*Edge
	Neighbors []string
}

func (m *Node) String() string {
	return fmt.Sprintf("%s, %d, %v", m.Name, m.Rate, m.Edges)
}

type Graph struct {
	Minutes     int
	Open        string
	TotalFlowed int
	FlowRate    int

	curr         *Node
	currOnWayRem int
	currPath     string

	curr2         *Node
	curr2OnWayRem int
	curr2Path     string
}

func (g *Graph) IsCurrClosed() bool {
	return !strings.Contains(g.Open, g.curr.Name)
}

func (g *Graph) IsCurr2Closed() bool {
	return !strings.Contains(g.Open, g.curr2.Name)
}

func (g Graph) CurrOpen() Graph {
	g.Open += g.curr.Name
	g.FlowRate += g.curr.Rate
	g.currPath += fmt.Sprintf("%s,", g.curr.Name)

	return g
}

func (g Graph) CurrOpen2() Graph {
	g.Open += g.curr2.Name
	g.FlowRate += g.curr2.Rate
	g.curr2Path += fmt.Sprintf("%s,", g.curr2.Name)

	return g
}

func (g *Graph) AdvanceMinutes(i int) *Graph {
	g.TotalFlowed += g.FlowRate * i
	g.Minutes -= i

	g.curr2OnWayRem -= i
	g.currOnWayRem -= i

	if g.currOnWayRem < 0 {
		g.currOnWayRem = 0
	}

	if g.curr2OnWayRem < 0 {
		g.curr2OnWayRem = 0
	}

	return g
}

func (g Graph) MoveTo(n *Node) Graph {
	destEdge := g.curr.Edges[n.Name]

	g.curr = destEdge.Dest
	g.currOnWayRem = destEdge.Dist

	return g
}

func (g Graph) MoveTo2(n *Node) Graph {
	destEdge := g.curr2.Edges[n.Name]

	g.curr2 = destEdge.Dest
	g.curr2OnWayRem = destEdge.Dist

	return g
}

func (g Graph) Search() int {
	if g.Minutes == 0 {
		return g.TotalFlowed
	}

	g.AdvanceMinutes(1)

	if g.currOnWayRem > 0 {
		return g.Search()
	}

	potential := []int{}

	if g.IsCurrClosed() && g.curr.Rate > 0 {
		return g.CurrOpen().Search()
	}

	for _, edge := range g.curr.Edges {
		if edge.Dist <= g.Minutes && !strings.Contains(g.Open, edge.Dest.Name) {
			potential = append(potential, g.MoveTo(edge.Dest).Search())
		}
	}

	sort.Ints(potential)

	if len(potential) == 0 {
		return g.Minutes*g.FlowRate + g.TotalFlowed
	}

	return potential[len(potential)-1]
}

var MAX int

func (g Graph) SearchPart2() int {
	g.AdvanceMinutes(1)

	if g.Minutes == 0 {
		if g.TotalFlowed > MAX {
			MAX = g.TotalFlowed
		}
		return g.TotalFlowed
	}

	if g.curr2OnWayRem > 0 && g.currOnWayRem > 0 {
		return g.SearchPart2()
	}

	if g.Minutes*206+g.TotalFlowed < MAX {
		return -1
	}

	potential := []Graph{}
	if g.curr2OnWayRem > 0 {
		if g.IsCurrClosed() && g.curr.Rate > 0 {
			return g.CurrOpen().SearchPart2()
		}

		for _, edge := range g.curr.Edges {
			if edge.Dist <= g.Minutes && !strings.Contains(g.Open, edge.Dest.Name) && edge.Dest != g.curr2 {
				potential = append(potential, g.MoveTo(edge.Dest))
			}
		}
	} else if g.currOnWayRem > 0 {
		if g.IsCurr2Closed() && g.curr2.Rate > 0 {
			return g.CurrOpen2().SearchPart2()
		}

		for _, edge := range g.curr2.Edges {
			if edge.Dist <= g.Minutes && !strings.Contains(g.Open, edge.Dest.Name) && edge.Dest != g.curr {
				potential = append(potential, g.MoveTo2(edge.Dest))
			}
		}
	} else {
		if g.IsCurrClosed() && g.curr.Rate > 0 && g.IsCurr2Closed() && g.curr2.Rate > 0 {
			return g.CurrOpen().CurrOpen2().SearchPart2()
		}

		if g.IsCurrClosed() && g.curr.Rate > 0 {
			found := false
			for _, edge := range g.curr2.Edges {
				if edge.Dist <= g.Minutes && !strings.Contains(g.Open, edge.Dest.Name) && edge.Dest != g.curr {
					potential = append(potential, g.CurrOpen().MoveTo2(edge.Dest))
					found = true
				}
			}
			if !found {
				potential = append(potential, g.CurrOpen())
			}
		} else if g.IsCurr2Closed() && g.curr2.Rate > 0 {
			found := false
			for _, edge := range g.curr.Edges {
				if edge.Dist <= g.Minutes && !strings.Contains(g.Open, edge.Dest.Name) && edge.Dest != g.curr2 {
					potential = append(potential, g.CurrOpen2().MoveTo(edge.Dest))
				}
			}
			if !found {
				potential = append(potential, g.CurrOpen2())
			}
		} else {
			found := false
			for _, edge := range g.curr.Edges {
				for _, edge2 := range g.curr2.Edges {
					if edge.Dist <= g.Minutes && !strings.Contains(g.Open, edge.Dest.Name) && edge2.Dist <= g.Minutes && !strings.Contains(g.Open, edge2.Dest.Name) && edge2.Dest != edge.Dest {
						potential = append(potential, g.MoveTo(edge.Dest).MoveTo2(edge2.Dest))
						found = true
					}
				}
			}

			if !found {
				for _, edge := range g.curr.Edges {
					for _, edge2 := range g.curr2.Edges {
						if edge.Dist <= g.Minutes && !strings.Contains(g.Open, edge.Dest.Name) {
							potential = append(potential, g.MoveTo(edge.Dest))
						} else if edge2.Dist <= g.Minutes && !strings.Contains(g.Open, edge2.Dest.Name) {
							potential = append(potential, g.MoveTo2(edge2.Dest))
						}
					}
				}
			}
		}
	}

	if len(potential) == 0 && (g.curr2OnWayRem > 0 || g.currOnWayRem > 0) {
		return g.SearchPart2()
	}

	if len(potential) == 0 {
		g.TotalFlowed += g.Minutes * g.FlowRate

		if g.TotalFlowed > MAX {
			MAX = g.TotalFlowed
		}

		return g.TotalFlowed
	}

	max := 0
	for _, g := range potential {
		pathMax := g.SearchPart2()

		if pathMax > max {
			max = pathMax
		}
	}

	return max
}

func part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	nodes := []*Node{}
	for scanner.Scan() {
		v := parseLine(scanner.Text())
		nodes = append(nodes, v.Node())
	}

	djikstra(nodes)

	var start *Node
	non0Nodes := []*Node{}
	for _, node := range nodes {
		if node.Rate != 0 || node.Name == "AA" {
			non0Nodes = append(non0Nodes, node)
		}

		if node.Name == "AA" {
			start = node
		}
	}

	g := &Graph{
		Minutes: 30,
		curr:    start,
	}

	return g.Search()
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	nodes := []*Node{}
	for scanner.Scan() {
		v := parseLine(scanner.Text())
		nodes = append(nodes, v.Node())
	}

	djikstra(nodes)

	var start *Node
	for _, node := range nodes {
		if node.Name == "AA" {
			start = node
		}
	}

	g := &Graph{
		Minutes: 26,
		curr:    start,
		curr2:   start,
	}

	return g.SearchPart2()
}
