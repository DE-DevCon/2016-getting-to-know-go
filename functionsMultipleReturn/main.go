package main

import "fmt"

func main() {
	firstName, lastName := getUser()
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}

func getUser() (string, string) {
	return "Robert", "Jones"
}
