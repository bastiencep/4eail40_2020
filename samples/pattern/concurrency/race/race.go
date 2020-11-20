package main

import "fmt"

// This code contains a data race.
// Try to run go's race detector on this one. (go run -race)

var shared int

func main() {
	for i := 0; i < 10; i++ {
		go writeSharedVariable()
	}
	fmt.Print(shared)
}

func writeSharedVariable() {
	shared++ // this is the critical section and should be protected
}
