// main.go, example to create arrays
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Define arrays")
	var numbers[5] int
	var cities[5] string
	var matrix[3][3] int

	// Insert data
	fmt.Println(">>>>>> Insert array data")
	for i := 0; i<5; i++ {
		numbers[i] = i
		cities[i] = fmt.Sprintf("City %d", i)
	}

	// Insert matrix data
	fmt.Println(">>>>>> Insert matrix data")
	for i := 0; i<3;i++ {
		for j := 0; j<3;j++ {
			matrix[i][i] = rand.Intn(100)
		}
	}

	// Display data
	fmt.Println(">>>>> Display array data")
	for j:=0; j<5;j++ {
		fmt.Println(numbers[j])
		fmt.Println(cities[j])
	}

	fmt.Println(">>>>> Display matrix data")
	for i := 0; i<3;i++ {
        	for j := 0; j<3;j++ {
		fmt.Println(matrix[i][j])
		}
	}
}
