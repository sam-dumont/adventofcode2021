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
	{"example", 7, []string{"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"240",
		"269",
		"260",
		"263"}},
	{"actual", 1713, utils.ReadFile("\n")},
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
	{"example", 5, []string{"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263"}},
	{"actual", 1734, utils.ReadFile("\n")},
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