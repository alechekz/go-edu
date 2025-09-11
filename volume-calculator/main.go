package main

import (
	"fmt"

	"alechekz/go-edu/calculator"
)

// main
func main() {

	//cube
	cube := calculator.NewCube(5.0)
	fmt.Printf("Volume of cube %#v, is - %v\n", cube, calculator.Volume(cube))

	//sphere
	sphere := calculator.NewSphere(5.0)
	fmt.Printf("Volume of sphere %#v, is - %v\n", sphere, calculator.Volume(sphere))
}
