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
	var next = s.space.Clone()
	for y := 0; y < s.space.height; y++ {
		for x := 0; x < s.space.width; x++ {
			switch s.space.Neighbors(x, y) {
			case 0, 1, 4, 5, 6, 7, 8:
				next.Set(x, y, 0)
			case 3:
				next.Set(x, y, 1)
			case 2:
			default:
			}
		}
	}
	s.space = next
	s.print()
}
