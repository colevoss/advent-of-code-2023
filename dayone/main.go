package main

import (
	"aoc/runner"
)

func main() {
	runner := runner.NewRunner("dayone/input.txt")

	partone := &PartOne{
		Nums: []int{},
	}

	runner.Run(partone)
}
