package main

import "fmt"

func main() {
	a := map[string]string{"foo": "bar", "boo": "baz"}
	for index, value := range a {
		fmt.Printf("%s => %s\n", index, value)
	}

	a["goo"] = "gaz"

	for key := range a {
		fmt.Printf("%s\n", key)
	}

	fmt.Println(a)
}
