package main

import (
	"github.com/schlac/go-life/sim"
	"github.com/schlac/go-life/ui"
	"log"
	"os"
)

var (
	version string = "private"
)

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
