/*
Create the types to simulate a shopping cart

Cart
	-> Collection of LineItems

LineItem
	-> Created for a product
	-> units

Product
	-> Id, Name, Cost

Functinality
	Create a ShoppingCart instance
	Add line items
	Print all the products & their respective units for each line item
	Print the Cart Value (sum(product cost * units))
*/

package main

import (
	"fmt"
	"strings"
)

type Product struct {
	Id   int
	Name string
	Cost float64
}

func NewProduct(id int, name string, cost float64) *Product {
	return &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
}

/*
func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}
*/

// fmt.Stringer interface implementation
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

type LineItem struct {
	Item  *Product
	Units int
}

func NewLineItem(item *Product, units int) *LineItem {
	return &LineItem{
		Item:  item,
		Units: units,
	}
}

func (li LineItem) GetItemValue() float64 {
	return float64(li.Units) * li.Item.Cost
}

/*
func (li LineItem) Format() string {
	return fmt.Sprintf("%s, Units = %d, Item Total = %0.2f", li.Item.Format(), li.Units, li.GetItemValue())
}
*/

func (li LineItem) String() string {
	return fmt.Sprintf("%s, Units = %d, Item Total = %0.2f", li.Item, li.Units, li.GetItemValue())
}

type ShoppingCart struct {
	LineItems []*LineItem
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (sc *ShoppingCart) AddLineItem(lineItem *LineItem) {
	sc.LineItems = append(sc.LineItems, lineItem)
}

func (sc ShoppingCart) CalculateTotal() float64 {
	var total float64
	for _, li := range sc.LineItems {
		total += li.GetItemValue()
	}
	return total
}

/*
func (sc ShoppingCart) Format() string {
	var sb strings.Builder
	for _, li := range sc.LineItems {
		sb.WriteString(li.Format())
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("Cart Total : %0.2f\n", sc.CalculateTotal()))
	return sb.String()
}
*/

func (sc ShoppingCart) String() string {
	var sb strings.Builder
	for _, li := range sc.LineItems {
		sb.WriteString(fmt.Sprintf("%s\n", li))
	}
	sb.WriteString(fmt.Sprintf("Cart Total : %0.2f\n", sc.CalculateTotal()))
	return sb.String()
}

func main() {
	p1 := NewProduct(100, "Pen", 10)
	p2 := NewProduct(101, "Pencil", 5)
	p3 := NewProduct(103, "Marker", 50)

	sc := NewShoppingCart()
	sc.AddLineItem(NewLineItem(p1, 50))
	sc.AddLineItem(NewLineItem(p2, 20))
	sc.AddLineItem(NewLineItem(p3, 5))

	fmt.Println(sc)
}
