package main

import (
	"fmt"
	"github.com/FelipeAz/solid/internal/app/dependency-inversion/category"
	"github.com/FelipeAz/solid/internal/app/dependency-inversion/product"
)

/*** Dependency Inversion Principle (DIP) ***/
/* “Depend upon Abstractions. Do not depend upon concretions” — Robert C. Martin */

/*
The Dependency Inversion Principle suggests that that every dependency between modules should target an abstract class
or an interface. No dependency should target a concrete class.
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
	hotBeverages := category.NewSubCategory(category.HotBeverages, beverage)

	var products []product.ProductInterface
	// In our example, the NewProduct function is expecting a CategoryInterface instead of Category Object. This
	// gives us more flexibility - Take HotBeverages as an example, it's a subcategory which came from a base category
	// but we can use it as a Category because we have an interface for that.
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
