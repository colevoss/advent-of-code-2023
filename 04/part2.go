package main

import (
	"fmt"
)

type PartTwo struct {
	Cards []*Card
}

func NewPartTwo() *PartTwo {
	return &PartTwo{
		Cards: []*Card{},
	}
}

func (p *PartTwo) Init() {
	fmt.Println("Initializing Day 04 Part 2")
}

func (p *PartTwo) ReadLine(line string, idx int) {
	card := NewCard(line, idx)
	card.Process()
	p.Cards = append(p.Cards, card)
}

func (p *PartTwo) Finish() {
	fmt.Printf("Day 04 Part 2:\n")
	sum := p.ProcessCopies()
	fmt.Printf("%d\n", sum)
}

func (p *PartTwo) ProcessCopies() int {
	sum := 0

	for cardIndex, card := range p.Cards {
		fmt.Printf("Card %d has %d matches and %d copies\n", cardIndex, card.WinCount, card.Copies)

		for i := cardIndex + 1; i <= card.WinCount+cardIndex; i++ {
			p.Cards[i].Copies += card.Copies
			fmt.Printf("  %d - %d\n", i, p.Cards[i].Copies)
		}

		sum += card.Copies
	}

	return sum
}
