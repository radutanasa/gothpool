package gothpool

import (
	"math/rand"
	"testing"
	"time"
)

func TestExecPool(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	exec := New(4, 1000)
	exec.Start()
	defer exec.Stop()

	for i := 0; i < 10000; i++ {
		exec.Run(func() {
			PrintWithDelay(rand.Intn(100))
		})
	}
}

func PrintWithDelay(delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	println(delay)
}