package main

import "fmt"

// https://golang.org/src/fmt/print.go?#L590

type foo struct {
	fooID int
}
type bar struct {
	barID int
}

func (f foo) String() string {
	return fmt.Sprintf("foo: %d", f.fooID)
}
func (b bar) String() string {
	return fmt.Sprintf("bar: %d", b.barID)
}
