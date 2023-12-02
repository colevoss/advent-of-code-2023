package main

import (
	"aoc/runner"
	"fmt"
	"strconv"
	"strings"
)

type DayOne struct {
	Nums []int
	Sum  int
}

func (d *DayOne) Init() {
	fmt.Println("Initializing Day One")
}

func (d *DayOne) ReadLine(line string, idx int) {
	fmt.Printf("Reading line %d: %s\n", idx, line)

	digits := []string{}

	for _, l := range strings.Split(line, "") {
		if _, err := strconv.Atoi(l); err == nil {
			fmt.Printf("Found digit %s\n", l)
			digits = append(digits, l)
		}
	}

	first := digits[0]
	last := digits[len(digits)-1]
	numStr := first + last

	num, err := strconv.Atoi(numStr)

	if err != nil {
		panic(err)
	}

	d.Sum += num
}

func (d *DayOne) Finish() {
}

func main() {
	runner := runner.NewRunner("dayone/input.txt")

	test := &DayOne{
		Nums: []int{},
	}

	runner.Run(test)

	fmt.Printf("Sum: %d\n", test.Sum)
}
