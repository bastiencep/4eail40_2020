package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var shared1 int
var mux1 sync.Mutex
var shared2 int
var mux2 sync.Mutex

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementSharedVariable()
		wg.Add(1)
		go decrementSharedVariable()
	}
	wg.Wait()
	fmt.Println(shared1)
	fmt.Println(shared2)
}

func incrementSharedVariable() {
	mux1.Lock()
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	mux2.Lock()
	shared1++
	shared2++
	mux1.Unlock()
	mux2.Unlock()
	wg.Done()
}

func decrementSharedVariable() {
	mux2.Lock()
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	mux1.Lock()
	shared1--
	shared2--
	mux1.Unlock()
	mux2.Unlock()
	wg.Done()
}
