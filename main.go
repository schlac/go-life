package main

import (
	"flag"
	"log"
	"time"
	"io/ioutil"

	bs "github.com/schlac/go-life/bits"
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
	seedPtr := flag.Int64("seed", 0, "seed to use for random numbers")
	verbosePtr := flag.Bool("verbose", false, "logging")

	flag.Parse()
	args := flag.Args()

	if !*verbosePtr {
		log.SetOutput(ioutil.Discard)
	}

	var world sim.World
	if len(args) >= 1 {
		filePath := args[0]
		var err error
		world, err = bs.NewSpaceFromFile(filePath)
		if err != nil {
			panic(err)
		}
	} else {
		var seed int64
		if *seedPtr != 0 {
			seed = *seedPtr
		} else {
			seed = time.Now().UnixNano()
		}
		world = bs.NewRandomSpace(20, 10, seed)
	}
	log.Println(world.Stats())
	s := bs.NewCloneSimulation(&world)
	display.PrintSimulation(s)
	for i := 1; i <= *roundsPtr; i++ {
		s.Play()
		display.PrintSimulation(s)
	}
}
