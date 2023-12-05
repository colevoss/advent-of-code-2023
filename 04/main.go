package main

import (
	"aoc/runner"
)

func main() {
	runner := runner.NewRunner("input/04.txt")

	// partone := NewPartOne()
	// runner.Run(partone)

	parttwo := NewPartTwo()
	runner.Run(parttwo)
}
