package bs

import "github.com/schlac/go-life/sim"

type stepper func(*bs)

type bs struct {
	round    int
	world    *sim.World
	stepFunc stepper
}

func NewSimulation(w *sim.World) sim.Simulation {
	return &bs{0, w, Play}
}

func (s *bs) Round() int {
	return s.round
}

func (s *bs) World() sim.World {
	return *s.world
}

func (s *bs) Play() {
	s.stepFunc(s)
}

func Play(s *bs) {
	s.round++
	var current = *s.world
	var next = current.Clone()
	for y, row := range *current.Cells() {
		for x, cell := range row {
			n := current.Neighbors(x, y)
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
	s.world = &next
}
