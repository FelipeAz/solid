package main

import (
	"fmt"
	"github.com/FelipeAz/solid/internal/app/single-responsibility/category"
	"github.com/FelipeAz/solid/internal/app/single-responsibility/product"
)

/*** Single Responsibility Principle (SRP) ***/
/* “Do one thing and do it well” — McIlroy (Unix philosophy) */

/*
The Single Responsibility Principle suggests that your class should have only one responsibility embedded to it and
every change that needs to be made on it should be originated from only one reason.
*/

func main() {
	taxMap := map[string]float64{
		"BRL": 1.30,
		"USD": 1.15,
		"KYD": 1.45,
		"EUR": 1.05,
	}

	beverage := category.NewCategory(category.Beverage, 0.05)
	meat := category.NewCategory(category.Meat, 0.20)

	var products []product.Product
	products = append(products, product.NewProduct("Water Bottle", beverage, 0.25, 1))
	products = append(products, product.NewProduct("Zero Coke", beverage, 0.95, 5))
	products = append(products, product.NewProduct("Sirloin Steak", meat, 7.50, 2))

	for coin, tax := range taxMap {
		fmt.Printf("Total %s: $ %.2f (Tax Included)\n", coin, product.GetTotal(products, tax))
	}
}
