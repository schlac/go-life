package sim

import (
	"log"
)

type Simulation struct {
	space *Space
	round int
}

func NewSimulation(sp *Space) *Simulation {
	s := new(Simulation)
	s.space = sp
	s.round = 0
	s.print()
	return s
}

func (s *Simulation) print() {
	log.Printf("%d:\n%s", s.round, s.space.String())
}

func (s *Simulation) Play() {
	s.round++
	s.print()
}
