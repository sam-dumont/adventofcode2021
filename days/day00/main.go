package day00

import (
	"github.com/samueldumont/adventofcode2021/v2/utils"
	"strconv"
)

func main() {
	part1(utils.ReadFile("\n"))
}

func part1(input []string) int {
	seen := map[int]bool{}

	for _, n := range parseInputs(input) {
		if seen[n] {
			return n * (2020 - n)
		}
		seen[2020-n] = true
	}

	return -1 // should not be hit
}

func part2(input []string) int {
	parsed := parseInputs(input)

	// O(n^3) is fast enough
	for i := 0; i < len(parsed); i++ {
		for j := i + 1; j < len(parsed); j++ {
			for k := j + 1; k < len(parsed); k++ {
				if parsed[i]+parsed[j]+parsed[k] == 2020 {
					return parsed[i] * parsed[j] * parsed[k]
				}
			}
		}
	}

	return -1 // should not be hit
}

func parseInputs(input []string) []int {

	nums := []int{}
	for _, n := range input {
		conv,err := strconv.Atoi(n)
		if err != nil{
			panic(err)
		}
		nums = append(nums, conv)
	}

	return nums
}