/* product = Id, Name, Cost */
package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	/*
		var p Product

		// accessing the attributes (using "." notation)
		p.Id = 100
		p.Name = "Pen"
		p.Cost = 10
	*/

	var p Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)

	var p2 Product
	p2 = p // a copy is created coz 'struct instances are just values'
	p2.Cost = 12
	fmt.Printf("p2.Cost = %0.2f, p.Cost = %0.2f\n", p2.Cost, p.Cost)

	// structs are compared by values
	var x = Product{
		Id:   200,
		Name: "Pencil",
		Cost: 5,
	}

	var y = Product{
		Id:   200,
		Name: "Pencil",
		Cost: 5,
	}

	fmt.Println("x == y :", x == y)

	// Handling references
	var pPtr *Product
	pPtr = &p

	// no need to deference to access the members through a pointer
	fmt.Println(pPtr.Id, pPtr.Name, pPtr.Cost)

	fmt.Println("Before applying discount, p: ", Format(p))
	ApplyDiscount(&p, 10) // apply 10% discount on p
	fmt.Println("After applying discount, p: ", Format(p))
}

// Assignment
/*
 Write the following functions
 Format(?) => returns a string (Id = 100, Name = "pen", Cost = 10) representation of the given product
 ApplyDiscount(?) => update the cost of the given product by applying the given discount percentage

*/

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(pPtr *Product, discountPercentage float64) {
	pPtr.Cost = pPtr.Cost * ((100 - discountPercentage) / 100)
}
