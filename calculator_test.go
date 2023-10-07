package calculator_test

import (
	"calculator"
	"testing"
)

// TODO: Change function parameters so Add, Subtract, Multiply and Divide can use variadic params
func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name   string
		inputs []float64
		want   float64
	}

	testCases := []testCase{
		{name: "Multiple positive numbers that sum to a positive.",
			inputs: []float64{1, 2, 3, 4, 5, 6}, want: 21},
		{name: "Multiple negative numbers that sum to a negative.",
			inputs: []float64{-2, -4, -6, -8}, want: -20},
		{name: "Multiple negative and positive numbers that cancel out.",
			inputs: []float64{-5, -10, -20, 5, 10, 20}, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.inputs...)

		if tc.want != got {
			t.Errorf("\n%s\nAdd(%v)\nwanted %f, got %f", tc.name, tc.inputs, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name   string
		inputs []float64
		want   float64
	}

	testCases := []testCase{
		{name: "Multiple positive numbers that subtract into a negative.",
			inputs: []float64{50, 25, 10, 5, 4, 3, 2, 1}, want: -100},
		{name: "Multiple negative numbers that subtract into a positive.",
			inputs: []float64{-100, 100, -25, -10, -10, -5}, want: 50},
		{name: "Multiple negative and positive numbers that cancel out.",
			inputs: []float64{25, -25, 10, -10, 5, -5}, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.inputs...)

		if tc.want != got {
			t.Errorf("\n%s\nSubtract(%v)\nwanted %f, got %f", tc.name, tc.inputs, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name   string
		inputs []float64
		want   float64
	}

	testCases := []testCase{
		{name: "Multiple positive numbers that multiply into a positive.",
			inputs: []float64{1, 2, 3, 4, 5, 6, 7}, want: 5040},
		{name: "Multiple negative numbers that multiply into a negative.",
			inputs: []float64{-2, -3, -3, -4, -4, -5, -5}, want: -7200},
		{name: "Multiple positive and negative numbers that multiply into a positive.",
			inputs: []float64{-3, 3, -5, 5, -7, -7}, want: 11025},
		{name: "Multiple positive and negative numbers that multiply into a negative.",
			inputs: []float64{-2, 2, -4, 4, -6, 6, -8}, want: 18432},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.inputs...)

		if tc.want != got {
			t.Errorf("\n%s\nMultiply(%v)\nwanted %f, got %f", tc.name, tc.inputs, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		inputs      []float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Divide only positive numbers. errExpected is false",
			inputs: []float64{50, 2, 5, 5}, want: 1, errExpected: false},
		{name: "Divide only negative numbers. errExpected is false",
			inputs: []float64{-50, -2, -5, -5}, want: 1, errExpected: false},
		{name: "Divide negative and positive numbers. errExpected is false",
			inputs: []float64{-50, 2, 5, -5}, want: 1, errExpected: false},
		{name: "Divide Positive by Zero. errExpected is true",
			inputs: []float64{10, 5, 0}, want: 0, errExpected: true},
		{name: "Divide Negative by Zero. errExpected is true",
			inputs: []float64{-9, 3, 0}, want: 0, errExpected: true},
		{name: "Result exceeds float64 precision. errExpected is false",
			inputs: []float64{3, 3, 3, 3}, want: 0.111111, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.inputs...)
		errReceived := err != nil // is true if there is an non-nil error

		// Error is not what I expected.
		if tc.errExpected != errReceived {
			t.Fatalf("\nDivide(%v): Unexpected error status: %v", tc.inputs, errReceived)
		}
		// Error value was what I expected, and data value doesn't approximate expectations.
		if !tc.errExpected && !calculator.CloseEnough(tc.want, got, 0.001) {
			t.Errorf("\nDivide(%v)\nwanted %f, got %f", tc.inputs, tc.want, got)
		}
		// Otherwise, error value was what I expected, and data value approximates expectations.
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       float64
		want        float64
		name        string
		errExpected bool
	}

	testCases := []testCase{
		{input: 16, want: 4, name: "Sqrt of an int, that results in an int.", errExpected: false},
		{input: 8, want: 2.828427, name: "Sqrt of an int, that results in a float.", errExpected: false},
		{input: 5.5, want: 2.345207, name: "Sqrt of a float, that results in a float.", errExpected: false},
		{input: -9, want: 0, name: "Sqrt of a negative int, that results in an error.", errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.input)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("\nSqrt(%f): Unexpected error status: %v", tc.input, errReceived)
		}

		if !calculator.CloseEnough(tc.want, got, 0.00001) {
			t.Errorf("\nSqrt(%f)\nwanted %f, got %f", tc.input, tc.want, got)
		}
	}
}
