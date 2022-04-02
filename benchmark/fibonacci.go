package main

import (
	"fmt"
)

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fibonacci1(n-1) + fibonacci1(n-2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	one := 1
	two := 0
	result := 0
	for i := 2; i <= n; i++ {
		result = one + two
		two = one
		one = result
	}
	return result
}

// go run .
func main() {
	result1 := fibonacci1(10)
	fmt.Printf("fibonacci: %d\n", result1)
	result2 := fibonacci2(10)
	fmt.Printf("fibonacci: %d\n", result2)
}
