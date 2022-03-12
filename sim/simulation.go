package sim

import (
	"log"
)

type Simulation struct {
	space Space
	round int
}

func NewSimulation(sp Space) Simulation {
	var s = Simulation{sp, 0}
	s.print()
	return s
}

func (s *Simulation) print() {
	log.Printf("%d:\n%s", s.round, s.space.String())
}

func (s *Simulation) Play() {
	s.round++
	var next = s.space.Clone()
	for y, row := range s.space.cells {
		for x, cell := range row {
			n := s.space.Neighbors(x, y)
			if cell > 0 {
				if n != 2 && n != 3 {
					next.Set(x, y, 0)
				}
			} else {
				if n == 3 {
					next.Set(x, y, 1)
				}
			}
		}
	}
	s.space = next
	s.print()
}
