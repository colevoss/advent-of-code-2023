package main

import (
	"fmt"
	"strconv"
	"strings"
)

type PartOne struct {
	Grid  *Grid
	Parts []*Part
}

func NewPartOne() *PartOne {
	return &PartOne{
		Grid: &Grid{
			Rows: []Row{},
		},
		Parts: []*Part{},
	}
}

func (p *PartOne) Init() {
	fmt.Println("Initializing Day 03 Part 1")
}

func (p *PartOne) ReadLine(line string, idx int) {
	// fmt.Println("Reading line", line, idx)
	p.ParseLine(line, idx)
}

func (p *PartOne) ParseLine(line string, idx int) {
	split := strings.Split(line, "")
	row := NewRow(len(split))

	var currentPart *Part

	for i, char := range split {
		// Check if char can convert to number
		_, err := strconv.Atoi(char)

		// if it can't it must be a symbol
		if err != nil {
			// if not a period, create symbol
			if char != "." {
				symbol := NewSymbol(char)
				row.AddSymbol(symbol, i)
			}

			// finish current part
			if currentPart != nil {
				currentPart.Finish()
				p.Parts = append(p.Parts, currentPart)
				currentPart = nil
			}

			continue
		}

		if currentPart == nil {
			currentPart = NewPart(idx, i, char)
			continue
		}

		currentPart.AddDigit(char)
	}

	if currentPart != nil {
		currentPart.Finish()
		p.Parts = append(p.Parts, currentPart)
	}

	p.Grid.Rows = append(p.Grid.Rows, row)
}

func (p *PartOne) Finish() {
	sum := 0

	for _, pa := range p.Parts {
		part := pa

		adj := p.Grid.IsAdjacent(part)

		if adj {
			// fmt.Printf("%+v\n", part)
			sum += part.Num
		}
	}

	fmt.Printf("Day 03 Part 1:\n")
	fmt.Printf("Solution: %d\n", sum)
}
