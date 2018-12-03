package main

import (
	"fmt"
	"math"
)

func main() {
	foo()
	circle_area(3)
	fmt.Println(advance_calculate(3, 3, 3))
	compute(2, 2, 2, "ramon")
	fmt.Println(add(2, 2, 5))
	closure_func("string2")

}

// Simple function
func foo() {
	fmt.Println("FOO function was called")
}

// Function with parameters
func circle_area(r float64) {
	area := math.Pi * math.Pow(r, 2)
	fmt.Println("Circle area (r=%.2f) = %.2f \n", r, area)
}

// Function with returning value
func advance_calculate(a, b, c float64) float64 {
	result := a * b * c
	return result
}

// Function with multiple returning values
func compute(a, b, c float64, name string) (float64, float64, string) {
	result1 := a * b * c
	result2 := a + b + c
	result3 := result1 + result2
	newName := fmt.Sprintf("%s value = %.2f", name, result3)
	return result1, result2, newName
}

// Function with Multiple Parameters and Returning value
func add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// Closure Function
func closure_func(name string) {
	hoo := func(a, b int) {
		result := a * b
		fmt.Println("hoo() = %d \n", result)
	}

	joo := func(a, b int) int {
		return a*b + a
	}

	fmt.Println("Closure_func(%s) was called", name)
	hoo(2, 3)
	val := joo(5, 8)
	fmt.Println("Val from joo() = %d \n", val)
}
