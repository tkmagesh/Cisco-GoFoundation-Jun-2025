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

	// dereferencing is not needed when the members of a container is accessed thorugh a pointer
	fmt.Println(nosPtr[0])

	fmt.Println("Before sorting : nos = ", nos)
	sort( /* ? */ ) // sort the 'nos' array
	fmt.Println("After sorting : nos = ", nos)
}

func sort( /* ? */ ) { // no return result

}
