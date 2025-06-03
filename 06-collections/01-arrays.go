package main

import "fmt"

func main() {
	// var nos [5]int // memory for 5 ints are allocated and initialized with the 'zero' value of int (0)

	// initilization with custom data
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// Use :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// accessing the elements using indexer
	fmt.Println(nos[3])

	// iterating an array [indexer]
	fmt.Println("iterating an array [indexer]")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	// iterating an array [range]
	fmt.Println("iterating an array [range]")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	var nos2 = nos // creates a copy coz "everything is a value in go"
	nos2[0] = 9999
	fmt.Printf("nos2[0] = %d, nos[0] = %d\n", nos2[0], nos[0])

	//handling references
	var nosPtr *[5]int
	nosPtr = &nos

	// accessing the elements using the "pointer to an array"
	fmt.Println((*nosPtr)[0])

	// dereferencing is not needed when the members of a container is accessed through a pointer
	fmt.Println(nosPtr[0])

	fmt.Println("Before sorting : nos = ", nos)
	sort(&nos) // sort the 'nos' array
	fmt.Println("After sorting : nos = ", nos)

	x := [5]int{32, 7, 43, 56, 4}
	y := [5]int{32, 7, 43, 56, 4}
	fmt.Printf("address of x = %p\n", &x)
	fmt.Printf("address of y = %p\n", &y)

	// compared by values
	fmt.Println("x == y : ", x == y)
}

func sort(list *[5]int) { // no return result
	for i := 0; i < (5 - 1); i++ {
		for j := i + 1; j < 5; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
