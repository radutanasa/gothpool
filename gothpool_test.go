package gothpool

import (
	"math/rand"
	"testing"
	"time"
)

func TestExecPool(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	exec := New(4, 10e5)
	exec.Start()

	for i := 0; i < 20; i++ {
		var j = i
		err := exec.Run(func() {
			PrintValue(j)
		})
		if err != nil {
			t.Error("Pool should not be stopped at this point.", err)
		}
	}

	exec.Stop()

	err := exec.Run(func() {
		t.Error("This should never run as the pool is stopped.")
	})
	if err == ExecPoolStoppedErr {
		t.Log("Correct behavior - pool is stopped and can't accept new jobs.")
	}
}

func PrintValue(value int) {
	println(value)
}