package main

import "fmt"

func main() {

	//
	// var nos []int

	// custom initialization
	// var nos []int = []int{3, 1, 4, 2, 5}

	// type inference
	// var nos = []int{3, 1, 4, 2, 5}

	// using :=
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// append new items
	nos = append(nos, 60)
	nos = append(nos, 10, 30, 70)

	// append data from another slice
	hundreds := []int{200, 700, 500}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	fmt.Printf("len(nos) : %d\n", len(nos))

	// iterating a slice [indexer]
	fmt.Println("iterating a slice [indexer]")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	// iterating a slice [range]
	fmt.Println("iterating a slice [range]")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	var nos2 = nos // creating a copy of the "underlying pointer to the array"
	nos2[0] = 9999
	fmt.Printf("nos2[0] = %d, nos[0] = %d\n", nos2[0], nos[0])

	// slicing
	fmt.Printf("nos[3:7] = %v\n", nos[3:7])
	fmt.Printf("nos[3:] = %v\n", nos[3:])
	fmt.Printf("nos[:7] = %v\n", nos[:7])

	fmt.Println("Before sorting : nos = ", nos)
	sortSlice(nos) // sort the 'nos' array
	fmt.Println("After sorting : nos = ", nos)

	// How to create a copy?
	nosCopy := make([]int, len(nos))
	copy(nosCopy, nos)
	nosCopy[0] = -8888
	fmt.Printf("nosCopy[0] = %d, nos[0] = %d\n", nosCopy[0], nos[0])
}

func sortSlice(list []int) { // no return result
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
