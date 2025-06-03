/* PerishableProduct = Product attributes + "expiry" */

package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	Product
	// Dummy
	Expiry string
}

func main() {
	milk := PerishableProduct{
		Product: Product{
			Id:   200,
			Name: "Milk",
			Cost: 50,
		},
		Expiry: "2 Days",
	}

	// fmt.Println(milk.Product.Id, milk.Product.Name, milk.Product.Cost, milk.Expiry)
	fmt.Println(milk.Id, milk.Name, milk.Cost, milk.Expiry)
}
