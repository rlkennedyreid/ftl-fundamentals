package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

type testCase struct {
	name          string
	a, b          float64
	want          float64
	errorExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{name: "Add two of the same numbers", a: 2, b: 2, want: 4},
		{name: "Add two different numbers", a: 1, b: 2, want: 3},
		{name: "Add 0 to a number", a: 5, b: 0, want: 5},
		{name: "Add a positive and negative number", a: 5, b: -6, want: -1},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAddRandomFloats(t *testing.T) {
	t.Parallel()

	for a := 1; a <= 100; a++ {
		first, second := rand.Float64(), rand.Float64()

		want := first + second
		got := calculator.Add(first, second)
		if want != got {
			t.Fatalf("Adding random integers step %v: Add(%f, %f): want %f, got %f", a, first, second, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{name: "Subtract same number", a: 2, b: 2, want: 0},
		{name: "Subtract two different numbers", a: 3, b: 2, want: 1},
		{name: "Subtract 0", a: 5, b: 0, want: 5},
		{name: "Subtract a negative number from a positive", a: 5, b: -6, want: 11},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Subtract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{name: "Multiply two of the same numbers", a: 2, b: 2, want: 4},
		{name: "Multiply two different numbers", a: 3, b: 2, want: 6},
		{name: "Multiply a number by 0", a: 5, b: 0, want: 0},
		{name: "Multiply a positive and negative number", a: 5, b: -6, want: -30},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Multiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{errorExpected: false, name: "Divide by same number", a: 2, b: 2, want: 1},
		{errorExpected: false, name: "Divide two different numbers", a: 3, b: 2, want: 1.5},
		{errorExpected: false, name: "Divide a positive number by a negative number", a: 12, b: -6, want: -2},
		{errorExpected: true, name: "Divide a number by 0", a: 1, b: 0, want: 999},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		errReceived := err != nil
		// We want an error if, and only if, we expect one
		if tc.errorExpected != errReceived {
			t.Fatalf("%s: Divide(%f, %f): Unexpected error status: %v", tc.name, tc.a, tc.b, errReceived)
		} else if tc.want != got {
			t.Errorf("%s: Divide(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}
