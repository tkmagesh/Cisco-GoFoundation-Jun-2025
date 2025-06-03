package main

import "fmt"

func main() {
	/*
		var productRanks map[string]int
		productRanks = make(map[string]int)
		productRanks["pen"] = 5
	*/

	// productRanks := map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	productRanks := map[string]int{
		"pen":    5,
		"pencil": 1,
		"marker": 3,
	}
	productRanks["scribble-pad"] = 4
	fmt.Println(productRanks)

	fmt.Printf("len(productRanks) : %d\n", len(productRanks))

	// iterating a map
	fmt.Println("iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// check for the existence of key
	// keyToCheck := "mouse" //non existant key
	keyToCheck := "pencil"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key [%q] doest not exist\n", keyToCheck)
	}

	// remove an item
	// keyToRemove := "pencil"
	keyToRemove := "mouse" // non existant key
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing key [%q], productRanks = %v\n", keyToRemove, productRanks)
}
