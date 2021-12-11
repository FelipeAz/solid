package main

import (
	"fmt"
)

/*** Single Responsibility Principle (SRP) ***/
/* “Do one thing and do it well” — McIlroy (Unix philosophy) */

/*
The Single Responsibility Principle suggests that your class should have only one responsibility embedded to it and
every change that needs to be made on it should be originated from only one reason.
*/

const (
	Beverage = "Beverage"
	Meat     = "Meat"
)

type Category struct {
	Name string
	Tax  float64
}

func NewCategory(name string, tax float64) Category {
	return Category{
		Name: name,
		Tax:  tax,
	}
}

type Product struct {
	Name     string
	Category Category
	Price    float64
	Amount   int
}

func NewProduct(name string, category Category, price float64, amount int) Product {
	return Product{
		Name:     name,
		Category: category,
		Price:    price,
		Amount:   amount,
	}
}

func (p Product) GetPriceWithTax(tax float64) float64 {
	return p.Price * (tax + p.Category.Tax)
}

//GetTotal should not be responsible to calculate the products price
func GetTotal(products []Product, tax float64) (total float64) {
	for _, p := range products {
		// Good
		total += p.GetPriceWithTax(tax)

		// Bad
		// total += p.price * tax -> generate coupling
	}
	return
}

func main() {
	taxMap := map[string]float64{
		"BRL": 1.30,
		"USD": 1.15,
		"KYD": 1.45,
		"EUR": 1.05,
	}

	beverage := NewCategory(Beverage, 0.05)
	meat := NewCategory(Meat, 0.20)

	var products []Product
	products = append(products, NewProduct("Water Bottle", beverage, 0.25, 1))
	products = append(products, NewProduct("Zero Coke", beverage, 0.95, 5))
	products = append(products, NewProduct("Sirloin Steak", meat, 7.50, 2))

	for coin, tax := range taxMap {
		fmt.Printf("Total %s: $ %.2f (Tax Included)\n", coin, GetTotal(products, tax))
	}
}
