package main

import "fmt"

// Add function
func add(a, b int) int {
	return a + b
}

// Swap function
func swap(x, y string) (string, string) {
	return y, x
}

// Quotient and remainder function
func divide(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	fmt.Println("Sum:", add(3, 4))
	fmt.Println("Swapped:", swap("Hello", "World"))
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
}
