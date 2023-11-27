package calculator

import (
	"errors"
	"math"
)

// Add function adds two numbers.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract function subtracts the second number from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply function multiplies two numbers.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide function divides the first number by the second.
// It returns an error if division by zero is attempted.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("undefined: division by zero")
	}
	return a / b, nil
}

// Power function calculates the power of a number.
// It returns an error if the base is zero and the exponent is negative.
func Power(base, exponent float64) (float64, error) {
	if base == 0 && exponent < 0 {
		return 0, errors.New("undefined: zero raised to a negative power")
	}
	return math.Pow(base, exponent), nil
}

func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot calculate square root of a negative number")
	}
	return math.Sqrt(a), nil
}

func ConvertDegreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

// Sin calculates the sine of an angle in radians.
func Sin(angleInDegrees float64) float64 {
	return math.Sin(ConvertDegreesToRadians(angleInDegrees))
}

func Cos(angleInDegrees float64) float64 {
	return math.Cos(ConvertDegreesToRadians(angleInDegrees))
}

// Tan calculates the tangent of an angle. The input angle is assumed to be in degrees.
func Tan(angleInDegrees float64) float64 {
	return math.Tan(ConvertDegreesToRadians(angleInDegrees))
}
