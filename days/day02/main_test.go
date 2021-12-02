package day00

import (
	"github.com/samueldumont/adventofcode2021/v2/utils"
	"testing"
)

var tests1 = []struct {
	name  string
	want  int
	input []string
	// add extra args if needed
}{
	{"example", 150, []string{"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2"}},
	{"actual", 1762050, utils.ReadFile("\n")},
}

func TestPart1(t *testing.T) {
	for _, test := range tests1 {
		t.Run(test.name, func(*testing.T) {
			got := part1(test.input)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

var tests2 = []struct {
	name  string
	want  int
	input []string
	// add extra args if needed
}{
	{"example", 900, []string{"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2"}},
	{"actual", 1855892637, utils.ReadFile("\n")},
}

func TestPart2(t *testing.T) {
	for _, test := range tests2 {
		t.Run(test.name, func(*testing.T) {
			got := part2(test.input)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}