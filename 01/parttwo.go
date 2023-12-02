package main

import (
	"fmt"
	"strconv"
	"strings"
)

type PartTwo struct {
	Sum int
}

func (d *PartTwo) Init() {
	fmt.Println("Initializing Day One Part Two")
}

func (d *PartTwo) ReadLine(line string, idx int) {
	num := d.PraseLine(line)
	d.Sum += num
}

func (d *PartTwo) PraseLine(line string) int {
	first := 0
	last := 0

	split := strings.Split(line, "")
	reverse := rev(split)

	digitWord := ""
	for _, s := range split {
		digitWord += s
		digit, ok := getDigitFromString(digitWord)

		if !ok {
			continue
		}

		first = digit
		break
	}

	// Once we have the first digit, we can search for the last
	// digit by parsing from the tail for cases like 3oneeightwo that should be 32 instead of 38
	digitWord = ""
	for _, s := range reverse {
		digitWord = s + digitWord
		digit, ok := getDigitFromString(digitWord)

		if !ok {
			continue
		}

		last = digit
		break
	}

	if last == 0 {
		last = first
	}

	num, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))

	if err != nil {
		panic(err)
	}

	return num
}

func rev(strs []string) []string {
	str := ""

	for _, s := range strs {
		str = s + str
	}

	return strings.Split(str, "")
}

func (d *PartTwo) Finish() {
	fmt.Printf("Day One Part Two Solution:\n%d\n", d.Sum)
}

func getDigitFromString(str string) (int, bool) {
	for k, i := range numMap {
		if strings.Contains(str, k) {
			return i, true
		}
	}

	return -1, false
}

var numMap = map[string]int{
	"one":   1,
	"1":     1,
	"two":   2,
	"2":     2,
	"three": 3,
	"3":     3,
	"four":  4,
	"4":     4,
	"five":  5,
	"5":     5,
	"six":   6,
	"6":     6,
	"seven": 7,
	"7":     7,
	"eight": 8,
	"8":     8,
	"nine":  9,
	"9":     9,
}
