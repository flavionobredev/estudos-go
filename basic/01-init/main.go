package main

import "fmt"

type Product struct {
	Name          string
	Description   string
	PurchasePrice float32
}

// func (product Product) getSalesPrice() float32 {
// 	product.Name = "Creu!" // This will not change the product.Name in memory
// 	return product.PurchasePrice * 2
// }

func (product *Product) getSalesPrice() float32 {
	product.Name = "Creu!" // This will CHANGE the product.Name in memory
	return product.PurchasePrice * 2
}

func main() {
	product := Product{
		Name:          "Product 1",
		Description:   "This is a product",
		PurchasePrice: 10.99,
	}
	fmt.Println(product.getSalesPrice(), product)
}
