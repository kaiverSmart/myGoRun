package main

import "fmt"

type Shaper interface {
	Area() float64
}

type square struct {
	side float64
}

func (sq *square) Area() float64  {
	return sq.side * sq.side
}

func main() {
	sq :=  square{side:23}
	var sint Shaper
	sint = &sq
	 fmt.Println(sint.Area())
}