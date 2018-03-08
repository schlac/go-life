package sim

import (
	"log"
	"os"
)

// Simulation Space
type Space struct {
	width int
	height int
	data []byte
}

// Create a new simulation space
func NewSpace() *Space {
	sp := new(Space)
	sp.width = 80
	sp.height = 80
	return sp
}

// Create a new simulation space from file
func NewSpaceFromFile(path string) *Space {
	log.Printf("Loading file '%s'", path)
	sp := NewSpace()

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := f.Read(sp.data)
	// post process read data
	for i := 0; i < n; i++ {
		c := sp.data[i]
		if c == 0 || c == ' ' {
			sp.data[i] = 0
		} else {
			sp.data[i] = 1
		}
	}
	sp.height = (n / sp.width ) + 1

	return sp
}

func (s *Space) ToString() string {
	ret := ""
	for i, val := range s.data {
		if i % s.width == 0 {
			ret += "\n"
		}
		if val > 0 {
			ret += "#"
		} else {
			ret += " "
		}
	}
	return ret
}

