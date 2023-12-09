package main

import (
	"fmt"
)

type PartTwo struct {
	Sum int
}

func NewPartTwo() *PartTwo {
	return &PartTwo{
		Sum: 0,
	}
}

func (p *PartTwo) Init() {
	fmt.Println("Initializing Day 09 Part 2")
}

func (p *PartTwo) ReadLine(line string, idx int) {
	ints := parseLine(line)
	prev := ints.FindPrevious()

	p.Sum += prev
}

func (p *PartTwo) Finish() {
	fmt.Printf("Day 09 Part 2:\n")
	fmt.Printf("%d\n", p.Sum)
}
