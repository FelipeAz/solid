package main

import (
	"fmt"
	"github.com/FelipeAz/solid/internal/app/liskov-substitution/category"
	"github.com/FelipeAz/solid/internal/app/liskov-substitution/product"
)

/*** Liskov Substitution Principle (LSP) ***/
/* “Derived methods should expect no more and provide no less” — Robert C. Martin */

/*
The Liskov Substitution Principle suggests that a base class should be able to function properly of all derived class
*/

func main() {
	taxMap := map[string]float64{
		"BRL": 1.30,
		"USD": 1.15,
		"KYD": 1.45,
		"EUR": 1.05,
	}

	meat := category.NewCategory(category.Meat, 0.20)
	beverage := category.NewCategory(category.Beverage, 0.05)
	// SubCategory came from the Category interface but it has its own differences, even though we can use it at
	// every place which we use the base module - Category
	hotBeverages := category.NewSubCategory(category.HotBeverages, beverage)

	var products []product.ProductInterface
	products = append(products, product.NewProduct("Water Bottle", beverage, 0.25, 1))
	products = append(products, product.NewProduct("Zero Coke", beverage, 0.95, 5))
	products = append(products, product.NewProduct("Coffee Mocha", hotBeverages, 2.50, 1))
	products = append(products, product.NewProductOffer(
		product.NewProduct("Sirloin Steak", meat, 7.50, 2),
		0.30),
	)

	for coin, tax := range taxMap {
		fmt.Printf("Total %s: $ %.2f (Tax Included)\n", coin, GetTotal(products, tax))
	}
}

//GetTotal should not be responsible to calculate the products price
func GetTotal(products []product.ProductInterface, tax float64) (total float64) {
	for _, p := range products {
		total += p.GetPriceWithTax(tax)
	}
	return
}
