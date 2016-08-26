package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for run := 0; run < 3; run++ {
		wg.Add(1)
		go func(run int) {
			for i := 0; i < 2; i++ {
				fmt.Printf("run %d: %d\n", run, i)
			}
			wg.Done()
		}(run)
	}

	wg.Wait()
}
