package bits

import (
	"math/rand"

	"github.com/schlac/go-life/sim"
)

// Create a new random simulation space
func NewRandomSpace(width int, height int, seed int64) sim.World {
	if seed != 0 {
		rand.Seed(seed)
	}

	var b = make([][]byte, height)
	for i := range b {
		var row = make([]byte, (width/blockSize)+1)
		rand.Read(row)
		//log.Printf("%08b", row)
		b[i] = row
	}
	var w sim.World = &world{&b, width}
	return w
}
