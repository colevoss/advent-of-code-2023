package main

import (
	"strconv"
)

type Part struct {
	ColStart int // 0     0, 1, 2
	ColEnd   int // 2
	Row      int
	Digits   string
	Num      int
}

func (p *Part) OccupiedCols() []int {
	cols := []int{}

	for i := p.ColStart; i <= p.ColEnd; i++ {
		cols = append(cols, i)
	}

	return cols
}

func (p *Part) IsAdjacentToCol(col int) bool {
	// is inside
	if col >= p.ColStart-1 && col <= p.ColEnd+1 {
		return true
	}

	return false
}

func NewPart(row int, colStart int, firstDigit string) *Part {
	return &Part{
		ColStart: colStart,
		ColEnd:   colStart,
		Row:      row,
		Digits:   firstDigit,
	}
}

func (p *Part) AddDigit(digit string) {
	p.Digits += digit
	p.ColEnd += 1
}

func (p *Part) Finish() {
	num, err := strconv.Atoi(p.Digits)

	if err != nil {
		panic(err)
	}

	p.Num = num
}

type Symbol struct {
	Value string
}

func NewSymbol(val string) *Symbol {
	return &Symbol{val}
}

type Row []*Symbol

func NewRow(length int) Row {
	return make(Row, length)
}

func (row *Row) AddSymbol(sym *Symbol, idx int) {
	(*row)[idx] = sym
}

type Grid struct {
	Rows []Row
}

func (g *Grid) IsAdjacent(part *Part) bool {
	prevRowIdx := part.Row - 1
	nextRowIdx := part.Row + 1

	leftCol := part.ColStart - 1
	rightCol := part.ColEnd + 1

	// check this row
	partRow := g.Rows[part.Row]

	if leftCol >= 0 && partRow[leftCol] != nil {
		return true
	}

	if rightCol < len(partRow) && partRow[rightCol] != nil {
		return true
	}

	// check prev row
	if prevRowIdx >= 0 {
		prevRow := g.Rows[prevRowIdx]

		// left
		if leftCol >= 0 && prevRow[leftCol] != nil {
			return true
		}

		// inside
		for i := part.ColStart; i <= part.ColEnd; i++ {
			if prevRow[i] != nil {
				return true
			}
		}

		// right
		if rightCol < len(prevRow) && prevRow[rightCol] != nil {
			return true
		}
	}

	// check next row
	if nextRowIdx < len(g.Rows) {
		nextRow := g.Rows[nextRowIdx]

		// left
		if leftCol >= 0 && nextRow[leftCol] != nil {
			return true
		}

		// inside
		for i := part.ColStart; i <= part.ColEnd; i++ {
			if nextRow[i] != nil {
				return true
			}
		}

		// right
		if rightCol < len(nextRow) && nextRow[rightCol] != nil {
			return true
		}
	}

	return false
}
