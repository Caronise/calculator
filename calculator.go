// Package calculator provides a library for
// simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// All functions return an error to have the same function signature
// So they can be mapped to an arithmetic symbol in Evaluate()

// Add takes multiple number of inputs, and
// returns the result of adding them together.
func Add(inputs ...float64) (float64, error) {
	var total float64
	for _, input := range inputs {
		total += input
	}
	return total, nil
}

// Subtract takes multiple number of inputs, and
// returns the result of subtracting them together.
func Subtract(inputs ...float64) (float64, error) {
	var total float64 = inputs[0]
	for _, input := range inputs[1:] {
		total -= input
	}
	return total, nil
}

// Multiply takes multiple number of inputs, and
// returns the result of multiplying them together.
func Multiply(inputs ...float64) (float64, error) {
	var total float64 = 1 // Because if it's 0, all results will be 0...
	for _, input := range inputs {
		total = total * input
	}
	return total, nil
}

// Divide takes multiple number of inputs, and
// returns the result of dividing them together.
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

// CloseEnough takes two float64 numbers, and checks
// returns true if the absolute difference between them
// is less than or equal to the tolerance value.
func CloseEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

// Sqrt takes a float64 value and returns the square root.
func Sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0, fmt.Errorf("invalid input: %f (Sqrt of negative numbers is undefined)", input)
	}
	return math.Sqrt(input), nil
}

// Evaluate takes a string arithmetic expression of two numerical values, and
// one symbol and returns the result of evaluating the expression as a float64.
func Evaluate(input string) (float64, error) {
	// Split string into a slice of elements
	splitInput := strings.Split(input, " ")
	// Eliminate all white space elements
	var expression []string
	for _, s := range splitInput {
		if s == "" {
			continue
		}
		expression = append(expression, s)
	}
	// If length is not 3, it's invalid
	if len(expression) != 3 {
		return 0, fmt.Errorf("invalid amount of input parameters: %s", input)
	}
	// If 0th and 2nd element can't be converted to float64, it's invalid
	if val1, err := strconv.ParseFloat(expression[0], 64); err == nil {
		if val2, err := strconv.ParseFloat(expression[2], 64); err == nil {
			// Otherwise, use a switch to call appropriate function.
			switch expression[1] {
			case "+":
				return Add(val1, val2)
			case "-":
				return Subtract(val1, val2)
			case "*":
				return Multiply(val1, val2)
			case "/":
				return Divide(val1, val2)
			}
		}
	}
	// Return error if val1 and val2 aren't float64 values.
	return 0, fmt.Errorf("invalid kind of input parameters: %s", input)
}
