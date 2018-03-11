package main

import (
	"github.com/schlac/go-life/sim"
	"github.com/schlac/go-life/ui"
	"log"
	"os"
)

var (
	// version string to be injected on build
	version string = "private"
)

// main entry point
// |-- print hello
// |-- {OR}
// |   |-- generate random simulation space
// |   |-- read simulation space from file
// |-- run simulation
func main() {
	ui.PrintHello(version)
	if len(os.Args) < 2 {
		log.Panicln("Arg file missing")
	}

	filePath := os.Args[1]
	sp := sim.NewSpaceFromFile(filePath)
	si := sim.NewSimulation(sp)
	si.Play()
}
