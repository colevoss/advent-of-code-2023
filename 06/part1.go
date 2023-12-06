package main

import (
	"fmt"
	"strconv"
	"strings"
)

type PartOne struct {
	Times     []int
	Distances []int
}

func NewPartOne() *PartOne {
	return &PartOne{}
}

func (p *PartOne) Init() {
	fmt.Println("Initializing Day 06 Part 1")
}

func (p *PartOne) ReadLine(line string, idx int) {
	// idx 0 is times
	// idx 1 is distances

	if idx == 0 {
		times := p.ParseLine(line)
		p.Times = times
		fmt.Printf("times: %+v\n", times)
	}

	if idx == 1 {
		distances := p.ParseLine(line)
		p.Distances = distances
		fmt.Printf("distances: %+v\n", distances)
	}
}

func (p *PartOne) Finish() {
	fmt.Printf("Day 06 Part 1:\n")
	solutions := []int{}

	for i := 0; i < len(p.Times); i++ {
		count := p.GetButtonTimeForRace(i)
		solutions = append(solutions, count)
	}

	product := 1
	for _, s := range solutions {
		product = product * s
	}

	fmt.Printf("%d\n", product)
}

func (p *PartOne) ParseLine(line string) []int {
	split := strings.Split(line, ":")
	timesStr := strings.TrimSpace(split[1])
	timesStrArr := strings.Split(timesStr, " ")

	// timeNumsArr := make([]int, len(timesStrArr))
	timeNumsArr := []int{}

	for _, ts := range timesStrArr {
		num, err := strconv.Atoi(strings.TrimSpace(ts))

		if err != nil {
			continue
		}

		timeNumsArr = append(timeNumsArr, num)
	}

	return timeNumsArr
}

func (p *PartOne) GetButtonTimeForRace(raceIdx int) int {
	time := p.Times[raceIdx]
	targetDistance := p.Distances[raceIdx]

	return getSolutionsForRace(time, targetDistance)
}

func getSolutionsForRace(time int, targetDistance int) int {
	solutionCount := 0

	for i := 0; i <= time; i++ {
		// i == speed
		speed := i
		remainingTime := time - i

		distanceTravelled := speed * remainingTime

		if distanceTravelled > targetDistance {
			solutionCount++
		}
	}

	return solutionCount
}
