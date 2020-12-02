package main

import (
	"fmt"
	"time"
)

func main() {
	messager1, messager2 := make(chan string), make(chan string)
	quit := time.After(5 * time.Second)
	go func() {
		for {
			time.Sleep(time.Second)
			messager1 <- "hello"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			messager2 <- "world"
		}
	}()
	for {
		select {
		case <-quit:
			fmt.Println("finished!")
			return
		case msg := <-messager1:
			fmt.Println(msg)
		case msg := <-messager2:
			fmt.Println(msg)
		}
	}
}
