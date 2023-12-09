
package main

import (
	"fmt"
)

type PartTwo struct {
}

func NewPartTwo() *PartTwo {
  return &PartTwo {
  }
}

func (p *PartTwo) Init() {
	fmt.Println("Initializing Day 09 Part 2")
}

func (p *PartTwo) ReadLine(line string, idx int) {
  fmt.Println("Reading line", line, idx)
}

func (p *PartTwo) Finish() {
	fmt.Printf("Day 09 Part 2:\n")
	// fmt.Printf("%d:\n", p.Sum)
}
