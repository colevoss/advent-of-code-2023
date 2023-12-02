package main

type Game struct {
	Id     int
	Rounds []*Round
}

func NewGame(id int) *Game {
	return &Game{
		Id:     id,
		Rounds: []*Round{},
	}
}

func (g *Game) AddRound(round *Round) {
	g.Rounds = append(g.Rounds, round)
}

func (g *Game) IsValid() bool {
	for _, r := range g.Rounds {
		if !r.IsValid() {
			return false
		}
	}

	return true
}

func (g *Game) MinCubeRound() *Round {
	minCubeRound := &Round{}

	for _, r := range g.Rounds {
		round := r

		minCubeRound.Red = max(round.Red, minCubeRound.Red)
		minCubeRound.Green = max(round.Green, minCubeRound.Green)
		minCubeRound.Blue = max(round.Blue, minCubeRound.Blue)
	}

	return minCubeRound
}

func max(intA int, intB int) int {
	if intA > intB {
		return intA
	} else {
		return intB
	}
}

type Round struct {
	Red   int
	Green int
	Blue  int
}

func (r *Round) AddCubes(count int, color string) {
	switch color {
	case "red":
		r.Red += count
		break
	case "blue":
		r.Blue += count
		break
	case "green":
		r.Green += count
		break
	}
}

func (r *Round) IsValid() bool {
	if r.Red > 12 || r.Green > 13 || r.Blue > 14 {
		return false
	}

	return true
}

func (r *Round) Product() int {
	return r.Red * r.Blue * r.Green
}
