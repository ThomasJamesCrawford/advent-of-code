package main

import (
	"io"
	"testing"
	"strings"
)

// paste the sample in here
const SAMPLE string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

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
			want: 13,
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
			args: args{ input: strings.NewReader(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`)},
			want: 36,
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
