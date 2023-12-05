package main

import (
	"fmt"
	"math"
	"sync"
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

func (p *PartTwo) Finish() {
	start := time.Now()

	p.FixSeedPairs()

	min := math.MaxInt

	var wg sync.WaitGroup
	c := make(chan int)

	for _, sp := range p.SeedPairs {
		wg.Add(1)
		go func(sp SeedRange) {
			defer wg.Done()
			spMin := math.MaxInt

			spStart := time.Now()
			fmt.Printf("Starting seed pair: %+v\n", sp)
			top := sp.Start + sp.Range

			for s := sp.Start; s < top; s++ {
				seed := s

				for _, m := range p.Maps {
					seed = m.MapSeed(seed)
				}

				if seed < spMin {
					spMin = seed
				}
			}

			c <- spMin

			spEnd := time.Now()
			dur := spEnd.Sub(spStart)
			fmt.Printf("Seed Pair Finished: %+v\n", dur)
		}(sp)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for minNum := range c {
		if minNum < min {
			min = minNum
		}
	}

	end := time.Now()
	dur := end.Sub(start)

	fmt.Printf("Day 05 Part 2:\n")
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
