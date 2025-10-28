package main

import "errors"

// Calculator provides basic mathematical operations
type Calculator struct{}

// Add returns the sum of two integers
func (c *Calculator) Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers
func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers
func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers
func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// IsEven checks if a number is even
func (c *Calculator) IsEven(n int) bool {
	return n%2 == 0
}

// Max returns the maximum of two integers
func (c *Calculator) Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
