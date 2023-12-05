package main

import (
	"fmt"
	"strconv"
	"strings"
)

const BLANK_LINE = ""

type PartOne struct {
	StartNewMap bool
	CurrentMap  *Map
	Seeds       []int
	Maps        SeedMap
}

func NewPartOne() *PartOne {
	return &PartOne{
		StartNewMap: false,
		CurrentMap:  nil,
		Maps:        []*Map{},
		Seeds:       []int{},
	}
}

func (p *PartOne) Init() {
	fmt.Println("Initializing Day 05 Part 1")
}

func (p *PartOne) ReadLine(line string, idx int) {
	// seeds is always the first line
	if idx == 0 {
		p.Seeds = parseNumsString(line)
		return
	}

	// map is finished when a blank line happens
	if line == BLANK_LINE {
		p.StartNewMap = true
		return
	}

	if p.StartNewMap {
		currentMap := parseMap(line)
		p.Maps = append(p.Maps, currentMap)
		p.StartNewMap = false
		return
	}

	currentMap := p.LatestMap()

	// this should never happen
	if currentMap == nil {
		panic("balls")
	}

	// If no other conditions met, must be range line
	mapRange := parseRangeString(line)
	currentMap.Ranges = append(currentMap.Ranges, mapRange)
}

func (p *PartOne) LatestMap() *Map {
	return p.Maps[len(p.Maps)-1]
}

func parseNumsString(line string) []int {
	split := strings.Split(line, ": ")
	numArr := splitNumStr(split[1])

	return numArr
}

func splitNumStr(str string) []int {
	split := strings.Split(str, " ")

	nums := make([]int, len(split))

	for idx, num := range split {
		trimmed := strings.TrimSpace(num)
		i, err := strconv.Atoi(trimmed)

		if err != nil {
			panic(err)
		}

		nums[idx] = i
	}

	return nums
}

func (p *PartOne) Finish() {
	mapped := make([]int, len(p.Seeds))

	for seedIdx, s := range p.Seeds {
		seed := s

		for _, m := range p.Maps {
			// fmt.Printf("%s", m)
			seed = m.MapSeed(seed)
		}

		mapped[seedIdx] = seed
		// fmt.Printf("%d -> %d\n", s, seed)
	}

	min := sliceMin(mapped)

	fmt.Printf("Day 05 Part 1:\n")
	fmt.Printf("%d\n", min)
}

func sliceMin(s []int) int {
	min := s[0]

	for _, i := range s {
		if i < min {
			min = i
		}
	}

	return min
}
