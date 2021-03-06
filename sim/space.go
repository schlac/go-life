package sim

import (
	"fmt"
	"log"
	"strings"
	"io/ioutil"
)

var (
	border byte = '#'
	pop byte = 'O'
	npop byte = ' '
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
	sp.width = 10
	sp.height = 10
	return sp
}

// Create a new simulation space from file
func NewSpaceFromFile(path string) *Space {
	log.Printf("Loading file '%s'", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln(err)
	}

	sp := NewSpace()
	sp.data = data

	// infer space width
	for i := 0; i < len(sp.data); i++ {
		c := sp.data[i]
		if c == '\r' || c == '\n' {
			sp.width = i + 1
			c = sp.data[i+1]
			if c == '\r' || c == '\n' {
				sp.width++
			}
			break
		}
	}

	// post process read data
	for i := 0; i < len(sp.data); i++ {
		switch sp.data[i] {
		case 0:
		case '\r':
		case '\n':
		case ' ':
			sp.data[i] = 0
		default:
			sp.data[i] = 1
		}
	}
	sp.height = (len(sp.data) / sp.width ) + 1

	return sp
}

// A statistics string
func (s *Space) StatsString() string {
	return fmt.Sprintf("space size: %d x %d", s.width, s.height)
}

// Convert the simulation space to a string
func (s *Space) String() string {
	var b strings.Builder
	// fmt.Fprintf(&b, "array: %v", s.data)
	// b.WriteByte('\n')

	for i := 0; i < s.width + 2; i++ {
		b.WriteByte(border)
	}
	b.WriteByte('\n')

	for i, val := range s.data {
		if i % s.width == 0 {
			b.WriteByte(border)
		}
		if val > 0 {
			b.WriteByte(pop)
		} else {
			b.WriteByte(npop)
		}
		if i % s.width == s.width - 1 {
			b.WriteByte(border)
			b.WriteByte('\n')
		}
	}

	for i := 0; i < s.width + 2; i++ {
		b.WriteByte(border)
	}
	b.WriteByte('\n')

	return b.String()
}

