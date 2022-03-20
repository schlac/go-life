package bs

import (
	"io/ioutil"
	"log"

	"github.com/schlac/go-life/sim"
)

// Create a new simulation space from file
func NewSpaceFromFile(path string) sim.World {
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
	var w = new(len(rows), maxLength)
	for i, row := range rows {
		copy((*w.cells)[i], data[row.start:row.start+row.length])
		for j, c := range (*w.cells)[i] {
			(*w.cells)[i][j] = c % 2
		}
	}

	return w
}
