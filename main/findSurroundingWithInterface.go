package main

import "math"

type Circle struct {
	r float64
}

type Rectangle struct {
	w float64
	h float64
}

type Shape interface {
	perimeter() float64
}

func Perimeter(shapes ...Shape) float64 {
	var total float64
	for _, v := range shapes {
		total = +v.perimeter()
	}

	return total

}

func (c *Circle) circlePeri() float64 {
	return math.Pi * c.r * 2.0
}

func (r *Rectangle) rectanglePeri() float64 {
	return (r.h + r.w) * 2
}

func main() {
	cPtr := new(Circle)
	rPtr := new(Rectangle)
	cPtr.r = 1.0
	rPtr.w = 1.0
	rPtr.h = 1.0

}
