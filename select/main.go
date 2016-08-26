package main

import "fmt"

func main() {
	chans := []chan string{make(chan string), make(chan string), make(chan string)}
	for i := 0; i < 3; i++ {
		go concurrentFunction(i, chans[i])
	}
	for i := 0; i < 3; i++ {
		select {
		case m := <-chans[0]:
			fmt.Println(m)
		case m := <-chans[1]:
			fmt.Println(m)
		case m := <-chans[2]:
			fmt.Println(m)
		}
	}
}

func concurrentFunction(i int, c chan string) {
	c <- fmt.Sprintf("i = %d", i)
}
