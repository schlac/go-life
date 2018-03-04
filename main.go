package main

import (
	"github.com/schlac/go-life/ui"
)

var (
	version string = "private"
)

func main() {
	ui.PrintHello(version)
}
