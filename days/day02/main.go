package day00

import (
	"github.com/samueldumont/adventofcode2021/v2/utils"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	length    int
}

func main() {
	part1(utils.ReadFile("\n"))
}

func part1(input []string) int {
	position := 0
	depth := 0
	parsed := parseInputs(input)

	for _, n := range parsed {
		if n.direction == "forward" {
			position += n.length
		} else if n.direction == "backward" {
			position -= n.length
		} else if n.direction == "down" {
			depth += n.length
		} else if n.direction == "up" {
			depth -= n.length
		}
	}

	return position * depth
}

func part2(input []string) int {
	position := 0
	depth := 0
	aim := 0
	parsed := parseInputs(input)

	for _, n := range parsed {
		if n.direction == "forward" {
			position += n.length
			depth += aim * n.length
		} else if n.direction == "down" {
			aim += n.length
		} else if n.direction == "up" {
			aim -= n.length
		}
	}

	return position * depth
}

func parseInputs(input []string) []instruction {

	var instructions []instruction
	for _, n := range input {
		line := strings.Fields(n)
		distance, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
		inst := instruction{direction: line[0], length: distance}
		instructions = append(instructions, inst)
	}

	return instructions
}
