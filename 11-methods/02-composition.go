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

type PerishableProduct struct {
	Product
	Expiry string
}

// overriding the Product.Format() method in PerishableProduct
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func main() {

	var milk = PerishableProduct{
		Product: Product{
			Id:   200,
			Name: "Milk",
			Cost: 50,
		},
		Expiry: "2 Days",
	}

	fmt.Println("Before applying discount, milk: ", milk.Format())
	milk.ApplyDiscount(10)
	fmt.Println("After applying discount, milk: ", milk.Format())

}
