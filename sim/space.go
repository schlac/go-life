package sim

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
)

// Simulation Space
type Space struct {
	cells [][]byte
}

func (s *Space) GetCells() [][]byte {
	return s.cells
}

func (s *Space) Clone() Space {
	var cells = make([][]byte, len(s.cells))
	for i := range cells {
		cells[i] = make([]byte, len(s.cells[0]))
		copy(cells[i], s.cells[i])
	}
	return Space{cells}
}

// Create a new simulation space
func NewRandomSpace() Space {
	width := 20
	height := 10
	var sp = Space{make([][]byte, height)}
	for i := range sp.cells {
		sp.cells[i] = make([]byte, width)
		for j, r := range rand.Perm(width) {
			sp.cells[i][j] = byte(r % 2)
		}
	}
	return sp
}

// Create a new simulation space from file
func NewSpaceFromFile(path string) Space {
	log.Printf("Loading file '%s'", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln(err)
	}

	type row struct {
		start  int
		length int
	}
	var rows []row = make([]row, 0, 100)
	var rowStart int = 0
	var maxLength = 0
	for i, c := range data {
		if c == '\n' || c == '\r' {
			var length = i - rowStart
			if length > 0 {
				rows = append(rows, row{rowStart, length})
			}
			rowStart = i + 1
			if maxLength < length {
				maxLength = length
			}
		}
	}

	// read data into slices
	var sp = Space{make([][]byte, len(rows))}
	for i, row := range rows {
		sp.cells[i] = make([]byte, maxLength)
		copy(sp.cells[i], data[row.start:row.start+row.length])
		for j, c := range sp.cells[i] {
			sp.cells[i][j] = c % 2
		}
	}

	return sp
}

func (s *Space) Y() int {
	return len(s.cells)
}

func (s *Space) X() int {
	return len(s.cells[0])
}

// A statistics string
func (s *Space) StatsString() string {
	return fmt.Sprintf("space size: %d x %d (%d cells)", s.X(), s.Y(), s.X()*s.Y())
}

func (s *Space) Get(x int, y int) byte {
	return s.cells[y][x]
}

func (s *Space) Set(x int, y int, value byte) {
	s.cells[y][x] = value
}

func (s *Space) Neighbors(x int, y int) byte {
	var sum byte = 0
	if y > 0 {
		if x > 0 {
			sum += s.cells[y-1][x-1]
		}
		sum += s.cells[y-1][x]
		if x < s.X()-1 {
			sum += s.cells[y-1][x+1]
		}
	}

	if x > 0 {
		sum += s.cells[y][x-1]
	}
	if x < s.X()-1 {
		sum += s.cells[y][x+1]
	}

	if y < s.Y()-1 {
		if x > 0 {
			sum += s.cells[y+1][x-1]
		}
		sum += s.cells[y+1][x]
		if x < s.X()-1 {
			sum += s.cells[y+1][x+1]
		}
	}
	return sum
}
