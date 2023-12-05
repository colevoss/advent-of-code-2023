package main

import (
	"fmt"
	"strings"
)

type MapRange struct {
	SourceStart int
	DestStart   int
	Range       int
}

func (mr *MapRange) MapSeed(seed int) (mapped int, inRange bool) {
	// if source start is 50 and range 2. seed must be 50, or 51
	// 51 >= 50 && 50 < (50 + 2)
	inRange = seed >= mr.SourceStart && seed < mr.SourceStart+mr.Range

	if !inRange {
		return seed, inRange
	}

	diff := mr.SourceStart - mr.DestStart
	return seed - diff, inRange
}

type Map struct {
	Source      string
	Destination string
	Ranges      []*MapRange
}

func (mr *Map) MapSeed(seed int) int {
	for _, mapRange := range mr.Ranges {
		out, inRange := mapRange.MapSeed(seed)

		if inRange {
			return out
		}
	}

	return seed
}

type SeedMap []*Map

func (m *Map) String() string {
	str := fmt.Sprintf("%s to %s\n", m.Source, m.Destination)

	for _, r := range m.Ranges {
		str += fmt.Sprintf("  %d -> %d (%d)\n", r.SourceStart, r.DestStart, r.Range)
	}

	return str
}

func NewMap(source, dest string) *Map {
	return &Map{
		Source:      source,
		Destination: dest,
		Ranges:      []*MapRange{},
	}
}

func parseRangeString(str string) *MapRange {
	numArr := splitNumStr(str)

	destStart := numArr[0]
	sourceStart := numArr[1]
	rnge := numArr[2]

	return &MapRange{
		DestStart:   destStart,
		SourceStart: sourceStart,
		Range:       rnge,
	}
}

func parseMap(line string) *Map {
	split := strings.Split(line, " ")
	mapName := split[0]

	mapParts := strings.Split(mapName, "-to-")

	source := mapParts[0]
	destination := mapParts[1]

	return NewMap(source, destination)
}
