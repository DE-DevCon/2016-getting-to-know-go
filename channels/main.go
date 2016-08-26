package main

import "fmt"

func main() {
	out := make(chan string)
	go func() {
		out <- "message"
	}()
	fmt.Println(<-out)
}
