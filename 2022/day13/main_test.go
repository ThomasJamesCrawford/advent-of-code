package main

import (
	"io"
	"testing"
	"strings"
)

// paste the sample in here
const SAMPLE string = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

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
			args: args{ input: strings.NewReader(SAMPLE)},
			want: 140,
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
