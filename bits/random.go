package bs

import (
	"math/rand"

	"github.com/schlac/go-life/sim"
)

// Create a new random simulation space
func NewRandomSpace(width int, height int, seed int64) sim.World {
	if seed != 0 {
		rand.Seed(seed)
	}
	var w = new(width, height)
	for i := range *w.cells {
		for j, r := range rand.Perm(width) { // TODO use rand.Read()
			(*w.cells)[i][j] = byte(r % 2)
		}
	}
	return w
}
