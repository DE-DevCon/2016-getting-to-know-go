package main

import (
	"fmt"
	"time"
)

func main() {
	for run := 0; run < 3; run++ {
		go concurrentFunction(run)
	}

	time.Sleep(500 * time.Millisecond)
}

func concurrentFunction(run int) {
	for i := 0; i < 2; i++ {
		fmt.Printf("run %d: %d\n", run, i)
	}
}
