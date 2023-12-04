package main

import (
	"fmt"
	"strconv"
	"strings"
)

type PartTwo struct {
	GG *GearGrid
}

func NewPartTwo() *PartTwo {
	return &PartTwo{
		GG: NewGearGrid(),
	}
}

func (p *PartTwo) Init() {
	fmt.Println("Initializing Day 03 Part 2")
}

func (p *PartTwo) ReadLine(line string, idx int) {
	// fmt.Println("Reading line", line, idx)
	p.ParseLine(line, idx)
}

func (p *PartTwo) ParseLine(line string, idx int) {
	split := strings.Split(line, "")
	symbolRow := NewRow(len(split))
	partRow := NewPartsRow()

	var currentPart *Part

	for i, char := range split {
		_, err := strconv.Atoi(char)

		if err != nil {
			if char == "*" {
				symbol := NewSymbol(char)
				symbolRow.AddSymbol(symbol, i)
			}

			if currentPart != nil {
				currentPart.Finish()
				partRow = append(partRow, currentPart)
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
		partRow = append(partRow, currentPart)
	}

	p.GG.Parts = append(p.GG.Parts, partRow)
	p.GG.SymbolRows = append(p.GG.SymbolRows, symbolRow)
}

func (p *PartTwo) Finish() {
	fmt.Printf("Day 03 Part 2:\n")

	sum := p.GG.GetSum()
	fmt.Printf("%d\n", sum)
}
