package main

import "fmt"

func main() {
	calc := &Calculator{}

	fmt.Println("Calculator Demo")
	fmt.Printf("2 + 3 = %d\n", calc.Add(2, 3))
	fmt.Printf("5 - 3 = %d\n", calc.Subtract(5, 3))
	fmt.Printf("2 * 3 = %d\n", calc.Multiply(2, 3))

	if result, err := calc.Divide(6, 3); err == nil {
		fmt.Printf("6 / 3 = %d\n", result)
	}

	fmt.Printf("Is 4 even? %v\n", calc.IsEven(4))
	fmt.Printf("Max(5, 3) = %d\n", calc.Max(5, 3))
}
