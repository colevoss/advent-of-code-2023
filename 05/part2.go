package main

import (
	"fmt"
	"math"
	"time"
)

type SeedRange struct {
	Start int
	Range int
}

type PartTwo struct {
	*PartOne
	SeedPairs []SeedRange
}

func NewPartTwo() *PartTwo {
	return &PartTwo{
		PartOne:   NewPartOne(),
		SeedPairs: []SeedRange{},
	}
}

func (p *PartTwo) Init() {
	fmt.Println("Initializing Day 05 Part 2")
}

// func (p *PartTwo) ReadLine(line string, idx int) {
// }

func (p *PartTwo) Finish() {
	start := time.Now()

	fmt.Printf("Day 05 Part 2:\n")
	p.FixSeedPairs()

	min := math.MaxInt
	// min := -1

	for _, sp := range p.SeedPairs {
		fmt.Printf("Starting seed pair: %+v\n", sp)
		for s := sp.Start; s < sp.Start+sp.Range; s++ {
			seed := s

			for _, m := range p.Maps {
				seed = m.MapSeed(seed)
			}

			// if seed < min || min == -1 {
			if seed < min {
				min = seed
			}
		}

		spEnd := time.Now()
		dur := spEnd.Sub(start)
		fmt.Printf("Seed Pair Finished: %+v\n", dur)
	}

	end := time.Now()
	dur := end.Sub(start)

	fmt.Printf("%d\n", min)
	fmt.Printf("Duration %+v\n", dur)
}

func (p *PartTwo) FixSeedPairs() {
	for i := 0; i <= len(p.Seeds)-1; i += 2 {
		start := p.Seeds[i]  // 0, 2, 4
		rnge := p.Seeds[i+1] // 1

		seedRange := SeedRange{start, rnge}
		p.SeedPairs = append(p.SeedPairs, seedRange)
	}
}
