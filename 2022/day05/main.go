package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f, _ := os.Open("input.txt")

	defer f.Close()

	fmt.Println("part1: ", part1(f))

	f.Seek(0, io.SeekStart)
	fmt.Println("part2: ", part2(f))
}

func processItems(s string, stacks map[int]*Stack) {
	for i, c := range s {
		if unicode.IsSpace(c) {
			continue
		}

		if i%4 == 1 {
			stackNumber := (i / 4) + 1

			stack, ok := stacks[stackNumber]

			if !ok {
				stacks[stackNumber] = &Stack{items: []rune{c}}
			} else {
				stack.Push(c)
			}
		}
	}
}

type Stack struct {
	items []rune
}

func (s *Stack) Reverse() {
	for i, j := 0, len(s.items)-1; i < j; i, j = i+1, j-1 {
		s.items[i], s.items[j] = s.items[j], s.items[i]
	}
}

func (s *Stack) Pop() rune {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Push(c rune) {
	s.items = append(s.items, c)
}

func (s *Stack) Print() {
	for _, c := range s.items {
		fmt.Printf("%c", c)
	}
}

func part1(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	ontoMoves := false
	stacks := map[int]*Stack{}
	for scanner.Scan() {
		line := scanner.Text()

		if !ontoMoves && strings.Contains(line, "1") {
			for _, s := range stacks {
				s.Reverse()
			}

			ontoMoves = true
			scanner.Scan()
			continue
		}

		if !ontoMoves {
			processItems(line, stacks)
		} else {
			values := strings.Split(line, " ")

			count, _ := strconv.Atoi(values[1])
			from, _ := strconv.Atoi(values[3])
			to, _ := strconv.Atoi(values[5])

			for i := 0; i < count; i++ {
				stacks[to].Push(stacks[from].Pop())
			}
		}
	}

	answer := ""
	for i := 1; i <= len(stacks); i++{
		answer += fmt.Sprintf("%c", stacks[i].Pop())
	}

	return answer
}

func part2(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	ontoMoves := false
	stacks := map[int]*Stack{}
	for scanner.Scan() {
		line := scanner.Text()

		if !ontoMoves && strings.Contains(line, "1") {
			for _, s := range stacks {
				s.Reverse()
			}

			ontoMoves = true
			scanner.Scan()
			continue
		}

		if !ontoMoves {
			processItems(line, stacks)
		} else {
			values := strings.Split(line, " ")

			count, _ := strconv.Atoi(values[1])
			from, _ := strconv.Atoi(values[3])
			to, _ := strconv.Atoi(values[5])

			moves := []rune{}
			for i := 0; i < count; i++ {
				moves = append(moves, stacks[from].Pop())
			}

			for i, j := 0, len(moves)-1; i < j; i, j = i+1, j-1 {
				moves[i], moves[j] = moves[j], moves[i]
			}

			for _, m := range moves {
				stacks[to].Push(m)
			}
		}
	}

	answer := ""
	for i := 1; i <= len(stacks); i++{
		answer += fmt.Sprintf("%c", stacks[i].Pop())
	}

	return answer
}
