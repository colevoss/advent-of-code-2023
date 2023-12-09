package main

type Sequence []int

func (s Sequence) FindNext() int {
	reduced := s.Reduce()

	toAdd := 0
	for i := len(reduced) - 1; i >= 0; i-- {
		seq := reduced[i]
		last := seq[len(seq)-1]
		toAdd += last
	}

	next := toAdd + s[len(s)-1]
	return next
}

func (s Sequence) Reduce() [][]int {
	seqs := [][]int{}

	done := false
	last := s

	for !done {
		red, fullyReduced := last.ReduceOnce()
		seqs = append(seqs, red)
		last = red

		done = fullyReduced
	}

	return seqs
}

func (s Sequence) ReduceOnce() (Sequence, bool) {
	reduced := make([]int, len(s)-1)
	fullyReduced := true

	for i := 0; i < len(s)-1; i++ {
		a := s[i]
		b := s[i+1]

		diff := b - a
		reduced[i] = diff

		if diff != 0 {
			fullyReduced = false
		}
	}

	return reduced, fullyReduced
}
