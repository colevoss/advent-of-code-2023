package main

import (
	"aoc/runner"
)

func main() {
	runner := runner.NewRunner("input/08.txt")

	// partone := NewPartOne()
	// runner.Run(partone)

	// Uncomment to work on part two
	parttwo := NewPartTwo()
	runner.Run(parttwo)
}
