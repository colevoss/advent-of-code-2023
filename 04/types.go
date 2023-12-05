package main

import (
	"regexp"
	"strings"
)

var cardReplaceReg = regexp.MustCompile(`Card\s\d+:\s`)
var splitReg = regexp.MustCompile(`\s+`)

type NumList []string

type Card struct {
	Id             int
	WinningNumbers NumList
	MyNumbers      NumList
	WinCount       int
	Copies         int
}

func (n NumList) HasNumber(num string) bool {
	for _, listNum := range n {
		if listNum == num {
			return true
		}
	}

	return false
}

func NewCard(line string, id int) *Card {
	noCard := cardReplaceReg.ReplaceAllString(line, "")
	split := strings.Split(noCard, "|")

	winningNumStr := split[0]
	myNumStr := split[1]

	winningNums := NumList(createNumArr(winningNumStr))
	myNums := NumList(createNumArr(myNumStr))

	return &Card{
		Id:             id,
		WinningNumbers: winningNums,
		MyNumbers:      myNums,
		WinCount:       0,
		Copies:         1, // every card gets at least one copy (the original copy)
	}
}

func (c *Card) Process() {
	for _, myNum := range c.MyNumbers {
		isMatched := c.WinningNumbers.HasNumber(myNum)

		if isMatched {
			c.WinCount++
		}
	}
}

func (c *Card) Score() int {
	if c.WinCount == 0 {
		return 0
	}

	return 1 << (c.WinCount - 1)
}
