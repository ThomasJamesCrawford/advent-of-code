package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	force := flag.Bool("force", false, "overwrite existing files")
	day := flag.Int("day", 0, "day")
	year := flag.Int("year", 2022, "year")

	flag.Parse()

	if *day == 0 {
		panic("--day must be set")
	}

	makeNewDay(*day, *year, *force)
}

func makeNewDay(day int, year int, force bool) {
	p := fmt.Sprintf("%d/day%02d/", year, day)

	_, err := os.Stat(path.Join(p, "input.txt"))
	if err == nil && !force {
		fmt.Printf("day/year already exists, use --force to override\n")
		return
	}

	os.MkdirAll(p, os.ModePerm)

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", os.Getenv("AOC_SESSION_TOKEN")))
	req.Header.Add("User-Agent", "input-pulling-script")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	file, err := os.Create(path.Join(p, "input.txt"))
	if err != nil {
		panic(err)
	}

	defer file.Close()

	io.Copy(file, resp.Body)

	f1, err := os.Create(path.Join(p, "main.go"))
	if err != nil {
		panic(err)
	}

	defer f1.Close()

	io.Copy(f1, strings.NewReader(
		`package main

import (
	"fmt"
	"io"
	"os"
	"bufio"
)

func main() {
	f, _ := os.Open("input.txt")

	defer f.Close()

	fmt.Println("part1: ", part1(f))
	fmt.Println("part2: ", part2(f))
}

func part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		//line := scanner.Text()
	}

	return 0
}

func part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		//line := scanner.Text()
	}

	return 0
}
`))

	f2, err := os.Create(path.Join(p, "main_test.go"))
	if err != nil {
		panic(err)
	}

	defer f2.Close()

	io.Copy(f2, strings.NewReader(`package main

import (
	"io"
	"testing"
	"strings"
)

// paste the sample in here
`+"const SAMPLE string = ``"+`

func Test_part1(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{ input: strings.NewReader(SAMPLE)},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{ input: strings.NewReader(SAMPLE)},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
`))
}
