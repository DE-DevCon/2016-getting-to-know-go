package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			fmt.Println("Two")
		case 4:
			fmt.Println("Four")
		case 6:
			fmt.Println("Six")
			fallthrough
		case 8:
			fmt.Println("Eight")
		default:
			fmt.Println(i)
		}
	}
}
