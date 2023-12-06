package main

import (
	"fmt"
	"strconv"
	"strings"
)

type PartTwo struct {
	Time     int
	Distance int
}

func NewPartTwo() *PartTwo {
	return &PartTwo{}
}

func (p *PartTwo) Init() {
	fmt.Println("Initializing Day 06 Part 2")
}

func (p *PartTwo) ReadLine(line string, idx int) {
	fmt.Println("Reading line", line, idx)

	if idx == 0 {
		time, err := p.ParseLine(line)

		if err != nil {
			panic(err)
		}

		p.Time = time
	}

	if idx == 1 {
		distance, err := p.ParseLine(line)

		if err != nil {
			panic(err)
		}

		p.Distance = distance
	}
}

func (p *PartTwo) Finish() {
	fmt.Printf("Day 06 Part 2:\n")
	solutions := getSolutionsForRace(p.Time, p.Distance)
	fmt.Printf("%d\n", solutions)
}

func (p *PartTwo) ParseLine(line string) (int, error) {
	split := strings.Split(line, ":")
	timesStr := strings.TrimSpace(split[1])
	timesStrArr := strings.Split(timesStr, " ")

	combinedNumStr := ""

	for _, str := range timesStrArr {
		trimmed := strings.TrimSpace(str)
		combinedNumStr += trimmed
	}

	return strconv.Atoi(combinedNumStr)
}
