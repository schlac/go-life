package main

import (
	"flag"
	"log"

	"github.com/schlac/go-life/sim"
	"github.com/schlac/go-life/ui"
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
	roundsPtr := flag.Int("rounds", 10, "number of rounds to play")

	flag.Parse()
	args := flag.Args()

	var sp sim.Space
	if len(args) >= 1 {
		filePath := args[0]
		sp = sim.NewSpaceFromFile(filePath)
		log.Println(sp.StatsString())
	} else {
		sp = sim.NewRandomSpace()
	}
	si := sim.NewSimulation(sp)
	for i := 1; i <= *roundsPtr; i++ {
		si.Play()
		ui.PrintSimulation(si)
	}
}
