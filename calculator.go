// Package calculator provides a library for
// simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the
// result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers a and b, and
// returns the result of multiplying a times b.
func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("invalid input: %f, %f  (division by zero is undefined)", a, b)
	}
	return a / b, nil
}

func CloseEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("invalid input: %f (Sqrt of negative numbers is undefined)", a)
	}
	return math.Sqrt(a), nil
}
