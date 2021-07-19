//using make to Control the Capacity of a Slice

package main

import "fmt"

// return 3 int Slices

func genSlices() ([]int, []int, []int) {
	var s1 []int

	// define slice using make + set only the length

	s2 := make([]int, 10)

	// // define slice using make + set  the length and capacity
	s3 := make([]int, 10, 50)

	return s1, s2, s3

}

func main() {
	s1, s2, s3 := genSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))

}
