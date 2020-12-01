// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying them together
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the first
// by the second. Returns error if second number is 0
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 999, fmt.Errorf("bad input: %f, %f (division by zero is undefined)", a, b)
	}
	return a / b, nil
}
