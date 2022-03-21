package bits

import (
	"testing"
)

func BenchmarkNewStep(b *testing.B) {
	var world = NewRandomSpace(500, 500, 1)
	var sim = NewNewSimulation(&world)
	b.ResetTimer()
	for i := 0; i < 500; i++ {
		sim.Play()
	}
}

func BenchmarkCloneStep(b *testing.B) {
	var world = NewRandomSpace(500, 500, 1)
	var sim = NewCloneSimulation(&world)
	b.ResetTimer()
	for i := 0; i < 500; i++ {
		sim.Play()
	}
}
