package main

import (
	"fmt"
	"strings"
)

type PartTwo struct {
	Pattern   []string
	Map       map[string][]string
	Positions []string
}

func NewPartTwo() *PartTwo {
	return &PartTwo{
		Map:       make(map[string][]string),
		Positions: []string{},
	}
}

func (p *PartTwo) Init() {
	fmt.Println("Initializing Day 08 Part 2")
}

func (p *PartTwo) ReadLine(line string, idx int) {
	if idx == 0 {
		p.Pattern = strings.Split(line, "")
		return
	}

	if line == "" {
		return
	}

	key, nodes := ParseMapLine(line)

	if strings.HasSuffix(key, "A") {
		p.Positions = append(p.Positions, key)
	}

	_, exists := p.Map[key]

	if exists {
		panic(fmt.Sprintf("%s exists already. You are fucked", key))
	}

	p.Map[key] = nodes
}

func (p *PartTwo) Finish() {
	fmt.Printf("Day 08 Part 2:\n")
	sol := p.FindSolution()
	fmt.Printf("sol: %v\n", sol)
}

func (p *PartTwo) FindSolution() int {
	results := make([]int, len(p.Positions))

	for i, pos := range p.Positions {
		res := p.walkTree(pos)
		results[i] = res
	}

	lcmResult := 1
	for _, res := range results {
		lcmResult = lcm(lcmResult, res)
	}

	return lcmResult
}

func (p *PartTwo) walkTree(start string) int {
	patternIdx := 0
	patternLen := len(p.Pattern)
	count := 0

	key := start

	for {
		count++
		pattern := p.Pattern[patternIdx]
		nodes := p.Map[key]

		if pattern == "L" {
			key = nodes[0]
		}

		if pattern == "R" {
			key = nodes[1]
		}

		if strings.HasSuffix(key, "Z") {
			return count
		}

		patternIdx++
		if patternIdx == patternLen {
			patternIdx = 0
		}
	}
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}
