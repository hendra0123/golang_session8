package mathops

import (
	"errors"
	"math"
)

// Addition function returns the sum of two numbers.
func Addition(a, b float64) float64 {
	return a + b
}

// Subtraction function returns the difference between two numbers.
func Subtraction(a, b float64) float64 {
	return a - b
}

// Multiplication function returns the product of two numbers.
func Multiplication(a, b float64) float64 {
	return a * b
}

// Division function returns the result of dividing two numbers.
// It also handles division by zero and returns an error in such cases.
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot calculate square root of a negative number")
	}
	return math.Sqrt(a), nil
}
