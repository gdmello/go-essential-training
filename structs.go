package main

import (
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func (p *Point) Move(dx int, dy int) {
	p.x += dx
	p.y += dy
}

type Square struct {
	center Point
	length int
}

func (s *Square) Move(dx int, dy int) {
	s.center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.length * s.length
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if x == 0 {
		return nil, fmt.Errorf("x co-ordinate missing")
	}
	if y ==0 {
		return nil, fmt.Errorf("y co-ordinate missing")
	}
	if length == 0 {
		return nil, fmt.Errorf("length missing")
	}
	p := Point{x, y}
	s := &Square{p, length}
	return s, nil
}

func main() {
	s, error := NewSquare(1,1,10)
	if error != nil {
		os.Exit(1)
	}
	s.Move(2,3)
	fmt.Printf("Move -> %+v\n",s)
	fmt.Printf("Area -> %+v\n",s.Area())
}