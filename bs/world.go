package bs

import (
	"fmt"

	"github.com/schlac/go-life/sim"
)

// Simulation Space
type world struct {
	cells *[][]byte
}

// Create a new empty simulation space
func new(x int, y int) *world {
	var b = make([][]byte, y)
	for i := range b {
		b[i] = make([]byte, x)
	}
	return &world{&b}
}

func (w *world) Clone() sim.World {
	var nw = new(w.X(), w.Y())
	for i := range *nw.cells {
		copy((*nw.cells)[i], (*w.cells)[i])
	}
	return nw
}

func (w *world) Cells() *[][]sim.Cell {
	return w.cells
}

func (w *world) Y() int {
	return len(*w.cells)
}

func (w *world) X() int {
	return len((*w.cells)[0])
}

// A statistics string
func (w *world) Stats() string {
	return fmt.Sprintf("space size: %d x %d (%d cells)", w.X(), w.Y(), w.X()*w.Y())
}

func (w *world) Get(x int, y int) sim.Cell {
	return (*w.cells)[y][x]
}

func (w *world) Set(x int, y int, value sim.Cell) {
	(*w.cells)[y][x] = value
}

func (w *world) Neighbors(x int, y int) uint8 {
	var sum uint8 = 0
	if y > 0 {
		if x > 0 {
			sum += (*w.cells)[y-1][x-1]
		}
		sum += (*w.cells)[y-1][x]
		if x < w.X()-1 {
			sum += (*w.cells)[y-1][x+1]
		}
	}

	if x > 0 {
		sum += (*w.cells)[y][x-1]
	}
	if x < w.X()-1 {
		sum += (*w.cells)[y][x+1]
	}

	if y < w.Y()-1 {
		if x > 0 {
			sum += (*w.cells)[y+1][x-1]
		}
		sum += (*w.cells)[y+1][x]
		if x < w.X()-1 {
			sum += (*w.cells)[y+1][x+1]
		}
	}
	return sum
}
