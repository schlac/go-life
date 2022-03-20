package main

import (
	"flag"
	"log"

	"github.com/schlac/go-life/bs"
	"github.com/schlac/go-life/display"
	"github.com/schlac/go-life/sim"
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
	display.PrintHello(version)
	roundsPtr := flag.Int("rounds", 10, "number of rounds to play")

	flag.Parse()
	args := flag.Args()

	var world sim.World
	if len(args) >= 1 {
		filePath := args[0]
		world = bs.NewSpaceFromFile(filePath)
	} else {
		world = bs.NewRandomSpace(20, 10)
	}
	log.Println(world.Stats())
	s := bs.NewSimulation(&world)
	for i := 1; i <= *roundsPtr; i++ {
		s.Play()
		display.PrintSimulation(s)
	}
}
