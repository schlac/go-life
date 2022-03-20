package bs

import (
	"math/rand"

	"github.com/schlac/go-life/sim"
)

// Create a new random simulation space
func NewRandomSpace(width int, height int) sim.World {
	var w = new(width, height)
	for i := range *w.cells {
		for j, r := range rand.Perm(width) {
			(*w.cells)[i][j] = byte(r % 2)
		}
	}
	return w
}
