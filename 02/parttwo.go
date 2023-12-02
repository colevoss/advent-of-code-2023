package main

import "fmt"

type PartTwo struct {
	*PartOne
}

func NewPartTwo() *PartTwo {
	p1 := NewPartOne()

	return &PartTwo{
		PartOne: p1,
	}
}

func (p *PartTwo) Init() {
	fmt.Printf("Running Day Two Part Two\n")
}

func (p *PartTwo) ReadLine(line string, idx int) {
	game := p.ParseGame(line)

	minRound := game.MinCubeRound()
	product := minRound.Product()

	p.Sum += product
}

func (p *PartTwo) Finish() {
	fmt.Printf("Day Two Part Two: %d\n", p.Sum)
}
