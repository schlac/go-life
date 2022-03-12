package sim

type Simulation struct {
	Space Space
	Round int
}

func NewSimulation(sp Space) Simulation {
	return Simulation{sp, 0}
}

func (s *Simulation) Play() {
	s.Round++
	var next = s.Space.Clone()
	for y, row := range s.Space.cells {
		for x, cell := range row {
			n := s.Space.Neighbors(x, y)
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
	s.Space = next
}
