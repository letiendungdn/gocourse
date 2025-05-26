package main

import "fmt"

func main() {

	// var sliceName[]ElementType

	// var numbers []int
	// var numbers1 []int = []int{1, 2, 3, 4, 5}

	// numbers2 := []int{9, 8, 7}

	// slice := make([]int, 5) // creates a slice with length 5 and capacity 5
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4] // creates a slice from the array a, starting at index 1 and ending before index 4

	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7) // appending an element to the slice
	fmt.Println(slice1)

	slice1Coppy := make([]int, len(slice1)) // creating a new slice with the same length as slice1
	copy(slice1Coppy, slice1) // copying the contents of slice1 to slice1Coppy

	fmt.Println("Slicecopy",slice1Coppy)

	// var nilSlice []int // declaring a nil slice

	
	
}