package main

import "fmt"

func main() {
	a := 5
	b := 8

	conditionalCase(a, b)
}

func conditionalCase(value1 int, value2 int) {
	if value1 > value2 || value1-value2 < value1 {
		fmt.Println("conditional-->a>b || a-b<a")
	} else {
		fmt.Println("..another")
	}
}
