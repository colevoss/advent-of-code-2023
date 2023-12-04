package main

type PartsRow []*Part

func NewPartsRow() PartsRow {
	return []*Part{}
}

type GearGrid struct {
	SymbolRows []Row
	Parts      []PartsRow
}

func NewGearGrid() *GearGrid {
	return &GearGrid{
		SymbolRows: []Row{},
		Parts:      []PartsRow{},
	}
}

func (g *GearGrid) GetSum() int {
	sum := 0
	for rowIdx, symbolRow := range g.SymbolRows {

	symbolLoop:
		for symboldIdx, symbol := range symbolRow {
			// might as well check if its not an asterisk/gear
			if symbol == nil || symbol.Value != "*" {
				continue symbolLoop
			}

			adjacentParts := []*Part{}

			// Find all adjacent parts to this symbol
			thisPartsRow := g.Parts[rowIdx]
			shouldCheckLeft := symboldIdx > 0
			shouldCheckRight := symboldIdx < len(symbolRow)-1

			for _, p := range thisPartsRow {
				part := p
				// check left
				if shouldCheckLeft && p.ColEnd == symboldIdx-1 {
					adjacentParts = append(adjacentParts, part)
				}

				// check right
				if shouldCheckRight && p.ColStart == symboldIdx+1 {
					adjacentParts = append(adjacentParts, part)
				}
			}

			shouldCheckPrev := rowIdx > 0
			shouldCheckNext := rowIdx < len(symbolRow)-1

			if shouldCheckPrev {
				prevRow := g.Parts[rowIdx-1]

				for _, part := range prevRow {
					p := part
					pLeftBound := max(0, p.ColStart-1)
					pRightBound := min(len(symbolRow), p.ColEnd+1)

					if symboldIdx >= pLeftBound && symboldIdx <= pRightBound {
						adjacentParts = append(adjacentParts, p)
					}
				}
			}

			if shouldCheckNext {
				prevRow := g.Parts[rowIdx+1]

				for _, part := range prevRow {
					p := part
					pLeftBound := max(0, p.ColStart-1)
					pRightBound := min(len(symbolRow), p.ColEnd+1)

					if symboldIdx >= pLeftBound && symboldIdx <= pRightBound {
						adjacentParts = append(adjacentParts, p)
					}
				}
			}

			if len(adjacentParts) == 2 {
				product := adjacentParts[0].Num * adjacentParts[1].Num
				sum += product
			}
		}
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
