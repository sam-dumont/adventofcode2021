package day00

import (
	"github.com/samueldumont/adventofcode2021/v2/utils"
	"strconv"
)

func main() {
	part1(utils.ReadFile("\n"))
}

func part1(input []string) int {
	counter := 0
	parsed := parseInputs(input)

	for i, n := range parsed[1:]{
		if n > parsed[i] {
			counter += 1
		}
	}

	return counter
}

func part2(input []string) int {
	counter := 0
	parsed := parseInputs(input)

	for i, n := range parsed[3:]{
		if n > parsed[i] {
			counter += 1
		}
	}

	return counter
}

func parseInputs(input []string) []int {

	nums := []int{}
	for _, n := range input {
		conv, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		nums = append(nums, conv)
	}

	return nums
}
