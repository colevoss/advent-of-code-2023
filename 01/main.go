package main

import (
	"aoc/runner"
)

func main() {
	runner := runner.NewRunner("01/input.txt")

	partone := &PartOne{
		Sum: 0,
	}
	runner.Run(partone)

	parttwo := &PartTwo{
		Sum: 0,
	}

	runner.Run(parttwo)
}
