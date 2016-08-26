package main

import "fmt"

func main() {
	a := []string{"foo", "bar"}
	for index, value := range a {
		fmt.Printf("%d => %s\n", index, value)
	}

	a = append(a, "baz")

	for _, value := range a {
		fmt.Printf("%s\n", value)
	}
}
