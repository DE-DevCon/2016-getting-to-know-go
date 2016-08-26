package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func (pt point) hypotenuse() float64 {
	return math.Sqrt(float64(pt.x)*float64(pt.x) + float64(pt.y)*float64(pt.y))
}

func main() {
	pt := point{x: 3, y: 5}
	fmt.Println(pt.hypotenuse())
}
