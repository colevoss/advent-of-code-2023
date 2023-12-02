package main

import (
	"fmt"
	"strconv"
	"strings"
)

type PartOne struct {
	Nums []int
	Sum  int
}

func (d *PartOne) Init() {
	fmt.Println("Initializing Day One")
}

func (d *PartOne) ReadLine(line string, idx int) {
	digits := []string{}

	for _, l := range strings.Split(line, "") {
		if _, err := strconv.Atoi(l); err == nil {
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

func (d *PartOne) Finish() {
	fmt.Printf("Day One Part One Solution:\n%d\n", d.Sum)
}
