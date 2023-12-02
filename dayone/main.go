package main

import (
	"aoc/runner"
)

func main() {
	runner := runner.NewRunner("dayone/input.txt")

	parttwo := &PartTwo{
		Sum: 0,
	}

	runner.Run(parttwo)
}
