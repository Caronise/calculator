// Package calculator provides a library for
// simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the
// result of adding them together.
func Add(inputs ...float64) float64 {
	var total float64
	for _, input := range inputs {
		total += input
	}
	return total
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(inputs ...float64) float64 {
	var total float64
	for _, input := range inputs {
		total -= input
	}
	return total
}

// Multiply takes two numbers a and b, and
// returns the result of multiplying a times b.
func Multiply(inputs ...float64) float64 {
	var total float64 = 1 // Because if it's 0, all results will be 0...
	for _, input := range inputs {
		total = total * input
	}
	return total
}

func Divide(inputs ...float64) (float64, error) {
	if len(inputs) == 0 {
		return 0, fmt.Errorf("invalid input: %v (length is 0)", inputs)
	}

	var total float64 = inputs[0]
	for _, input := range inputs[1:] {
		if input == 0 {
			return 0, fmt.Errorf("invalid input: %v (division by zero is undefined)", inputs)
		}
		total /= input
	}
	return total, nil
}

func CloseEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func Sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0, fmt.Errorf("invalid input: %f (Sqrt of negative numbers is undefined)", input)
	}
	return math.Sqrt(input), nil
}
