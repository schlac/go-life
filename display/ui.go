package display

import (
	"fmt"
	"strings"

	"github.com/schlac/go-life/sim"
)

const (
	border byte = '#'
	pop    byte = 'o'
	npop   byte = ' '
)

func PrintHello(version string) {
	fmt.Printf("go-life version %s\n", version)
}

func PrintSimulation(s sim.Simulation) {
	fmt.Printf("\n%d:\n", s.Round())
	PrintSpace(s.World())
}

func PrintSpace(s sim.World) {
	fmt.Print(String(s))
}

// Convert the simulation space to a string
func String(s sim.World) string {
	var b strings.Builder

	printHLine := func() {
		for i := 0; i < s.X()+2; i++ {
			b.WriteByte(border)
		}
		b.WriteByte('\n')
	}

	printHLine()
	for _, row := range *s.Cells() {
		b.WriteByte(border)
		for _, cell := range row {
			if cell == 0 {
				b.WriteByte(npop)
			} else {
				b.WriteByte(pop)
			}
		}
		b.WriteByte(border)
		b.WriteByte('\n')
	}
	printHLine()

	return b.String()
}
