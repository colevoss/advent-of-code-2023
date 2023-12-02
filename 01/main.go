package main

import (
	"aoc/runner"
)

func main() {
	runner := runner.NewRunner("input/01.txt")

	partone := &PartOne{
		Sum: 0,
	}
	runner.Run(partone)

	parttwo := &PartTwo{
		Sum: 0,
	}

	runner.Run(parttwo)
}
