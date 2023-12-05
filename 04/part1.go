package main

import (
	"fmt"
	"strings"
)

type PartOne struct {
	Sum int
}

func NewPartOne() *PartOne {
	return &PartOne{
		Sum: 0,
	}
}

func (p *PartOne) Init() {
	fmt.Println("Initializing Day 04 Part 1")
}

func createNumArr(str string) []string {
	return splitReg.Split(strings.TrimSpace(str), -1)
}

func (p *PartOne) ReadLine(line string, idx int) {
	card := NewCard(line, idx)
	card.Process()
	score := card.Score()
	p.Sum += score
	// fmt.Printf("%d MatchCount: %d, Score: %d, total %d\n", idx+1, matchCount, score, p.Sum)
}

func (p *PartOne) Finish() {
	fmt.Printf("Day 04 Part 1:\n")
	fmt.Printf("%d\n", p.Sum) // 24542
}
