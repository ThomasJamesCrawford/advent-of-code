package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

	alphabet := ".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		first := line[:len(line)/2]
		second := line[len(line)/2:]

		for _, c := range first {
			if strings.ContainsRune(second, c) {
				sum += strings.IndexRune(alphabet, c)
				break
			}
		}
	}

	return sum
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	alphabet := ".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	sum := 0
	for scanner.Scan() {
		elf1 := scanner.Text()
		scanner.Scan()
		elf2 := scanner.Text()
		scanner.Scan()
		elf3 := scanner.Text()

		fmt.Println(elf1)
		fmt.Println(elf2)
		fmt.Println(elf3)

		for _, c := range elf1 {
			if strings.ContainsRune(elf2, c) && strings.ContainsRune(elf3, c) {
				sum += strings.IndexRune(alphabet, c)
				break
			}
		}
	}

	return sum
}
