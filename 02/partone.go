package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PartOne struct {
	Sum         int
	gameRegexp  *regexp.Regexp
	roundRegexp *regexp.Regexp
}

func NewPartOne() *PartOne {
	reg := regexp.MustCompile("Game\\s(\\d+):\\s(.*)")
	roundReg := regexp.MustCompile("(\\d+)\\s(red|green|blue)")

	return &PartOne{
		Sum:         0,
		gameRegexp:  reg,
		roundRegexp: roundReg,
	}
}

func (p *PartOne) Init() {
	fmt.Printf("Running Day Two Part One\n")
}

func (p *PartOne) ReadLine(line string, idx int) {
	game := p.ParseGame(line)

	if !game.IsValid() {
		return
	}

	p.Sum += game.Id
}

func (p *PartOne) ParseGame(line string) *Game {
	gameNr, cubeStr, err := p.GetGame(line)

	if err != nil {
		panic(err)
	}

	game := NewGame(gameNr)

	rounds := strings.Split(cubeStr, ";")

	for _, round := range rounds {
		round := p.ParseRound(round)
		game.AddRound(round)
	}

	return game
}

func (p *PartOne) Finish() {
	fmt.Printf("Day Two Part One: %d\n", p.Sum)
}

func (p *PartOne) GetGame(line string) (int, string, error) {
	found := p.gameRegexp.FindStringSubmatch(line)

	gameNumStr := found[1]
	cubeStr := found[2]

	i, err := strconv.Atoi(gameNumStr)

	return i, cubeStr, err
}

func (p *PartOne) ParseRound(roundStr string) *Round {
	matches := p.roundRegexp.FindAllStringSubmatch(roundStr, -1)

	round := &Round{}

	for _, m := range matches {
		countStr := m[1]
		color := m[2]

		count, err := strconv.Atoi(countStr)

		if err != nil {
			panic(err)
		}

		round.AddCubes(count, color)

		// fmt.Printf("%d %s - %s\n", i, countStr, color)
	}

	return round
}
