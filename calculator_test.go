package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		name string
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Two positive numbers that sum to a positive."},
		{a: -1, b: -1, want: -2, name: "Two negative numbers that sum to a negative."},
		{a: -2, b: 2, want: 0, name: "One negative and one positive numbers that cancel out."},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("\n%s\nAdd(%f, %f)\nwanted %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		name string
	}

	testCases := []testCase{
		{a: 4, b: 2, want: 2, name: "Two positive numbers that subtract into a positive."},
		{a: -2, b: -2, want: 0, name: "Two negative numbers that cancel out."},
		{a: -2, b: 2, want: -4, name: "One negative and one positive that subtract into a negative."},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("\n%s\nSubtract(%f, %f)\nwanted %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		name string
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Two positive numbers that multiply into a positive."},
		{a: -2, b: -2, want: 4, name: "Two negative numbers that multiply into a positive."},
		{a: 2, b: -2, want: -4, name: "One positive and one negative number that multiply into a negative."},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("\n%s\nMultiply(%f, %f)\nwanted %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b        float64
		want        float64
		name        string
		errExpected bool
	}

	testCases := []testCase{
		{a: 4, b: 2, want: 2, name: "", errExpected: false},
		{a: -4, b: -2, want: 2, name: "", errExpected: false},
		{a: 4, b: -2, want: -2, name: "", errExpected: false},
		{a: 4, b: 0, want: 0, name: "Divide Positive by Zero", errExpected: true},
		{a: -4, b: 0, want: 0, name: "Divide Negative by Zero", errExpected: true},
		{a: 1, b: 3, want: 0.333333, name: "Result exceeds float64 precision", errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil // is true if there is an non-nil error

		// Error is not what I expected.
		if tc.errExpected != errReceived {
			t.Fatalf("\nDivide(%f, %f): Unexpected error status: %v", tc.a, tc.b, errReceived)
		}
		// Error value was what I expected, and data value doesn't approximate expectations.
		if !tc.errExpected && !calculator.CloseEnough(tc.want, got, 0.001) {
			t.Errorf("\nDivide(%f, %f)\nwanted %f, got %f", tc.a, tc.b, tc.want, got)
		}
		// Otherwise, error value was what I expected, and data value approximates expectations.
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a           float64
		want        float64
		name        string
		errExpected bool
	}

	testCases := []testCase{
		{a: 16, want: 4, name: "Sqrt of an int, that results in an int.", errExpected: false},
		{a: 8, want: 2.828427, name: "Sqrt of an int, that results in a float.", errExpected: false},
		{a: 5.5, want: 2.345207, name: "Sqrt of a float, that results in a float.", errExpected: false},
		{a: -9, want: 0, name: "Sqrt of a negative int, that results in an error.", errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("\nSqrt(%f): Unexpected error status: %v", tc.a, errReceived)
		}

		if !calculator.CloseEnough(tc.want, got, 0.00001) {
			t.Errorf("\nSqrt(%f)\nwanted %f, got %f", tc.a, tc.want, got)
		}
	}
}
