package main

import (
	"fmt"
	"sort"
	"strings"
)

type PartOne struct {
	Hands []*Hand
}

func NewPartOne() *PartOne {
	return &PartOne{}
}

func (p *PartOne) Init() {
	fmt.Println("Initializing Day 07 Part 1")
}

func (p *PartOne) ReadLine(line string, idx int) {
	parts := strings.Split(line, " ")

	handStr := parts[0]
	bidStr := parts[1]

	hand := NewHand(handStr, bidStr)
	// fmt.Printf("hand: %+v\n", hand)

	p.Hands = append(p.Hands, hand)
}

func (p *PartOne) Finish() {
	fmt.Printf("Day 07 Part 1:\n")
	sort.Sort(SortByHand(p.Hands))

	score := 0
	for i, hand := range p.Hands {
		rank := len(p.Hands) - i
		handScore := rank * hand.Bid
		score += handScore
	}

	fmt.Printf("score: %v\n", score)
}
