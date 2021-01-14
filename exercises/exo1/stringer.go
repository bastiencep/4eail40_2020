package main

import "fmt"

// Implement types Rectangle, Circle and Shape
type Rectangle struct {
	Width  int64
	Length int64
}

type Circle struct {
	Radius int64
}

func (b Rectangle) String() string {
	return fmt.Sprintf("Rectangle of width %d and length %d", b.Width, b.Length)
}

func (b Circle) String() string {
	return fmt.Sprintf("Circle of radius %d", b.Radius)
}

type Shape interface {
	String() string
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
	}
}
