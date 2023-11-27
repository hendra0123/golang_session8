package mathops

import "testing"

func TestAddition(t *testing.T) {
	result := Addition(3, 5)
	if result != 8 {
		t.Errorf("Expected 8, but got %f", result)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(7, 4)
	if result != 3 {
		t.Errorf("Expected 3, but got %f", result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(2, 6)
	if result != 12 {
		t.Errorf("Expected 12, but got %f", result)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(8, 2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != 4 {
		t.Errorf("Expected 4, but got %f", result)
	}

	// Test division by zero
	_, err = Division(10, 0)
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}
