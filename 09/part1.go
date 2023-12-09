package main

import (
	"fmt"
	"strconv"
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
	fmt.Println("Initializing Day 09 Part 1")
}

func (p *PartOne) ReadLine(line string, idx int) {
	ints := parseLine(line)
	fmt.Printf("ints: %v\n", ints)
	next := ints.FindNext()

	p.Sum += next
}

func (p *PartOne) Finish() {
	fmt.Printf("Day 09 Part 1:\n")
	fmt.Printf("%d\n", p.Sum)
}

func parseLine(line string) Sequence {
	parts := strings.Split(line, " ")

	ints := []int{}

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)

		if trimmed == "" {
			continue
		}

		i, err := strconv.Atoi(trimmed)

		if err != nil {
			panic(err)
		}

		ints = append(ints, i)
	}

	return Sequence(ints)
}
