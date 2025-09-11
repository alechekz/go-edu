package calculator

import "math"

// shape is an abstraction for geometric figures
type shape interface {
	volume() float64
}

// cube is a custom type of data that represent geometric figure cube
type cube struct {
	side float64
}

// sphere is a custom type of data that represent geometric figure sphere
type sphere struct {
	radius float64
}

// NewCube returns a pointer to new cube
func NewCube(side float64) *cube {
	return &cube{side}
}

// volume returns volume of the cube c
func (c *cube) volume() float64 {
	return math.Pow(c.side, 3)
}

// NewSphere returns a pointer to new sphere
func NewSphere(r float64) *sphere {
	return &sphere{r}
}

// volume returns volume of the sphere sp
func (sp *sphere) volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(sp.radius, 3)
}

// Volume calculates volume of given shape
func Volume(s shape) float64 {
	return s.volume()
}
