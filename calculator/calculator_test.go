package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(3, 5)
	if result != 8 {
		t.Errorf("Expected 8, but got %f", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(7, 4)
	if result != 3 {
		t.Errorf("Expected 3, but got %f", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 6)
	if result != 12 {
		t.Errorf("Expected 12, but got %f", result)
	}
}

func TestDivide(t *testing.T) {
	// Test normal division
	result, err := Divide(8, 2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != 4 {
		t.Errorf("Expected 4, but got %f", result)
	}

	// Test division by zero
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestPower(t *testing.T) {
	result, err := Power(2, 3)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != 8 {
		t.Errorf("Expected 8, but got %f", result)
	}
}
