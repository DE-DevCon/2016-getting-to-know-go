package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("go.tpp")
	defer file.Close()

	info, _ := file.Stat()
	fmt.Println(info.Size())
}
