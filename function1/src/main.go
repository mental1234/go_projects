// Function to calculate the average in a series of numbers
package main

import (
	"fmt"
)

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total = total + v
	}
	return total / float64(len(xs))
}

func main() {
	fmt.Println("Hello World!")
	numbers := []float64{98,93,77,82,83}
	fmt.Println(average(numbers))
}
