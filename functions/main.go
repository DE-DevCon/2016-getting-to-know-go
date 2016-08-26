package main

import "fmt"

func main() {
	fmt.Println(sayHello("Bob"))
}

func sayHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
