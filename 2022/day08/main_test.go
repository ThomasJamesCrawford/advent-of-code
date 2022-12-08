package main

import (
	"io"
	"testing"
	"strings"
)

// paste the sample in here
const SAMPLE string = `30373
25512
65332
33549
35390`

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
			want: 21,
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
			want: 8,
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
