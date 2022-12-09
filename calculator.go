// Package calculator does simple calculations.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying them together.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of their division.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

// Sqrt takes one number and returns the square root of it.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative number is not a real number")
	}
	return math.Sqrt(a), nil
}
