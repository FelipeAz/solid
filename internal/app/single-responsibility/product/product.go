package product

import "github.com/FelipeAz/solid/internal/app/single-responsibility/category"

type Product struct {
	Name     string
	Category category.Category
	Price    float64
	Amount   int
}

func NewProduct(name string, category category.Category, price float64, amount int) Product {
	return Product{
		Name:     name,
		Category: category,
		Price:    price,
		Amount:   amount,
	}
}

func (p Product) GetPriceWithTax(tax float64) float64 {
	return p.Price * (tax + p.Category.Tax) * float64(p.Amount)
}
