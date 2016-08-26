package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	pt := point{x: 3, y: 4}
	fmt.Printf("x: %d\n", pt.x)
	fmt.Printf("y: %d\n", pt.y)
}
