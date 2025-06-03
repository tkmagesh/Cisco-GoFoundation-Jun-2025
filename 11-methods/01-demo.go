/* product = Id, Name, Cost */
package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (pPtr *Product) ApplyDiscount(discountPercentage float64) {
	pPtr.Cost = pPtr.Cost * ((100 - discountPercentage) / 100)
}

func main() {

	var p Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	fmt.Println("Before applying discount, p: ", p.Format())
	p.ApplyDiscount(10) // apply 10% discount on p
	fmt.Println("After applying discount, p: ", p.Format())

}
