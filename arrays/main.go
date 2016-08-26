package main

import "fmt"

func main() {
	a := [2]string{"foo", "bar"}
	b := [...]string{"boo", "baz"}

	for index, value := range a {
		fmt.Printf("%d => %s\n", index, value)
	}

	for index, value := range b {
		fmt.Printf("%d => %s\n", index, value)
	}
}
