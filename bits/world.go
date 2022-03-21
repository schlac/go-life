package bits

import (
	"fmt"

	"github.com/schlac/go-life/sim"
)

// Simulation Space
type world struct {
	cells *[][]byte
	x     int
}

const blockSize = 8

// Create a new empty simulation space
func new(x int, y int) *world {
	var b = make([][]byte, y)
	for i := range b {
		nBlocks := (x / blockSize) + 1
		b[i] = make([]byte, nBlocks)
	}
	return &world{&b, x}
}

func (w *world) Clone() sim.World {
	var nw = new(w.X(), w.Y())
	for i := range *nw.cells {
		copy((*nw.cells)[i], (*w.cells)[i])
	}
	return nw
}

func (w *world) Cells() *[][]sim.Cell {
	b := make([][]byte, w.Y())
	for y := range b {
		b[y] = make([]byte, w.X())
		for x := 0; x < w.X(); x++ {
			b[y][x] = w.Get(x, y)
		}
	}
	return &b
}

func (w *world) Y() int {
	return len(*w.cells)
}

func (w *world) X() int {
	return w.x
}

// A statistics string
func (w *world) Stats() string {
	return fmt.Sprintf("space size: %d x %d (%d cells)", w.X(), w.Y(), w.X()*w.Y())
}

func (w *world) Get(x int, y int) sim.Cell {
	//log.Printf("get %d %d %d %08b", x, y, (block & (1 << (x % blockSize))), block)
	if ((*w.cells)[y][x/blockSize] & (1 << (x % blockSize))) == 0 {
		return 0
	} else {
		return 1
	}
}

func (w *world) Set(x int, y int, value sim.Cell) {
	//log.Printf("%04d %d %08b %d", x, value, (*w.cells)[y], w.Get(x, y))
	block := (*w.cells)[y][x/blockSize]
	if value == 0 {
		block &= ^(1 << (x % blockSize))
	} else {
		block |= (1 << (x % blockSize))
	}
	(*w.cells)[y][x/blockSize] = block
	//log.Printf("%04d %d %08b %d %08b", x, value, (*w.cells)[y], w.Get(x, y), block)
}

func (w *world) Neighbors(x int, y int) uint8 {
	var sum uint8 = 0
	if y > 0 {
		if x > 0 {
			sum += w.Get(x-1, y-1)
		}
		sum += w.Get(x, y-1)
		if x < w.X()-1 {
			sum += w.Get(x+1, y-1)
		}
	}

	if x > 0 {
		sum += w.Get(x-1, y)
	}
	if x < w.X()-1 {
		sum += w.Get(x+1, y)
	}

	if y < w.Y()-1 {
		if x > 0 {
			sum += w.Get(x-1, y+1)
		}
		sum += w.Get(x, y+1)
		if x < w.X()-1 {
			sum += w.Get(x+1, y+1)
		}
	}
	return sum
}
