package sim

type Simulation interface {
	Play()
	Round() int
	World() World
}

type Cell = byte

type World interface {
	X() int
	Y() int
	Cells() *[][]Cell
	Clone() World
	Set(int, int, Cell)
	Get(int, int) Cell
	Neighbors(int, int) uint8
	Stats() string
}
