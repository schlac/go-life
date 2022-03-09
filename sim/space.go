package sim

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

const (
	Border byte = '#'
	Pop    byte = 'X'
	NPop   byte = ' '
)

// Simulation Space
type Space struct {
	width  int
	height int
	data   []byte
}

func (s *Space) Clone() *Space {
	data := make([]byte, len(s.data))
	copy(data, s.data)
	return &Space{s.width, s.height, data}
}

// Create a new simulation space
func NewRandomSpace() *Space {
	sp := new(Space)
	sp.width = 20
	sp.height = 10
	sp.data = make([]byte, sp.width*sp.height)
	for i, v := range rand.Perm(sp.width * sp.height) {
		sp.data[i] = byte(v % 2)
	}
	return sp
}

// Create a new simulation space from file
func NewSpaceFromFile(path string) *Space {
	log.Printf("Loading file '%s'", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln(err)
	}

	sp := new(Space)
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
		case 0, '\r', '\n', ' ':
			sp.data[i] = 0
		default:
			sp.data[i] = 1
		}
	}
	sp.height = (len(sp.data) / sp.width)

	return sp
}

// A statistics string
func (s *Space) StatsString() string {
	return fmt.Sprintf("space size: %d x %d (%d cells)", s.width, s.height, s.width*s.height)
}

// Convert the simulation space to a string
func (s *Space) String() string {
	var b strings.Builder
	// fmt.Fprintf(&b, "array: %v", s.data)
	// b.WriteByte('\n')

	for i := 0; i < s.width+2; i++ {
		b.WriteByte(Border)
	}
	b.WriteByte('\n')

	for i, val := range s.data {
		if i%s.width == 0 {
			b.WriteByte(Border)
		}
		if val > 0 {
			b.WriteByte(Pop)
		} else {
			b.WriteByte(NPop)
		}
		if i%s.width == s.width-1 {
			b.WriteByte(Border)
			b.WriteByte('\n')
		}
	}

	for i := 0; i < s.width+2; i++ {
		b.WriteByte(Border)
	}
	b.WriteByte('\n')

	return b.String()
}

func (s *Space) Get(x int, y int) byte {
	return s.data[x*y+x]
}

func (s *Space) Set(x int, y int, value byte) {
	s.data[x*y+x] = value
}

func (s *Space) Neighbors(x int, y int) int {
	return 0
}
