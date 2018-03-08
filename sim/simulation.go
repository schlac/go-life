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
	return s
}

func (s *Simulation) print() {
	log.Printf("%i:\n", s.round)
	log.Println(s.space.ToString())
}

func (s *Simulation) Play() {
	log.Printf("Starting round %i", s.round)
	s.print()
}

