package main

import (
	"fmt"
	"strings"
)

type PartOne struct {
	Pattern []string
	Map     map[string][]string
}

func NewPartOne() *PartOne {
	return &PartOne{
		Map: make(map[string][]string),
	}
}

func (p *PartOne) Init() {
	fmt.Println("Initializing Day 08 Part 1")
}

func (p *PartOne) ReadLine(line string, idx int) {
	if idx == 0 {
		p.Pattern = strings.Split(line, "")
		return
	}

	if line == "" {
		return
	}

	key, nodes := ParseMapLine(line)

	// if strings.HasSuffix(key, "A") {
	// 	fmt.Printf("key: %v\n", key)
	// }

	p.Map[key] = nodes
}

func ParseMapLine(line string) (key string, nodes []string) {
	parts := strings.Split(line, " = ")

	key = parts[0]
	nodesStr := strings.Trim(parts[1], "()")
	nodes = strings.Split(nodesStr, ", ")
	// left, right := nodeParts[0], nodeParts[1]

	// fmt.Printf("k: %s, l:%s, r:%s\n", key, left, right)
	return
}

func (p *PartOne) Finish() {
	fmt.Printf("Day 08 Part 1:\n")
	// fmt.Printf("p.Map: %+v\n", p.Map)
	// count := p.WalkTree()
	// fmt.Printf("%d\n", count)
}

func (p *PartOne) WalkTree() int {
	patternIdx := 0
	patternLen := len(p.Pattern)
	count := 0

	key := "AAA"

	for {
		count++
		pattern := p.Pattern[patternIdx]
		nodes := p.Map[key]

		fmt.Printf("Pattern: %s - nodes: %+v\n", pattern, nodes)

		if pattern == "L" {
			key = nodes[0]
		}

		if pattern == "R" {
			key = nodes[1]
		}

		if key == "ZZZ" {
			break
		}

		patternIdx++
		if patternIdx == patternLen {
			patternIdx = 0
			// continue
		}
	}

	return count
}
