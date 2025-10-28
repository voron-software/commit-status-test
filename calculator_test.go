package main

import "testing"

func TestAdd(t *testing.T) {
	calc := &Calculator{}

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"zero", 0, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	calc := &Calculator{}

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", -5, -3, -2},
		{"mixed signs", 5, -3, 8},
		{"zero", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	calc := &Calculator{}

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 6},
		{"negative numbers", -2, -3, 6},
		{"mixed signs", -2, 3, -6},
		{"zero", 0, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	calc := &Calculator{}

	tests := []struct {
		name      string
		a, b      int
		expected  int
		shouldErr bool
	}{
		{"positive numbers", 6, 3, 2, false},
		{"negative numbers", -6, -3, 2, false},
		{"mixed signs", -6, 3, -2, false},
		{"division by zero", 5, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Divide(tt.a, tt.b)
			if tt.shouldErr {
				if err == nil {
					t.Errorf("Divide(%d, %d) expected error but got none", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.expected {
					t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	calc := &Calculator{}

	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{"even positive", 4, true},
		{"odd positive", 5, false},
		{"even negative", -4, true},
		{"odd negative", -5, false},
		{"zero", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.IsEven(tt.n)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %v; want %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	calc := &Calculator{}

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"first larger", 5, 3, 5},
		{"second larger", 3, 5, 5},
		{"equal", 5, 5, 5},
		{"negative numbers", -3, -5, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Max(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
