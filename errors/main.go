package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("go.tpp")
	if err != nil {
		log.Fatal(err)
	}

	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info.Size())
}
