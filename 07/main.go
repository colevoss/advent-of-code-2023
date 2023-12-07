package main

import (
	"aoc/runner"
)

func main() {
	runner := runner.NewRunner("input/07.txt")

	partone := NewPartOne()
	runner.Run(partone)

	// parttwo := NewPartTwo()
	// runner.Run(parttwo)
}
