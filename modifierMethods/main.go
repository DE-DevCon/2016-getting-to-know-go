package main

import (
	"fmt"
)

type stringHolder struct {
	content string
}

func (sh stringHolder) setContent(content string) {
	sh.content = content
}

func (sh stringHolder) getContent() string {
	return sh.content
}

func main() {
	sh := stringHolder{"foobar"}
	fmt.Println(sh.getContent())
	sh.setContent("goobaz")
	fmt.Println(sh.getContent())
}
