package main

import "fmt"

func main() {
	things := []interface{}{foo{1}, bar{2}, 4, 5.5, "foo"}
	for _, t := range things {
		if f, ok := t.(foo); ok {
			fmt.Printf("foo: %d\n", f.fooID)
		} else if b, ok := t.(bar); ok {
			fmt.Printf("bar: %d\n", b.barID)
		} else if i, ok := t.(int); ok {
			fmt.Printf("int: %d\n", i)
		} else if s, ok := t.(string); ok {
			fmt.Printf("string: %s\n", s)
		} else if f, ok := t.(float64); ok {
			fmt.Printf("float: %f\n", f)
		}
	}
}
