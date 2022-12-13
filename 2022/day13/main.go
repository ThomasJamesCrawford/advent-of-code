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

type Node struct {
	Value  int
	Values []*Node
}

func (n *Node) String() string {
	if n.Value != -1 {
		return fmt.Sprintf("%d", n.Value)
	}

	return fmt.Sprintf("%s", n.Values)
}

func (n *Node) Sum() int {
	if n.Value != -1 {
		return n.Value
	}

	sum := 0

	for _, v := range n.Values {
		sum += v.Sum()
	}

	return sum
}

func (l *Node) Compare(r *Node) int {
	if l.IsValue() && r.IsValue() {
		fmt.Println("\t- COMPARE", l.Value, "vs", r.Value)
		if l.Value < r.Value {
			fmt.Println("\t- Left side is smaller, so inputs are in the right order")
			return 1
		}

		if l.Value > r.Value {
			fmt.Println("\t- Right side is smaller, so inputs are not in the right order")
			return 0
		}

		return -1
	}

	if l.IsValue() || r.IsValue() {
		fmt.Println("\t- COMPARE", l, "vs", r)

		if l.IsValue() {
			l.Values = []*Node{{Value: l.Value}}
			fmt.Println("\t- Mixed types; convert left to", l.Values)
		}

		if r.IsValue() {
			r.Values = []*Node{{Value: r.Value}}
			fmt.Println("\t- Mixed types; convert right to", r.Values)
		}
	}

	fmt.Println("\t- COMPARE", l.Values, "vs", r.Values)

	for i := 0; true; i++ {

		if len(l.Values) <= i && len(r.Values) <= i {
			return -1
		}

		if len(l.Values) <= i {
			fmt.Println("\t- Left side ran out of items, so inputs are in the right order")
			return 1
		}

		if len(r.Values) <= i {
			fmt.Println("\t- Right side ran out of items, so inputs are not in the right order")
			return 0
		}

		l := l.Values[i]
		r := r.Values[i]

		res := l.Compare(r)

		if res == 0 || res == 1 {
			return res
		}
	}

	return -1
}

func (n *Node) IsValue() bool {
	return n.Value != -1
}

func pLine(l string) *Node {
	n := &Node{
		Value: -1,
	}

	if len(l) == 2 {
		n.Value = -1
		return n
	}

	l = l[1 : len(l)-1]

	open := 0
	segments := []string{}

	segment := ""
	for _, c := range l {
		if c == '[' {
			open++
		}

		if c == ']' {
			open--
		}

		if c == ',' && open == 0 {
			segments = append(segments, segment)
			segment = ""
		} else {
			segment += fmt.Sprintf("%c", c)
		}
	}
	segments = append(segments, segment)

	for _, segment := range segments {
		if intVal, err := strconv.Atoi(segment); err != nil {
			n.Values = append(n.Values, pLine(segment))
			n.Value = -1
		} else {
			n.Values = append(n.Values, &Node{Value: intVal})
			n.Value = -1
		}
	}

	return n
}

func part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	sum := 0
	index := 1
	for scanner.Scan() {
		fmt.Println("== Pair", index, "==")
		packet1 := pLine(scanner.Text())

		scanner.Scan()
		packet2 := pLine(scanner.Text())

		fmt.Println(packet1)
		fmt.Println(packet2)

		fmt.Println()

		res := packet1.Compare(packet2)

		if res == 1 {
			fmt.Println(index, "was correct!")
			sum += index
		}
		fmt.Println()
		scanner.Scan()

		index++
	}

	return sum
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	d1 := &Node{
		Value:  -1,
		Values: []*Node{{Value: 2}},
	}

	d2 := &Node{
		Value:  -1,
		Values: []*Node{{Value: 6}},
	}

	packets := []*Node{d1, d2}
	for scanner.Scan() {
		packet1 := pLine(scanner.Text())

		scanner.Scan()
		packet2 := pLine(scanner.Text())

		packets = append(packets, packet1, packet2)

		scanner.Scan()
	}

	sort.Slice(packets, func(i, j int) bool {
		l := packets[i]
		r := packets[j]

		comp := l.Compare(r)

		return comp == 1
	})

	product := 1
	for i, p := range packets {
		fmt.Println(p)
		if p == d1 || p == d2 {
			product *= (i + 1)
		}

	}

	return product
}
