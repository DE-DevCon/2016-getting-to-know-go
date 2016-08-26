package main

import "fmt"

func main() {
	things := []interface{}{foo{1}, bar{2}, 4, 5.5, "foo"}
	for _, t := range things {
		switch thing := t.(type) {
		case foo:
			fmt.Printf("foo: %d\n", thing.fooID)
		case bar:
			fmt.Printf("bar: %d\n", thing.barID)
		case int:
			fmt.Printf("int: %d\n", thing)
		case string:
			fmt.Printf("string: %s\n", thing)
		case float64:
			fmt.Printf("float: %f\n", thing)
		}
	}
}
