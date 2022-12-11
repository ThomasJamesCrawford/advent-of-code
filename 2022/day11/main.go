package main

import (
	"bufio"
	"fmt"
	"io"
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

type Monkey struct {
	items     []int
	operation func(old int) int
	test      int
	ifTrue    int
	ifFalse   int
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	monkeys := map[int]*Monkey{}
	inspect := map[int]int{}
	monkeyCount := 0

	for scanner.Scan() {
		m := &Monkey{}

		scanner.Scan()
		_, itemsString, _ := strings.Cut(scanner.Text(), "Starting items: ")

		items := strings.Split(itemsString, ", ")
		for _, i := range items {
			itemInt, _ := strconv.Atoi(i)
			m.items = append(m.items, itemInt)
		}

		scanner.Scan()
		_, fnString, _ := strings.Cut(scanner.Text(), "Operation: new = old ")
		op, value, _ := strings.Cut(fnString, " ")

		if value == "old" {
			switch op {
			case "+":
				m.operation = func(a int) int {
					return a + a
				}
			case "-":
				m.operation = func(a int) int {
					return 0
				}
			case "*":
				m.operation = func(a int) int {
					return a * a
				}
			}
		} else {
			valueInt, _ := strconv.Atoi(value)
			switch op {
			case "+":
				m.operation = func(a int) int {
					return a + valueInt
				}
			case "-":
				m.operation = func(a int) int {
					return a - valueInt
				}
			case "*":
				m.operation = func(a int) int {
					return a * valueInt
				}
			}
		}

		scanner.Scan()
		_, valueString, _ := strings.Cut(scanner.Text(), "Test: divisible by ")

		test, _ := strconv.Atoi(valueString)
		m.test = test

		scanner.Scan()
		_, valueString, _ = strings.Cut(scanner.Text(), "If true: throw to monkey ")
		test, _ = strconv.Atoi(valueString)
		m.ifTrue = test

		scanner.Scan()
		_, valueString, _ = strings.Cut(scanner.Text(), "If false: throw to monkey ")
		test, _ = strconv.Atoi(valueString)
		m.ifFalse = test

		monkeys[monkeyCount] = m
		monkeyCount++

		scanner.Scan()
	}

	for r := 1; r <= 20; r++ {
		for i := 0; i < len(monkeys); i++ {
			m := monkeys[i]

			for _, item := range m.items {
				value := m.operation(item) / 3

				if value % m.test == 0 {
					monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, value)
				} else {
					monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, value)
				}

				if _, ok := inspect[i]; !ok {
					inspect[i] = 1
				} else {
					inspect[i] += 1
				}
			}

			m.items = []int{}
		}
	}

	inspectCounts := []int{}
	for _, count := range inspect {
		inspectCounts = append(inspectCounts, count)
	}

	sort.Ints(inspectCounts)

	return inspectCounts[len(inspectCounts)-1] * inspectCounts[len(inspectCounts)-2]
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	monkeys := map[int]*Monkey{}
	inspect := map[int]int{}
	monkeyCount := 0

	divisors := []int{}

	for scanner.Scan() {
		m := &Monkey{}

		scanner.Scan()
		_, itemsString, _ := strings.Cut(scanner.Text(), "Starting items: ")

		items := strings.Split(itemsString, ", ")
		for _, i := range items {
			itemInt, _ := strconv.Atoi(i)
			m.items = append(m.items, itemInt)
		}

		scanner.Scan()
		_, fnString, _ := strings.Cut(scanner.Text(), "Operation: new = old ")
		op, value, _ := strings.Cut(fnString, " ")

		if value == "old" {
			switch op {
			case "+":
				m.operation = func(a int) int {
					return a + a
				}
			case "-":
				m.operation = func(a int) int {
					return 0
				}
			case "*":
				m.operation = func(a int) int {
					return a * a
				}
			}
		} else {
			valueInt, _ := strconv.Atoi(value)
			switch op {
			case "+":
				m.operation = func(a int) int {
					return a + valueInt
				}
			case "-":
				m.operation = func(a int) int {
					return a - valueInt
				}
			case "*":
				m.operation = func(a int) int {
					return a * valueInt
				}
			}
		}

		scanner.Scan()
		_, valueString, _ := strings.Cut(scanner.Text(), "Test: divisible by ")

		test, _ := strconv.Atoi(valueString)
		m.test = test

		scanner.Scan()
		_, valueString, _ = strings.Cut(scanner.Text(), "If true: throw to monkey ")
		test, _ = strconv.Atoi(valueString)
		m.ifTrue = test

		scanner.Scan()
		_, valueString, _ = strings.Cut(scanner.Text(), "If false: throw to monkey ")
		test, _ = strconv.Atoi(valueString)
		m.ifFalse = test

		monkeys[monkeyCount] = m
		monkeyCount++

		divisors = append(divisors, m.test)

		scanner.Scan()
	}

	lcm := LCM(divisors[0], divisors[1], divisors[2:]...)

	for r := 1; r <= 10000; r++ {
		for i := 0; i < len(monkeys); i++ {
			m := monkeys[i]

			for _, item := range m.items {
				value := m.operation(item) % lcm

				if value % m.test == 0 {
					monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, value)
				} else {
					monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, value)
				}

				if _, ok := inspect[i]; !ok {
					inspect[i] = 1
				} else {
					inspect[i] += 1
				}
			}

			m.items = []int{}
		}
	}

	inspectCounts := []int{}
	for _, count := range inspect {
		inspectCounts = append(inspectCounts, count)
	}

	sort.Ints(inspectCounts)

	return inspectCounts[len(inspectCounts)-1] * inspectCounts[len(inspectCounts)-2]
}
