package main

import (
	"fmt"
	"sync"
)

// Fixing the data race.

var shared int
var mux sync.Mutex

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeSharedVariable()
	}
	wg.Wait() // we need to wait for all other goroutines to complete
	fmt.Print(shared)
}

func writeSharedVariable() {
	mux.Lock()
	shared++ // this is the critical section and should be protected
	mux.Unlock()
	wg.Done()
}
