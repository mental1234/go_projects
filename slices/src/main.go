// slices example
package main

import "fmt"

func main() {
	fmt.Println("Slices example")
	x := make([]float64, 5, 10)
	fmt.Println(x)
	arr := []float64{1,2,3,4,5,6,7,8,9}
	fmt.Println(arr)
	y := arr[0:5]
	fmt.Println(y)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4,5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}
