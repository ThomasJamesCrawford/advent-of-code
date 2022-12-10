package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"path"
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

func part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	prefix := []string{"/"}
	prefixSizes := map[string]int{}
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " ")

		if words[1] == "cd" {
			if words[2] == "/" {
				prefix = []string{"/"}
			} else if words[2] == ".." {
				if len(prefix) > 1 {
					prefix = prefix[:len(prefix)-1]
				}
			} else {
				prefix = append(prefix, words[2])
			}
		}

		if words[0] == "dir" {
			continue
		}

		if words[1] == "ls" {
			continue
		}

		size, err := strconv.Atoi(words[0])

		if err == nil {
			for i := len(prefix); i > 0; i-- {
				p := strings.Join(prefix[:i], "/")
				_, ok := prefixSizes[p]

				if !ok {
					prefixSizes[p] = size
				} else {
					prefixSizes[p] += size
				}
			}
		}
	}

	sum := 0
	for _, ps := range prefixSizes {
		if ps <= 100000 {
			sum += ps
		}
	}

	return sum
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	p := "/"
	pathSizes := map[string]int{}
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " ")

		if words[1] == "cd" {
			p = path.Join(p, words[2])
		}

		if words[0] == "dir" {
			continue
		}

		if words[1] == "ls" {
			continue
		}

		size, err := strconv.Atoi(words[0])

		if err == nil {
			_, ok := pathSizes[p]

			if !ok {
				pathSizes[p] = size
			} else {
				pathSizes[p] += size
			}
		}
	}

	aggSizes := map[string]int{}
	for p, size := range pathSizes {
		accum := p
		for {
			_, ok := aggSizes[accum]

			if !ok {
				aggSizes[accum] = size
			} else {
				aggSizes[accum] += size
			}

			if accum == "/" {
				break
			}

			accum = accum[:strings.LastIndex(accum, "/")]

			if accum == "" {
				accum = "/"
			}
		}
	}

	size := math.MaxInt64
	max := aggSizes["/"]
	need := 30000000 - (70000000 - max)
	for _, ps := range aggSizes {
		if ps <= size && ps >= need {
			size = ps
		}
	}

	return size
}
