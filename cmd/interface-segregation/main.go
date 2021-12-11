package main

import (
	"fmt"
	"github.com/FelipeAz/solid/internal/app/interface-segregation/category"
	"github.com/FelipeAz/solid/internal/app/interface-segregation/product"
)

/*** Interface Segregation Principle (ISP) ***/
/* “Many client specific interfaces are better than one general purpose interface” — Robert C. Martin */

/*
The Interface Segregation Principle suggests that if a class provides methods to multiple clients, then rather than
having a generic interface loaded with all methods, provide a separate interface for each client and implement
all of them in the class.
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
	products = append(products, product.NewProduct("Water Bottle", beverage, 0.25, 1))
	products = append(products, product.NewProduct("Zero Coke", beverage, 0.95, 5))
	products = append(products, product.NewProduct("Coffee Mocha", hotBeverages, 2.50, 1))
	// For that Principle, We have split the Product interface in two -> ProductInterface & ProductOfferInterface
	// ProductOfferInterface has a method called GetPriceWithOffer, but we can still using both interfaces
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
