package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// init test case struct
type testCase struct {
	name   string
	data   func() shape
	answer float64
}

// TestVolume tests function Volume
func TestVolume(t *testing.T) {

	//init test cases
	var cases []testCase = []testCase{
		{
			name: "Cube with side 5.0",
			data: func() shape {
				return NewCube(5.0)
			},
			answer: 125.0,
		},
		{
			name: "Cube with site 23.5",
			data: func() shape {
				return NewCube(23.5)
			},
			answer: 12977.875,
		},
		{
			name: "Sphere with radius 10.0",
			data: func() shape {
				return NewSphere(10.0)
			},
			answer: 4188.790204786392,
		},
		{
			name: "Sphere with radius 45.5",
			data: func() shape {
				return NewSphere(45.5)
			},
			answer: 394568.8529263857,
		},
	}

	//run tests
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.answer, Volume(test.data()))
		})
	}
}
