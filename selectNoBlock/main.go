package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan time.Time)
	go func() {
		timer := time.NewTimer(2 * time.Second)
		channel <- <-timer.C
	}()
	for {
		select {
		case <-channel:
			fmt.Println("Timer Expired")
			return
		default:
			fmt.Println("No message received")
		}
		time.Sleep(50 * time.Millisecond)
	}
}
