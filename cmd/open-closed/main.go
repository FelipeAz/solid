package main

import (
	"fmt"
	"github.com/FelipeAz/solid/internal/app/open-closed/category"
	"github.com/FelipeAz/solid/internal/app/open-closed/product"
)

/*** Open Closed Principle (OCP) ***/
/* “A module should be open for extensions, but closed for modification” — Robert C. Martin */

/*
The Open Closed Principle suggests that a module should be closed for modification and open for extension, in other
words, we can add new modules or functionalities without requiring modifications to existing modules
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

	var products []product.ProductInterface
	products = append(products, product.NewProduct("Water Bottle", beverage, 0.25, 1))
	products = append(products, product.NewProduct("Zero Coke", beverage, 0.95, 5))

	// Instead of modifying the Product to add an offer, we Extended this Module by creating the ProductOffer
	// To do so, we must use a common interface (ProductInterface) which has all the functions we need
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
