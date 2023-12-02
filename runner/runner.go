package runner

import (
	"bufio"
	"os"
)

type Runner struct {
	Path string
}

func NewRunner(path string) *Runner {
	return &Runner{path}
}

type Challenge interface {
	Init()
	ReadLine(line string, idx int)
	Finish()
}

func (r *Runner) Run(c Challenge) {
	f, err := os.Open(r.Path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	idx := 0

	c.Init()

	for scanner.Scan() {
		txt := scanner.Text()

		c.ReadLine(txt, idx)
		idx++
	}

	c.Finish()
}
