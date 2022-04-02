package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Length() float64
}

type Rect struct {
	width float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Length() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	var s Shape

	s = Rect{4.0, 5.0}
	
	fmt.Println("Area is:", s.Area())
}
