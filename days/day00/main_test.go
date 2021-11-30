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
	{"example", 514579, []string{"1721", "979", "366", "299", "675", "1456"}},
	{"actual", 444019, utils.ReadFile("\n")},
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
	{"example", 241861950, []string{"1721", "979", "366", "299", "675", "1456"}},
	{"actual", 29212176, utils.ReadFile("\n")},
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