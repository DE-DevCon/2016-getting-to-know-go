package main

import "fmt"

func main() {
	firstName, lastName := getUser()
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}

func getUser() (firstName string, lastName string) {
	firstName = "Robert"
	lastName = "Jones"
	return
}
