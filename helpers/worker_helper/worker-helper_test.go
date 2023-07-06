package worker_helper

import (
	"testing"
	"time"
)

func TestInstantWorker(t *testing.T) {
	stop := make(chan bool)

	var functionCalled bool

	myFunc := func() {
		functionCalled = true
	}

	go InstantWorker(myFunc, 100*time.Millisecond, stop)

	time.Sleep(500 * time.Millisecond)

	stop <- true

	if !functionCalled {
		t.Errorf("Function was not called by InstantWorker")
	}
}
