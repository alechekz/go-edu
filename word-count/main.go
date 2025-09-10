package main

import (
	"fmt"

	"alechekz/go-edu/wordcount"
)

// main
func main() {
	s := "some some words"
	m := wordcount.Count(s)
	fmt.Printf("%#v\n", m)
}
