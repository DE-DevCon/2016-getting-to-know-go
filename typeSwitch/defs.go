package main

import "fmt"

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
